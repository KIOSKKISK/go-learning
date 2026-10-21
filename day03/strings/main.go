package main

import (
	"fmt"
	"unicode/utf8"
)

func info(s string) {
	fmt.Println(s, utf8.RuneCountInString(s), len(s))
}

func main() {
	info("Go")
	info("Привет!")
	info("Ok 👍")
}
