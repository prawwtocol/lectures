package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	wrongPrintRune("a🙂🙃🌚😑😐z") // ð
	printRune("a🙂🙃🌚😑😐z")      // 🙂
	fmt.Println(containsRune("🙂", '🙂'))
	fmt.Println(isLetter('L'))
}

func wrongPrintRune(s string) {
	fmt.Println(string(s[1]))
}

func printRune(s string) {
	fmt.Println(string([]rune(s)[1]))
}

func containsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}
