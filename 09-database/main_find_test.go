package main_test

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      string
	Priority    string
	Description string
	CreatedAt   time.Time
}

type repository struct {
	db *sql.DB
}

func (r repository) Find(id int) (Task, error) {
	var task Task

	rows, _ := r.db.Query("SELECT * FROM tasks WHERE id=?", id)
	defer rows.Close()

	if rows.Next() == false {
		return task, errors.New("Kayıt bulunamadı!")
	}

	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Title, &task.StartDate, &task.DueDate, &task.Status, &task.Priority, &task.Description, &task.CreatedAt)
		if err != nil {
			return task, errors.New("Kayıtlar eşleştirilemedi!")
		}
	}

	if err = rows.Err(); err != nil {
		return task, err
	}

	return task, nil
}

func TestrepositoryFind(t *testing.T) {
	// Setting up mock DB or actual DB connection for testing
	// TODO: Initialize a mock DB connection or actual DB connection for testing

	// Mock data for testing
	mockTask := Task{
		ID:          1,
		Title:       "Test Task",
		StartDate:   time.Now(),
		DueDate:     time.Now().Add(time.Hour * 24),
		Status:      "Pending",
		Priority:    "High",
		Description: "This is a test task",
		CreatedAt:   time.Now(),
	}

	// Test Case 1: Valid ID
	t.Log("Testing with a valid ID")
	repo := repository{} // Initialize repository with mock DB connection
	result, err := repo.Find(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if result.ID != mockTask.ID {
		t.Errorf("Expected task ID to be %d, but got %d", mockTask.ID, result.ID)
	}
	// Add more assertions for other fields if needed

	// Test Case 2: Invalid ID
	t.Log("Testing with an invalid ID")
	_, err = repo.Find(999) // Assuming 999 is an invalid ID
	if err == nil {
		t.Error("Expected an error for invalid ID, but got nil")
	}
	expectedErrMsg := "Kayıt bulunamadı!"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message: %s, but got: %s", expectedErrMsg, err.Error())
	}
}
