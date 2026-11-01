package main

import (
	"fmt"
)

func main() {
	salaries := []int{40000, 55000, 30000, 70000, 48000}
	raised := make([]int, len(salaries))
	for i, v := range salaries {
		raised[i] = v * 115 / 100
	}
	fmt.Println(salaries)
	fmt.Println(raised)
	var high []int
	for _, v := range raised {
		if v > 50000 {
			high = append(high, v)
		}
	}
	fmt.Println(high)

	sum := 0
	for _, v := range raised {
		sum = sum + v
	}
	fmt.Println(sum)

	maxSalary := raised[0]
	for _, v := range raised {
		if v > maxSalary {
			maxSalary = v
		}
	}
	fmt.Println(maxSalary)
}
