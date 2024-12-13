package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

// Mocking the DB for testing purposes
type MockDB struct{}

func (m MockDB) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}

func (m MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (m MockDB) LastInsertId() (int64, error) {
	return 1, nil
}

func TestrepositoryAdd(t *testing.T) {
	mockDB := MockDB{}
	repo := repository{db: mockDB}

	// Test scenario 1: utils/newFile.json
	t.Log("Test scenario 1: Adding a new task to the repository")
	task := Task{
		ID:          1, // TODO: Assign a valid ID
		Title:       "Write unit tests",
		StartDate:   time.Now(),
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      false,
		Priority:    true,
		Description: "Write comprehensive unit tests for Go code",
		CreatedAt:   time.Now(),
	}

	lastID, err := repo.Add(task)

	if err != nil {
		t.Errorf("Error adding task: %v", err)
	}

	if lastID == 0 {
		t.Error("Expected last ID to be non-zero")
	}

	t.Logf("Task added successfully. Last ID: %d", lastID)
}
