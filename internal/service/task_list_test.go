package service

import (
	"TaskManager/internal/domain/model"
	"errors"
	"testing"
)

func TestValidateCreateTaskList(t *testing.T) {
	testTable := []struct {
		name     string
		list     model.TaskList
		expected error
	}{
		{
			name: "Valid Task List",
			list: model.TaskList{
				Title:       "Valid Task List",
				Description: "Task List",
			},
			expected: nil,
		},
		{
			name: "Empty Title",
			list: model.TaskList{
				Title:       "",
				Description: "Task List with empty title",
			},
			expected: errors.New("task list name cannot be empty"),
		},
		{
			name: "Empty Description",
			list: model.TaskList{
				Title:       "Task List with empty description",
				Description: "",
			},
			expected: nil, // Assuming empty description is allowed
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCreateTaskList(tt.list)
			if (err != nil && tt.expected == nil) || (err == nil && tt.expected != nil) || (err != nil && tt.expected != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("validateCreateTaskList(%v) = %v, expected %v", tt.list, err, tt.expected)
			}
		})
	}
}

func TestValidateDeleteTaskList(t *testing.T) {
	testTable := []struct {
		name     string
		listId   int
		expected error
	}{
		{
			name:     "Valid ID",
			listId:   1,
			expected: nil,
		},
		{
			name:     "Zero ID",
			listId:   0,
			expected: errors.New("invalid task list ID"),
		},
		{
			name:     "Negative ID",
			listId:   -5,
			expected: errors.New("invalid task list ID"),
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDeleteTaskList(tt.listId)
			if (err != nil && tt.expected == nil) || (err == nil && tt.expected != nil) || (err != nil && tt.expected != nil && err.Error() != tt.expected.Error()) {
				t.Errorf("validateDeleteTaskList(%d) = %v, expected %v", tt.listId, err, tt.expected)
			}
		})
	}
}
