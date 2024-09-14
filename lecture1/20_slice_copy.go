package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := make([]int64, 1000)
	b := make([]int64, 1000)

	copyA := intCopy(a)
	copyB := templateCopy(b)

	fmt.Println(reflect.DeepEqual(a, copyA))
	fmt.Println(reflect.DeepEqual(b, copyB))
}

func intCopy(source []int64) []int64 {
	dst := make([]int64, len(source))
	copy(dst, source)

	return dst
}

func templateCopy[S ~[]E, E any](source S) []E {
	dst := make([]E, len(source))
	copy(dst, source)

	return dst
}
