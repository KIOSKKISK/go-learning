package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

func main() {
	prices := make([]int, 10000)
	prices[0] = 1000
	for i := 1; i < len(prices); i++ {
		prices[i] = prices[i-1] + rand.IntN(21) - 10
		if prices[i] < 1 {
			prices[i] = 1
		}
	}
	fmt.Println("максимум:", slices.Max(prices), "минута:", slices.Index(prices, slices.Max(prices)),
		"минимум:", slices.Min(prices), "минута:", slices.Index(prices, slices.Min(prices)))

	maxDif := []int{}
	for i := 1; i < len(prices); i++ {
		if (prices[i] - prices[i-1]) == 10 {
			maxDif = append(maxDif, prices[i])
		}
	}
	fmt.Println(10, len(maxDif), maxDif[0:10])

}
