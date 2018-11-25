package main

import (
	"fmt"
	"time"
	"math/rand"
)

type stack struct {
	data int
	link *stack
}

func main() {
	// arr := []int {3, 5, 6, 7, 1} // It works!
	// var arr []int // It doesn't work
	// arr := []int {0} // It works!
	var arr []int // It works!

	var top *stack

	arr = init_arr(arr)
	print_arr(arr)

	fmt.Printf("Stack Data Structure\n")

	for i := 0; i < len(arr); i++ {
		push(&top, arr[i]);
	}

	print_stack(top)
}

func init_arr(arr []int) []int {
	rand.Seed(time.Now().UnixNano())

	var cmp int = rand.Intn(7) + 1
	arr = make([]int, cmp)

	for i := 0; i < cmp; i++ {
redo:
		var tmp int = rand.Intn(100)

		for j := 0; j < i; j++ {
			if(arr[j] == tmp) {
				goto redo
			}
		}

		arr[i] = tmp;
	}

	return arr
}

func print_arr(arr []int) {
	println("Arr Data =")

	for i, data := range arr {
		fmt.Printf("arr[%d] = %d\n", i, data)
	}
}

func get_node() *stack {
	var tmp *stack = new(stack)
	tmp.link = nil
	return tmp
}

func push(top **stack, data int) {
	var tmp *stack = *top

	*top = get_node()
	(*top).data = data
	(*top).link = tmp
}

func print_stack(top *stack) {
	for top != nil {
		fmt.Printf("top.data = %d\n", top.data);
		top = top.link
	}
}
