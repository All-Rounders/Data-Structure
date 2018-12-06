package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0

	for i := 0; i < len(arguments); i++ {
		tmp, _ := strconv.Atoi(arguments[i])
		sum = sum + tmp
	}

	fmt.Println("Sum: ", sum)
}
