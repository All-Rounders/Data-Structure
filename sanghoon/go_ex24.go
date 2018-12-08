package main

import (
	"fmt"
	"reflect"
)

func main() {
	type t1 int
	type t2 int

	x1 := t1(1)
	x2 := t2(1)
	x3 := 1

	st1 := reflect.ValueOf(&x1).Elem();
	st2 := reflect.ValueOf(&x2).Elem();
	st3 := reflect.ValueOf(&x3).Elem();

	typeOfx1 := st1.Type()
	typeOfx2 := st2.Type()
	typeOfx3 := st3.Type()

	fmt.Printf("x1 Type: %s\n", typeOfx1)
	fmt.Printf("x2 Type: %s\n", typeOfx2)
	fmt.Printf("x3 Type: %s\n", typeOfx3)

	type aStructure struct {
		x		uint
		y		float64
		text	string
	}

	x4 := aStructure{123, 3.14, "A Structure"}
	st4 := reflect.ValueOf(&x4).Elem()
	typeOfx4 := st4.Type()

	fmt.Printf("x4 Type: %s\n", typeOfx4)
	fmt.Printf("The fields of %s are:\n", typeOfx4)

	for i := 0; i < st4.NumField(); i++ {
		fmt.Printf("%d: Field name: %s ", i, typeOfx4.Field(i).Name)
		fmt.Printf("type: %s ", st4.Field(i).Type())
		fmt.Printf("and value: %v\n", st4.Field(i).Interface())
	}
}
