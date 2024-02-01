package main

import "fmt"

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var countEarlyStudent int32
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			countEarlyStudent++
		}
	}

	if countEarlyStudent < k {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(angryProfessor(3, []int32{1, 3, -1, 2, 0}))
}
