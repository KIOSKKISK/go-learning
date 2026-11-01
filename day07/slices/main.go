package main

import (
	"fmt"
	"strconv"
)

func main() {
	items := []string{"хлеб", "молоко", "яйца", "сыр"}
	fmt.Println(items)
	fmt.Println(len(items))
	items[2] = "масло"
	fmt.Println(items[2])
	for i, v := range items {
		fmt.Println(strconv.Itoa(i+1), v)
	}
	var nothing []string
	fmt.Println(len(nothing))

	var scores []int
	scores = append(scores, 90)
	scores = append(scores, 75)
	scores = append(scores, 60)

	scores = append(scores, 85, 100)
	sum := 0
	for _, s := range scores {
		sum = sum + s
	}
	fmt.Println(sum)

	var good []int
	for _, v := range scores {
		if v >= 80 {
			good = append(good, v)
		}
	}
	fmt.Println(good)

	prices := []int{100, 250, 80, 400}
	discounted := make([]int, len(prices))
	for i, p := range prices {
		discounted[i] = p * 9 / 10
	}
	fmt.Println(prices)
	fmt.Println(discounted)

	var cheap []int
	for _, c := range discounted {
		if c < 100 {
			cheap = append(cheap, c)
		}
	}
	fmt.Println(cheap)
}
