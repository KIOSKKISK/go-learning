package main

import (
	"fmt"
	"strconv"
)

func average(a int, b int, c int) float64 {
	return (float64(a+b+c) / 3.0)
}

func main() {

	fmt.Println(average(1, 2, 10))
	age := 25
	word1 := "мне "
	word2 := " лет"
	fmt.Println(word1 + strconv.Itoa(age) + word2)
}
