package main

import (
	"fmt"
	"math"
)

// ScoringStrategy defines the interface for scoring strategies
type ScoringStrategy interface {
	CalculateScore(attempt int) float64
}

// CorrectAnswerStrategy implements scoring for correct answers
type CorrectAnswerStrategy struct {
	BaseScore    float64
	DecreaseRate float64
}

// CalculateScore calculates the score for a correct answer
func (c *CorrectAnswerStrategy) CalculateScore(attempt int) float64 {
	return c.BaseScore / math.Pow(float64(attempt), c.DecreaseRate)
}

// IncorrectAnswerStrategy implements scoring for incorrect answers
type IncorrectAnswerStrategy struct {
	BaseScore    float64
	DecreaseRate float64
}

// CalculateScore calculates the score for an incorrect answer
func (i *IncorrectAnswerStrategy) CalculateScore(attempt int) float64 {
	return i.BaseScore / math.Pow(float64(attempt), i.DecreaseRate)
}

// ScoringSystem handles the overall scoring system
type ScoringSystem struct {
	CorrectStrategy   ScoringStrategy
	IncorrectStrategy ScoringStrategy
}

// CalculateTotalScore calculates the total score based on correct and incorrect attempts
func (s *ScoringSystem) CalculateTotalScore(correctAttempts, incorrectAttempts int) float64 {
	totalCorrectScore := 0.0
	for i := 1; i <= correctAttempts; i++ {
		score := s.CorrectStrategy.CalculateScore(i)
		fmt.Printf("%d. correct answer score: %f\n", i, score)
		totalCorrectScore += score
	}

	totalIncorrectScore := 0.0
	for i := 1; i <= incorrectAttempts; i++ {
		score := s.IncorrectStrategy.CalculateScore(i)
		fmt.Printf("%d. wrong answer score: %f\n", i, score)

		totalIncorrectScore += score
	}

	return totalCorrectScore + totalIncorrectScore
}

func main() {
	// Define strategies
	correctStrategy := &CorrectAnswerStrategy{BaseScore: 1.0, DecreaseRate: 0.8}
	incorrectStrategy := &IncorrectAnswerStrategy{BaseScore: -1.0, DecreaseRate: 0.4}

	// Initialize scoring system
	scoringSystem := &ScoringSystem{
		CorrectStrategy:   correctStrategy,
		IncorrectStrategy: incorrectStrategy,
	}

	// Calculate total score
	correctAttempts := 9
	incorrectAttempts := 3
	totalScore := scoringSystem.CalculateTotalScore(correctAttempts, incorrectAttempts)

	fmt.Printf("Total score: %f\n", totalScore)
}
