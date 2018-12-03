package main

import "fmt"

type Car struct {
	speed int
	color string
	method interface {}
}

func (c *Car) setSpeed(s int) {
	c.speed = s;
}

func (c *Car) getSpeed() int {
	return c.speed;
}

func main() {
	println("Make Method Test")
	c := new(Car)
	c.setSpeed(33);
	fmt.Printf("Speed = %d\n", c.getSpeed());
}
