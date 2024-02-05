package main

import "fmt"

func countingValleys(steps int32, path string) int32 {
	var currAltitude int32
	var countValley int32
	for i := 0; i < int(steps); i++ {
		var diff int32
		if path[i] == 'U' {
			diff = 1
		} else if path[i] == 'D' {
			diff = -1
		}
		if currAltitude == 0 && diff == -1 {
			countValley++
		}
		currAltitude = currAltitude + diff
	}

	return countValley

}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))

}
