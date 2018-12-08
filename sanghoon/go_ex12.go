package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [4]int{1, -2, 14, 0}
	aMap := make(map[string]int)

	length := len(arr)
	for i := 0; i < length; i++ {
		fmt.Printf("%s ", strconv.Itoa(i))
		aMap[strconv.Itoa(i)] = arr[i]
	}

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
