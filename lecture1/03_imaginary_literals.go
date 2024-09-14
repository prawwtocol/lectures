package main

func main() {
	//println(0123i == 0123)
	//println(0123 == 0o123)

	println(0i)
	println(123i) // == 0123i for backward-compatibility
	//println(0123)   // == 0123i for backward-compatibility
	println(0o123i) // == 0o123 * 1i == 83i
	println(0xabci) // == 0xabc * 1i == 2748i
	println(0.i)
	println(2.71828i)
	println(1.e+0i)
	println(6.67428e-11i)
	println(1e6i)
	println(.25i)
	println(.12345e+5i)
	println(0x1p-2i) // == 0x1p-2 * 1i == 0.25i
}
