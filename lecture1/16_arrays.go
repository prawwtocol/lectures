package main

const n = 5

func main() {
	//[2*N] struct { x, y int32 }
	//[1000]*float64
	//[3][5]int
	//[2][2][2]float64  // same as [2]([2]([2]float64))
	fooArray([1024]byte{})
}

func fooArray(a [1024]byte) {

}
