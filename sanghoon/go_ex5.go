package main

import (
	"fmt"
)

func getMinMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

func main() {
	min, max := getMinMax(3, 4)
	fmt.Printf("min = %d, max = %d\n", min, max)
}
