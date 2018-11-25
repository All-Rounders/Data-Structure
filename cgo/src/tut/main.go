package main

/*
#cgo CFLAGS: -g -Wall
#include "test.h"
#include <stdlib.h>
*/
import "C"
import (
		"fmt"
		"unsafe"
)

func main() {
	str := C.CString("Gopher")
	defer C.free(unsafe.Pointer(str))

	num := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.test(str, num)
	fmt.Printf("size = %d\n", size)

	fmt.Println("cgo test Success")
}
