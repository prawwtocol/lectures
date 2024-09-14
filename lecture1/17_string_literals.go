package main

func main() {
	println(`abc`) // same as "abc"
	println(`\n
\n`) // same as "\\n\n\\n"
	println("\n")
	println("\"") // same as `"`
	println("Hello, world!\n")
	println("日本語")
	println("\u65e5本\U00008a9e")
	println("\xff\u00FF")
	//println("\uD800")     // illegal: surrogate half
	//println("\U00110000") // illegal: invalid Unicode code point
}
