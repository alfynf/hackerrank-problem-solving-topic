package main

import "fmt"

func utopianTree(n int32) int32 {
	// Write your code here
	treeHeight := int32(1)
	for i := 1; i <= int(n); i++ {
		if i%2 != 0 {
			treeHeight = treeHeight * 2
		} else {
			treeHeight++
		}
	}

	return treeHeight

}
func main() {
	fmt.Println(utopianTree(6))
}
