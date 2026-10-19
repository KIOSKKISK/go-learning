package main

import "fmt"

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func main() {
	for i := 1; i <= 30; i++ {
		fmt.Println(i, isAdult(i))
	}
}
