package main

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func isLong(b Book) bool {
	if b.pages > 300 {
		return true
	}
	return false
}

func NewBook(t string, a string, p int) Book {
	return Book{title: t, author: a, pages: p}
}

func main() {
	N1 := NewBook("Морфий", "Булгаков", 200)
	N2 := NewBook("Война и мир", "Толстой", 500)

	fmt.Println(N1, isLong(N1))
	fmt.Println(N2, isLong(N2))
}
