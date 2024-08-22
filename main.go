package main

import (
	"fmt"
	"math"
)

// Function to calculate score
func CalculateScore(B float64, A float64, attempt int) float64 {
	return B / math.Pow(float64(attempt), A)
}

func main() {
	B := 1.0  // Base score for the first correct answer
	A1 := 0.8 // Decrease factor
	A2 := 0.4 // Decrease factor
	correctAttempts := 9
	incorrectAttempts := 3

	// Total score from correct answers
	totalCorrectScore := 0.0
	for i := 1; i <= correctAttempts; i++ {
		score := CalculateScore(B, A1, i)
		totalCorrectScore += score
		fmt.Printf("%d. correct answer score: %f\n", i, score)
	}

	// Total score from incorrect answers
	BIncorrect := -1.0
	totalIncorrectScore := 0.0
	for i := 1; i <= incorrectAttempts; i++ {
		score := CalculateScore(BIncorrect, A2, i)
		totalIncorrectScore += score
		fmt.Printf("%d. incorrect answer score: %f\n", i, score)
	}

	// Final total score
	totalScore := totalCorrectScore + totalIncorrectScore
	fmt.Printf("Total score: %f\n", totalScore)
}
