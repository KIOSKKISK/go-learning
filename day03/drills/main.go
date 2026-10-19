package main

import "fmt"

func isPositive(n int) bool {
	if n > 0 {
		return true
	}
	return false

}

func bigger(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func sign(n int) string {
	if n == 0 {
		return "ноль"
	} else if n > 0 {
		return "плюс"
	}
	return "минус"
}

func countEven(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(isPositive(5))
	fmt.Println(isPositive(-3))

	fmt.Println(bigger(3, 8))
	fmt.Println(bigger(10, 2))

	fmt.Println(sign(7))
	fmt.Println(sign(-1))
	fmt.Println(sign(0))

	fmt.Println(countEven(10))
	fmt.Println(countEven(7))

	i := -2
	for ; i <= 2; i++ {
		fmt.Println(i, sign(i))
	}

}
