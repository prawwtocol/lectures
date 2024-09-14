package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := false
	b := 2.0
	c := 5
	r := 'r'

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(r))
}
