package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	ID        int
	Name      string `json:"name"`
	CreatedAt time.Time
}

func TestdbConn(t *testing.T) {
	// Scenario 1: utils/newFile.json
	t.Log("Scenario 1: Validating dbConn function for successful database connection")
	
	// Mock data for expected values
	expectedDBUser := "root"
	expectedDBPass := "12345"
	expectedDBName := "gotest"

	// Call the dbConn function
	db := dbConn()

	// Check if db is not nil
	if db == nil {
		t.Errorf("Failed to establish database connection")
	}

	// Check if the database connection parameters match the expected values
	dsn := db.Stats().DriverInfo
	if dsn != "mysql" {
		t.Errorf("Expected database driver: mysql, got: %s", dsn)
	}
	if db.Stats().OpenConnections != 1 {
		t.Errorf("Expected one open connection, got: %d", db.Stats().OpenConnections)
	}

	// Verify the connection details
	dbUser, dbPass, dbName, err := extractDBDetailsFromDSN(dsn)
	if err != nil {
		t.Errorf("Error extracting DB details from DSN: %v", err)
	}

	if dbUser != expectedDBUser {
		t.Errorf("Expected DB user: %s, got: %s", expectedDBUser, dbUser)
	}
	if dbPass != expectedDBPass {
		t.Errorf("Expected DB pass: %s, got: %s", expectedDBPass, dbPass)
	}
	if dbName != expectedDBName {
		t.Errorf("Expected DB name: %s, got: %s", expectedDBName, dbName)
	}

	// Close the database connection
	err = db.Close()
	if err != nil {
		t.Errorf("Error closing database connection: %v", err)
	}
}

func extractDBDetailsFromDSN(dsn string) (dbUser, dbPass, dbName string, err error) {
	// Parse the DSN to extract DB details
	// TODO: Implement the logic to extract DB details from DSN
	return "", "", "", fmt.Errorf("Not implemented")
}
