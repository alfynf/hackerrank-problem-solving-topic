package main

import "fmt"

func minValue(arr []int32) int32 {
	var minNum int32 = 99999
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 && arr[i] < minNum {
			minNum = arr[i]
		}
	}
	return minNum
}

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	minLength := minValue(arr)
	countSticks := []int32{int32(len(arr))}
	var discardedCount int32

	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] - minLength

		if arr[i] <= 0 {
			discardedCount++
		}
		if i == len(arr)-1 {
			if discardedCount == int32(len(arr)) {
				break
			}
			countSticks = append(countSticks, int32(len(arr))-discardedCount)

			i = -1
			minLength = minValue(arr)
			discardedCount = 0
		}
	}

	return countSticks

}

func main() {
	fmt.Println(cutTheSticks([]int32{8, 8, 14, 10, 3, 5, 14, 12}))
}
