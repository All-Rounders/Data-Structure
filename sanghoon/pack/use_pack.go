package main

import (
	"sim_pack"
	"fmt"
)

func main() {
	tmp := sim_pack.Add(5, 10)
	sim_pack.Println(tmp)

	fmt.Println(sim_pack.Pi)
}
