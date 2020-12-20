package main

import (
	"fmt"
)

func main() {

	// mencari nilai rata-rata / average
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	length := len(scores)
	sum := 0
	// for i := 0; i < length; i++ {
	// 	sum += scores[i]
	// }

	for _, value := range scores {
		sum += value
	}

	average := (float64(sum) / float64(length))
	fmt.Println("Total: ", sum, " Average : ", average)

	fmt.Println("============================")

	scoresTwo := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScore []int

	// lengthTwo := len(scoresTwo)

	// for i := 0; i < lengthTwo; i++ {
	// 	if scoresTwo[i] >= 90 {
	// 		goodScore = append(goodScore, scoresTwo[i])
	// 	}
	// }

	for _, value := range scoresTwo {
		if value >= 90 {
			goodScore = append(goodScore, value)
		}
	}

	fmt.Println("Good Scores: ", goodScore)
}
