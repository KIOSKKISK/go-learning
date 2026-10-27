package main

import "fmt"

type Cat struct {
	name string
	age  int
}

func NewCat(n string, a int) Cat {
	return Cat{name: n, age: a}
}

func main() {
	fmt.Println(NewCat("Джамал", 200))
}
