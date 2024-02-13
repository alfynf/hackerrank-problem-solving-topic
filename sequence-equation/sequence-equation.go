package main

import (
	"fmt"
)

func arrayToHash(p []int32) map[int32]int {
	data := make(map[int32]int)
	for i := 0; i < len(p); i++ {
		data[p[i]] = i + 1
	}
	return data
}

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var newArray []int32
	hashP := arrayToHash(p)
	for i := 1; i <= len(p); i++ {
		firstP := hashP[int32(i)]
		secondP := p[firstP-1]

		if i == int(secondP) {
			newArray = append(newArray, int32(hashP[int32(firstP)]))
		}
	}

	return newArray
}

func main() {
	fmt.Println(permutationEquation([]int32{5, 2, 1, 3, 4}))
}
