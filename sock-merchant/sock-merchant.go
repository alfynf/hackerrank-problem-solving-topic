package main

import (
	"fmt"
	"math"
)

func sockMerchant(n int32, ar []int32) int32 {
	countSocks := make(map[int32]int32)
	for i := 0; i < len(ar); i++ {
		countSocks[ar[i]]++
	}

	var sumPairs int32
	for _, v := range countSocks {
		countPairs := int32(math.Floor(float64(v) / float64(2)))
		sumPairs = sumPairs + countPairs
	}

	return sumPairs
}
func main() {
	fmt.Println(sockMerchant(7, []int32{1, 2, 1, 2, 1, 3, 2}))
}
