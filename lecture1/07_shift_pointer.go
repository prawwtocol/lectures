package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := make([]int64, 42)
	a[0] = 1
	a[1] = 42

	p := &a[0]
	s := unsafe.Add(unsafe.Pointer(p),
		unsafe.Sizeof(*new(int64)))

	fmt.Println(*(*int64)(s))
}
