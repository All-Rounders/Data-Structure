package main

import (
	"fmt"
)

func main() {
	y := 4
	square := func(s int) int {
		return s * s
	}
	fmt.Println("Square of", y, "is", square(y))

	square = func(s int) int {
		return s + s
	}
	fmt.Println("Square of", y, "is", square(y))
}
