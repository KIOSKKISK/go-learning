package main

import "fmt"

type Car struct {
	brand string
	year  int
	color string
}

func Newcar(b string, y int, c string) Car {
	return Car{brand: b, year: y, color: c}
}

func main() {
	car1 := Newcar("Cadillac", 2025, "black")
	car2 := Newcar("BMW", 2018, "black")
	fmt.Println(car1, car2)
}
