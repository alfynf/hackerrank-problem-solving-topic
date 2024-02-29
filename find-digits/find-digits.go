package main

import (
	"fmt"
	"strconv"
)

func findDigits(n int32) int32 {
	// Write your code here
	var countDigits int32
	nString := strconv.Itoa(int(n))
	for _, d := range nString {
		dInt, _ := strconv.Atoi(string(d))
		if dInt != 0 {
			if n%int32(dInt) == 0 {
				countDigits++
			}
		}
	}

	return countDigits
}

func main() {
	fmt.Println(findDigits(124))
}
