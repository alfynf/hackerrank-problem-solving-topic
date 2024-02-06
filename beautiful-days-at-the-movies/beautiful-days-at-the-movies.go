package main

import (
	"fmt"
	"strconv"
)

func reverseString(s string) string {
	var newString string
	for i := len(s) - 1; i >= 0; i-- {
		newString = newString + string(s[i])
	}
	return newString
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var countBeautifulDay int32
	for a := i; a <= j; a++ {
		aString := strconv.Itoa(int(a))
		reverseDay := reverseString(aString)
		reverseDayNum, _ := strconv.Atoi(reverseDay)

		if (a-int32(reverseDayNum))%k == 0 {
			countBeautifulDay++
		}
	}
	return countBeautifulDay
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
