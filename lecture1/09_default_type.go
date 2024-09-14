package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := false
	b := 2.0
	c := 5 // int64
	r := 'r'

	fmt.Println(reflect.TypeOf(a)) // bool
	fmt.Println(reflect.TypeOf(b)) // float64
	fmt.Println(reflect.TypeOf(c)) // int
	fmt.Println(reflect.TypeOf(r)) // int32 ~ rune
}
