package main

func main() {
	println(42)
	println(4_2)
	println(0600)
	println(0_600)
	println(0o600)
	println(0o600)
	println(0xBad)
	println(0xB_a_d)
	println(0x_67_7a_2f_cc_40_c6)
	println(170141183460)
	println(170_141183_460)
	println(0_0_0_0_0_0_0_0_0)

	//_42         // an identifier, not an integer literal
	//42_         // invalid: _ must separate successive digits
	//4__2        // invalid: only one _ at a time
	//0_xBadFace  // invalid: _ must separate successive digits
}
