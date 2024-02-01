package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	countBird := make(map[int32]int32)
	var maxCount int32
	for i := 0; i < len(arr); i++ {
		countBird[arr[i]]++
		if countBird[arr[i]] > maxCount {
			maxCount = countBird[arr[i]]
		}
	}

	var maxCountedIds []int32
	for k, v := range countBird {
		if v == maxCount {
			maxCountedIds = append(maxCountedIds, k)
		}
	}

	var minId int32

	if len(maxCountedIds) == 1 {
		return maxCountedIds[0]
	} else {
		minId = maxCountedIds[0]
		for i := 1; i < len(maxCountedIds); i++ {
			if maxCountedIds[i] < minId {
				minId = maxCountedIds[i]
			}
		}
	}
	return minId

}

func main() {
	fmt.Println(migratoryBirds([]int32{8, 4, 2, 5, 10}))
}
