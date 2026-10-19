package main

import "fmt"

func isSmall(n int) bool {
	if n < 10 {
		return true
	}
	return false
}

func countThree(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(isSmall(5))
	fmt.Println(isSmall(20))
	fmt.Println(countThree(10))
	fmt.Println(countThree(20))
}
