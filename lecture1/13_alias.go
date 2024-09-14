package main

type NewType int64
type AliasType = int64

func main() {
}

//compilation error
//func foo(a NewType) int64 {
//	return a + 2
//}

func bar(a AliasType) int64 {
	return a + 2
}
