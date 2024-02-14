package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	var countJump int32
	for i := 0; i < len(c); i++ {
		if i < len(c)-2 && c[i+2] < 1 {
			countJump++
			i = i + 1
		} else if i < len(c)-1 && c[i+1] < 1 {
			countJump++
		}
	}
	return countJump
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
}
