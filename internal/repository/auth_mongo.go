package repository

import (
	"TaskManager/internal/domain/model"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
	"hash/fnv"
	"time"
)

// Auth defines the interface for authentication-related operations.
type AuthMongo struct {
	collection *mongo.Collection
}

// NewAuthMongo creates a new instance of AuthMongo with the provided MongoDB client and database name.
func NewAuthMongo(client *mongo.Client, dbName string) *AuthMongo {
	return &AuthMongo{
		collection: client.Database(dbName).Collection("users"),
	}
}

// CreateUser inserts a new user into the MongoDB collection and returns the user ID.
func (a *AuthMongo) CreateUser(user model.User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if user.Id.IsZero() {
		user.Id = bson.NewObjectID()
	}

	_, err := a.collection.InsertOne(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("error inserting user: %w", err)
	}

	h := fnv.New32a()
	h.Write(user.Id[:])
	idInt := int(h.Sum32())

	return idInt, nil
}

// GetUser is a repository method for finding users from MongoDb
func (a *AuthMongo) GetUser(username string, password string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"username": username}

	var user model.User
	err := a.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Printf("Password comparison error: %v\n", err)
		return model.User{}, fmt.Errorf("invalid password")
	}

	return user, nil
}
