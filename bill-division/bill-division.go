package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
    var annaBill int32
    for i := 0; i < len(bill); i++ {
        if k != int32(i) {
            annaBill += bill[i]
        }
    }
    bCharged := b
    bActual := annaBill / 2
    diff := bCharged - bActual
    if diff > 0 {
        fmt.Println(diff)
    } else {
        fmt.Println("Bon Appetit")
    }
}

func main() {
	bonAppetit(int32{3,10,2,9}, 1, 12)
}