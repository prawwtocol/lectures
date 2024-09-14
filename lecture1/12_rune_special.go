package main

import "time"

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * time.Duration(100*i))
		println("\a")
	}

	println("\a")
	println("1\b")
	println("1\f1")
	println("\n")
	println("\r")
	println("1\t1")
	println("1\v")
	println("\\")
	println(string('\''))
	println("\"")
}
