package main

import "fmt"

func minMax(a int, b int, c int) (int, int) {
	// if a > b && b > c {
	// 	return c, a
	// } else if b > a && a > c {
	// 	return c, b
	// } else if a > c && c > b {
	// 	return b, a
	// } else if c > a && a > b {
	// 	return b, c
	// } else if b > c && c > a {
	// 	return a, b
	// } else {
	// 	return a, c
	// }
	mn := a
	mx := a

	if b < mn {
		mn = b
	}
	if c < mn {
		mn = c
	}
	if b > mx {
		mx = b
	}
	if c > mx {
		mx = c
	}
	return mn, mx
}

func main() {
	mn, mx := minMax(7, 2, 9)
	fmt.Println(mn, mx) // 2 9
}
