package main

const (
	someConst    string = "123"
	anotherConst        = "123"
)

const (
	_ = 1 << (10 * iota)
	Kb
	Mb
	Gb
	Tb
)
