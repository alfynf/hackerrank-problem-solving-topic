package main

import (
	"fmt"
	"math"
)

func viralAdvertising(n int32) int32 {
	// Write your code here
	var countLike int32
	countShared := 5
	for i := 1; i <= int(n); i++ {
		numLikeEachDay := math.Floor(float64(countShared) / float64(2))
		countLike = countLike + int32(numLikeEachDay)
		countShared = int(numLikeEachDay) * 3
	}

	return countLike
}

func main() {
	fmt.Println(viralAdvertising(6))
}
