package main

import (
	"fmt"
	"strings"
)

type NaiveBayes struct {
	whiteCount map[string]int
	blackCount map[string]int
	totalWhite int
	totalBlack int
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		whiteCount: make(map[string]int),
		blackCount: make(map[string]int),
	}
}

func (nb *NaiveBayes) TrainWhite(data string) {
	nb.totalWhite++
	for _, word := range strings.Fields(data) {
		nb.whiteCount[word]++
	}
}

func (nb *NaiveBayes) TrainBlack(data string) {
	nb.totalBlack++
	for _, word := range strings.Fields(data) {
		nb.blackCount[word]++
	}
}

func (nb *NaiveBayes) Predict(data string) string {
	whiteProb := float64(nb.totalWhite) / float64(nb.totalWhite+nb.totalBlack)
	blackProb := float64(nb.totalBlack) / float64(nb.totalWhite+nb.totalBlack)

	for _, word := range strings.Fields(data) {
		whiteProb *= float64(nb.whiteCount[word]+1) / float64(nb.totalWhite+len(nb.whiteCount))
		blackProb *= float64(nb.blackCount[word]+1) / float64(nb.totalBlack+len(nb.blackCount))
	}

	if whiteProb > blackProb {
		return "White"
	} else {
		return "Black"
	}
}

func main() {
	nb := NewNaiveBayes()

	// Training
	nb.TrainWhite("SELECT name FROM users")
	nb.TrainBlack("DROP TABLE users")

	// Prediction
	fmt.Println(nb.Predict("SELECT age FROM users")) // Expected: White
	fmt.Println(nb.Predict("DELETE FROM users"))     // Expected: Black
}
