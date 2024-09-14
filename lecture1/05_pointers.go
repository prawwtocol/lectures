package main

func main() {
	var a *int
	println(a)

	b := new(int)
	println(b)
	println(*b)

	println(*ptr[int]())
}

func ptr[T any]() *T {
	return new(T)
}
