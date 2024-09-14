package main

import "fmt"

func main() {
	sum := func(a, b int64) int64 {
		return a + b
	}

	fmt.Println(sum(1, 2))
}
