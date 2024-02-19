package main

import "fmt"

func abs(a int32, b int32) int32 {
	diff := a - b
	if diff < 0 {
		return -1 * diff
	}
	return diff
}
func minimumDistances(a []int32) int32 {
	// Write your code here
	sameValue := make(map[int32]int32)
	var minDistance int32 = 9999
	for i := 0; i < len(a); i++ {
		lastIndex := sameValue[a[i]]
		if lastIndex != 0 {
			if lastIndex == -5 {
				lastIndex = 0
			}
			diff := abs(int32(i), lastIndex)
			if diff < minDistance {
				minDistance = diff
			}
		}

		if i == 0 {
			sameValue[a[i]] = -5
		} else {
			sameValue[a[i]] = int32(i)
		}
	}

	if minDistance == 9999 {
		return -1
	}

	return minDistance

}

func main() {
	fmt.Println(minimumDistances([]int32{0, 0, 1, 2, 3, 4}))
}
