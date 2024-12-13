package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository defines a mock repository for testing purposes
type MockRepository struct {
	mock.Mock
}

// islem is a struct representing the operations to be tested
type islem struct{}

// MockTopla is the function to be tested
func (*islem) MockTopla(sayilar []int) (int, error) {
	toplam := 0
	for i := range sayilar {
		toplam = toplam + sayilar[i]
	}
	return toplam, nil
}

// TestMockRepositoryMockTopla is a unit test function for MockTopla
func TestMockRepositoryMockTopla(t *testing.T) {
	// Initialize test data
	testData := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Sum of positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Sum of negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
		{
			name:     "Sum of mixed numbers",
			input:    []int{1, -2, 3, -4, 5},
			expected: 3,
		},
	}

	operations := &islem{}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			// Perform the operation
			result, err := operations.MockTopla(data.input)

			// Log the scenario
			t.Log("Scenario:", data.name)

			// Check for errors
			assert.NoError(t, err, "Unexpected error occurred")

			// Validate the result
			assert.Equal(t, data.expected, result, "Sum calculation incorrect")
		})
	}
}
