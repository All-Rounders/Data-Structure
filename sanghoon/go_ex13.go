package main

import (
	"fmt"
	"reflect"
)

func main() {
	type msg struct {
		x int
		y int
		label string
	}

	p1 := msg{23, 12, "A msg"}
	p2 := msg{}
	p2.label = "msg 2"

	s1 := reflect.ValueOf(&p1).Elem()
	s2 := reflect.ValueOf(&p2).Elem()

	fmt.Println("S1 = ", s1)
	fmt.Println("S2 = ", s2)

	fmt.Println("P1 = ", p1)
	fmt.Println("P2 = ", p2)
}
