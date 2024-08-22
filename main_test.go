package main

import (
	"fmt"
	"math"
	"testing"
)

// TestCorrectAnswerStrategy tests the correct answer score calculation
func TestCorrectAnswerStrategy(t *testing.T) {
	strategy := &CorrectAnswerStrategy{BaseScore: 1.0, DecreaseRate: 0.8}

	tests := []struct {
		attempt       int
		expectedScore float64
	}{
		{1, 1.0},
		{2, 0.574349},
		{3, 0.415244},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Attempt %d", tt.attempt), func(t *testing.T) {
			score := strategy.CalculateScore(tt.attempt)
			if !almostEqual(score, tt.expectedScore, 1e-6) {
				t.Errorf("Expected %f but got %f", tt.expectedScore, score)
			}
		})
	}
}

// TestIncorrectAnswerStrategy tests the incorrect answer score calculation
func TestIncorrectAnswerStrategy(t *testing.T) {
	strategy := &IncorrectAnswerStrategy{BaseScore: -1.0, DecreaseRate: 0.4}

	tests := []struct {
		attempt       int
		expectedScore float64
	}{
		{1, -1.0},
		{2, -0.757858},
		{3, -0.644394},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Attempt %d", tt.attempt), func(t *testing.T) {
			score := strategy.CalculateScore(tt.attempt)
			if !almostEqual(score, tt.expectedScore, 1e-6) {
				t.Errorf("Expected %f but got %f", tt.expectedScore, score)
			}
		})
	}
}

// TestCalculateTotalScore tests the overall score calculation
func TestCalculateTotalScore(t *testing.T) {
	correctStrategy := &CorrectAnswerStrategy{BaseScore: 1.0, DecreaseRate: 0.8}
	incorrectStrategy := &IncorrectAnswerStrategy{BaseScore: -1.0, DecreaseRate: 0.4}

	scoringSystem := &ScoringSystem{
		CorrectStrategy:   correctStrategy,
		IncorrectStrategy: incorrectStrategy,
	}

	tests := []struct {
		correctAttempts   int
		incorrectAttempts int
		expectedScore     float64
	}{
		{0, 0, 0.0},
		{9, 0, 3.406627},
		{0, 3, -2.402252},
		{9, 3, 1.004375},
		{5, 5, -0.906491},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Correct: %d, Incorrect: %d", tt.correctAttempts, tt.incorrectAttempts), func(t *testing.T) {
			totalScore := scoringSystem.CalculateTotalScore(tt.correctAttempts, tt.incorrectAttempts)
			if !almostEqual(totalScore, tt.expectedScore, 1e-6) {
				t.Errorf("Expected %f but got %f", tt.expectedScore, totalScore)
			}
		})
	}
}

// almostEqual compares two floating point numbers with a tolerance
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
