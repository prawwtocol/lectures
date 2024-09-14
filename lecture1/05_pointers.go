package main

func main() {
	var a *int
	println(a)

	t := 5
	println(&t)

	p := &t
	*p = 10

	b := new(int)
	println(b)
	println(*b)

	println(*ptr[int]())
}

func ptr[T any]() *T {
	return new(T)
}
