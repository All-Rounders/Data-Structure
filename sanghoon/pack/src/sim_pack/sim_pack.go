package sim_pack

import (
	"fmt"
)

const Pi = "3.14159"

func Add(x, y int) int {
	return x + y
}

func Println(x int) {
	fmt.Println(x)
}

func init() {
	fmt.Println("Init Function of Sim Pack")
}
