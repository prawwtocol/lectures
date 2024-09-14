package main

import "fmt"

func foo(a **int, p *int) {
	*a = p
}

func main() {
	a := 5
	p := &a

	b := 42

	foo(&p, &b)

	fmt.Println(*p)
}
