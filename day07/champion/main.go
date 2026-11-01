package main

import "fmt"

func main() {
	ages := []int{23, 41, 19, 35, 52, 28}
	MaxAge := ages[0]
	for _, v := range ages {
		if v > MaxAge {
			MaxAge = v
		}
	}
	fmt.Println(MaxAge)

	minAge := ages[0]
	for _, v := range ages {
		if v < minAge {
			minAge = v
		}
	}
	fmt.Println(minAge)

	oldestAge := ages[0]
	oldestIndex := 0
	for i, v := range ages {
		if v > oldestAge {
			oldestAge = v
			oldestIndex = i
		}
	}
	fmt.Println(oldestAge, oldestIndex+1)

	
	
}
