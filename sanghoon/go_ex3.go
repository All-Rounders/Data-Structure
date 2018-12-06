package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	optI := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			optI = true
			break
		}
	}

	if optI {
		fmt.Println("-i Option Detected!")
		fmt.Print("y/n: ")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You entered: ", answer)
	} else {
		fmt.Println("The -i Option is not set")
	}
}
