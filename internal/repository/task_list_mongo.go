package repository

import (
	"TaskManager/internal/domain/model"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

// Authorization defines the interface for user authentication and authorization operations.
type TaskListMongo struct {
	collection *mongo.Collection
}

// NewTaskListMongo initializes a new TaskListMongo instance with the provided MongoDB client and database name.
func NewTaskListMongo(client *mongo.Client, dbName string) *TaskListMongo {
	return &TaskListMongo{
		collection: client.Database(dbName).Collection("task_lists"),
	}
}

// Create implements the TaskList interface for creating a task list in MongoDB.
func (t *TaskListMongo) Create(userId string, list model.TaskList) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	counterColl := t.collection.Database().Collection("counters")
	var result struct{ Seq int }
	filter := bson.M{"_id": "task_list_id"}
	update := bson.M{"$inc": bson.M{"seq": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err := counterColl.FindOneAndUpdate(ctx, filter, update, opts).Decode(&result)
	if err != nil {
		return 0, err
	}

	list.Id = result.Seq
	list.UserId = userId
	_, err = t.collection.InsertOne(ctx, list)
	if err != nil {
		return 0, err
	}

	return list.Id, nil
}

// GetAll implements the TaskList interface for retrieving all task lists for a user from MongoDB.
func (t *TaskListMongo) GetAll(userId string) ([]model.TaskList, error) {
	// Implementation for retrieving all task lists for a user

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"user_id": userId}
	cursor, err := t.collection.Find(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no task lists found for user %s", userId)
		}
		return nil, fmt.Errorf("error retrieving task lists: %w", err)
	}

	var taskLists []model.TaskList
	if err := cursor.All(ctx, &taskLists); err != nil {
		return nil, fmt.Errorf("error decoding task lists: %w", err)
	}
	if len(taskLists) == 0 {
		return nil, fmt.Errorf("no task lists found for user %s", userId)
	}
	return taskLists, nil
}

// GetById implements the TaskList interface for retrieving a specific task list by ID from MongoDB.
func (t *TaskListMongo) GetById(userId string, listId int) (model.TaskList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId, "id": listId}
	var taskList model.TaskList
	err := t.collection.FindOne(ctx, filter).Decode(&taskList)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.TaskList{}, fmt.Errorf("task list with ID %d not found for user %s", listId, userId)
		}
		return model.TaskList{}, fmt.Errorf("error retrieving task list: %w", err)
	}
	if taskList.UserId != userId {
		return model.TaskList{}, fmt.Errorf("task list with ID %d does not belong to user %s", listId, userId)
	}
	return taskList, nil
}

// Delete implements the TaskList interface for deleting a task list by ID from MongoDB.
func (t *TaskListMongo) Delete(userId string, listId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId, "id": listId}
	_, err := t.collection.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("task list with ID %d not found for user %s", listId, userId)
		}
	}

	return nil
}

// Update implements the TaskList interface for updating a task list in MongoDB.
func (t *TaskListMongo) Update(userId string, listId int, input model.UpdateTaskListInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userId, "id": listId}
	update := bson.M{}
	if input.Title != nil {
		update["title"] = *input.Title
	}
	if input.Description != nil {
		update["description"] = *input.Description
	}
	if len(update) == 0 {
		return errors.New("no fields to update")
	}
	_, err := t.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("task list with ID %d not found for user %s", listId, userId)
		}
		return fmt.Errorf("error updating task list: %w", err)
	}
	return nil
}
