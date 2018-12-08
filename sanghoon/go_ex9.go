package main

import (
	"fmt"
)

func main() {
	my_arr := [4]int{1, 2, 4, -4}
	double := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	tripple := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("my_arr: ", my_arr)
	fmt.Println("double: ", double)
	fmt.Println("tripple: ", tripple)

	double[1][2] = 15
	tripple[0][1][1] = -1

	fmt.Println("double: ", double)
	fmt.Println("tripple: ", tripple)
}
