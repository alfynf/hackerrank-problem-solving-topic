package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	var counter int32
	for i := 0; i <= len(s)-int(m); i++ {
		var sum int32
		for j := i; j <= i+int(m)-1; j++ {
			sum = sum + s[j]
		}
		if sum == d {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(birthday([]int32{2, 2, 1, 3, 2}, 3, 2))
}
