package main

import "fmt"

func main() {
	a := complex(1, 2.5)
	fmt.Println(real(a))
	fmt.Println(imag(a))
}
