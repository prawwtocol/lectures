package main

func main() {
	println(`abc`)                 // same as "abc"
	println(`Привет, 123 123 123`) // same as "\\n\n\\n"
	println("\n")
	println("\"") // same as `"`
	println("Hello, world!\n")
	println("日本語")
	println("\u65e5本\U00008a9e")
	println("\xff\u00FF")
}
