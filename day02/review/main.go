package main

import "fmt"

func rectangleArea(length float64, width float64) float64 {
	return length * width
}

func main() {
	const width = 2.5
	height := 4.0
	fmt.Println(rectangleArea(width, height))
	height = 11.0
	fmt.Println(rectangleArea(width, height))
}
