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
			name:     "Test case 1: Sum of positive numbers",
			input:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Test case 2: Sum of negative numbers",
			input:    []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
		{
			name:     "Test case 3: Sum of mixed positive and negative numbers",
			input:    []int{1, -2, 3, -4, 5},
			expected: 3,
		},
	}

	// Create an instance of islem
	op := &islem{}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			// Call the function under test
			result, err := op.MockTopla(data.input)

			// Log the scenario details
			t.Log("Scenario:", data.name)
			t.Log("Input:", data.input)
			t.Log("Expected:", data.expected)
			t.Log("Result:", result)

			// Assert the result
			assert.NoError(t, err, "Unexpected error occurred")
			assert.Equal(t, data.expected, result, "Expected and actual results do not match")
		})
	}
}
