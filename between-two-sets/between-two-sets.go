package main

import "fmt"

func getTotalX(a []int32, b []int32) int32 {
	minNum := a[0]
	maxNum := b[len(b)-1]
	
	var countInt int32
	for i := minNum; i <= maxNum; i++ {
		isFirstConValid := true
		isSecConValid := true
		
		// validate first condition
		for j := 0; j < len(a); j++ {
			if i % a[j] != 0 {
				isFirstConValid = false
			}  
		}
		
		// validate second condition
		for k := 0; k < len(b); k++ {
			if b[k] % i != 0 {
				isSecConValid = false
			}  
		}
		
		if isFirstConValid && isSecConValid {
			countInt++
		}
	}
	return countInt
}

func main() {
	fmt.Println(getTotalX([]int32{2,4}, []int32{16,32,96}))
}