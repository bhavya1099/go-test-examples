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

	// Test scenario 1: Inserting a new task successfully
	t.Log("Scenario 1: Inserting a new task successfully")
	testTask := Task{
		ID:          0, // TODO: Update with valid ID if needed
		Title:       "Sample Task",
		StartDate:   time.Now(),
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      true,
		Priority:    false,
		Description: "This is a sample task for testing",
		CreatedAt:   time.Now(),
	}

	insertedID, err := repo.Add(testTask)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if insertedID == 0 {
		t.Error("Expected non-zero ID after insertion")
	} else {
		t.Logf("Task inserted successfully with ID: %d", insertedID)
	}

	// Additional test scenarios can be added here for more coverage
}
