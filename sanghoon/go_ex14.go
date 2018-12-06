package main

import (
	"fmt"
)

type coords interface {
	xAxis() int
	yAxis() int
}

type vector2d struct {
	x int
	y int
}

func (s vector2d) xAxis() int {
	return s.x
}

func (s vector2d) yAxis() int {
	return s.y
}

func findCoord(a coords) {
	fmt.Println("x: ", a.xAxis(), "y: ", a.yAxis())
}

type coord int

func (s coord) xAxis() int {
	return int(s)
}

func (s coord) yAxis() int {
	return 0
}

func main() {
	x := vector2d{x: -1, y: 12}
	fmt.Println(x)
	findCoord(x)

	y := coord(10)
	findCoord(y)
}
