package main

import (
	"fmt"
)

func sort(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}

func main() {
	min, max := sort(7, 3)
	fmt.Printf("min = %d, max = %d\n", min, max)
}
