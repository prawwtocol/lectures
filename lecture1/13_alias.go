package main

type NewType int64
type AliasType = int64

func main() {
	foo5(NewType(int64(1)))
	bar(int64(2))
}

func foo5(a NewType) NewType {
	return a + 2
}

func bar(a AliasType) int64 {
	return a + 2
}
