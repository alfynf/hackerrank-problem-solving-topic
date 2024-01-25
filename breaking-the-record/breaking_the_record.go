package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var maxScore int32
	var minScore int32
	var countBreakMaxScore int32
	var countBreakMinScore int32

	maxScore = scores[0]
	minScore = scores[0]

	for i := 1; i < len(scores); i++ {
		score := scores[i]
		if score < minScore {
			minScore = score
			countBreakMinScore++
		} else if score > maxScore {
			maxScore = score
			countBreakMaxScore++
		}
	}

	return []int32{countBreakMaxScore, countBreakMinScore}
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
