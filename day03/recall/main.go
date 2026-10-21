package main

import "fmt"

func isSmall(n int) bool { // задача 1
	if n < 10 {
		return true
	}
	return false
}

func countThree(n int) int { // задача 2
	count := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			count++
		}
	}
	return count
}

func temperature(n int) string { // задача 3

	if n < 0 {
		return "холодно"
	} else if n > 20 {
		return "жарко"
	}
	return "нормально"
}

func main() {
	fmt.Println(isSmall(5))
	fmt.Println(isSmall(20))

	fmt.Println(countThree(10))
	fmt.Println(countThree(20))

	fmt.Println(temperature(-5))
	fmt.Println(temperature(15))
	fmt.Println(temperature(30))
}
