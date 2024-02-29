package main

import "fmt"

// beautiful-triplets
func beautifulTriplets(d int32, arr []int32) int32 {
	// Write your code here
	var countTriplets int32
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j]-arr[i] == d {
				for k := j + 1; k < len(arr); k++ {
					if arr[k]-arr[j] == d {
						countTriplets++
					}
				}
			}
		}
	}
	return countTriplets
}

func main() {
	fmt.Println(beautifulTriplets(1, []int32{2, 2, 3, 4, 5}))
}
