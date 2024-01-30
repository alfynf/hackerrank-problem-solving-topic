package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var maxSum int32
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			sum := keyboards[i] + drives[j]
			if sum <= b && sum > maxSum {
				maxSum = sum
			}
		}
	}
	if maxSum > b || maxSum == 0 {
		return -1
	}

	return maxSum
}

func main() {
	fmt.Println(getMoneySpent([]int32{50, 25, 14}, []int32{20, 10}, 80))
}
