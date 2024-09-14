package main

import "fmt"

type MyNewType int64
type MyNewTypeSlice []MyNewType
type MyTemplateNewTypeSlice[T any] []T

// http.Handler example

type MyType string

func (m MyType) Print() string {
	return "example: " + string(m)
}

func main() {
	m := MyNewType(5)
	_ = m

	s := make(MyNewTypeSlice, 10)
	_ = s

	t := make(MyTemplateNewTypeSlice[int64], 10)
	_ = t

	myType := MyType("123")
	fmt.Println(myType)
}
