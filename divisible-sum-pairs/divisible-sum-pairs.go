package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var counter int32
	for i := 0; i < int(n)-1; i++ {
		for j := i + 1; j < int(n); j++ {
			if (ar[i]+ar[j])%k == 0 {
				counter++
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(divisibleSumPairs(6, 5, []int32{1, 2, 3, 4, 5, 6}))

}
