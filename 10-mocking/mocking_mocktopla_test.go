package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository defines a mock repository
type MockRepository struct {
	mock.Mock
}

// islem is a struct for operations
type islem struct{}

// TestMockRepositoryMockTopla tests the MockTopla function
func TestMockRepositoryMockTopla(t *testing.T) {
	// Setup
	mockRepo := new(MockRepository)
	op := &islem{}

	// Test data
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test case 1: Positive numbers",
			input:    []int{2, 3, 5},
			expected: 10,
		},
		{
			name:     "Test case 2: Negative numbers",
			input:    []int{-2, -3, -5},
			expected: -10,
		},
		{
			name:     "Test case 3: Mixed numbers",
			input:    []int{2, -3, 5, -7},
			expected: -3,
		},
	}

	// Test cases execution
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock expectation
			mockRepo.On("MockTopla", tc.input).Return(tc.expected, nil)

			// Call the function
			result, err := op.MockTopla(tc.input)

			// Assertions
			assert.NoError(t, err, "Unexpected error")
			assert.Equal(t, tc.expected, result, "Incorrect result")
		})
	}

	// Assert all expectations were met
	mockRepo.AssertExpectations(t)
}
