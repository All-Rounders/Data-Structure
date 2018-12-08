package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	a := aStructure{"Gunwoo", 176, 72}
	mySlice = append(mySlice, a)

	a = aStructure{"Yuin", 172, 52}
	mySlice = append(mySlice, a)

	a = aStructure{"Taeyoon", 182, 86}
	mySlice = append(mySlice, a)

	a = aStructure{"Dakyung", 160, 48}
	mySlice = append(mySlice, a)

	fmt.Println("0: ", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight < mySlice[j].weight
	})
	fmt.Println("<: ", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].weight > mySlice[j].weight
	})
	fmt.Println(">: ", mySlice)
}
