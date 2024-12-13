package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db *sql.DB
}

type Task struct {
	ID          int
	Title       string
	StartDate   time.Time
	DueDate     time.Time
	Status      bool
	Priority    bool
	Description string
	CreatedAt   time.Time
}

func (r repository) Add(task Task) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO tasks (title,start_date,due_date,status,priority,description,created_at) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(task.Title, task.StartDate, task.DueDate, task.Status, task.Priority, task.Description, task.CreatedAt)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func TestrepositoryAdd(t *testing.T) {
	// Set up a mock database connection for testing
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer mockDB.Close()

	repo := repository{db: mockDB}

	// Test scenario 1: Add task successfully
	task := Task{
		Title:       "Test Task",
		StartDate:   time.Now(),
		DueDate:     time.Now().AddDate(0, 0, 7),
		Status:      false,
		Priority:    true,
		Description: "Test Description",
		CreatedAt:   time.Now(),
	}

	insertID, err := repo.Add(task)
	if err != nil {
		t.Errorf("Scenario 1: Expected no error, got: %v", err)
	}
	if insertID <= 0 {
		t.Errorf("Scenario 1: Expected positive insert ID, got: %d", insertID)
	}

	t.Log("Scenario 1: Add task successfully passed")

	// Additional scenarios can be added for comprehensive testing
}
