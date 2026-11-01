package main

import (
	"fmt"
	"slices"
)

func main() {
	times := []int{99, 32, 21, 53, 49, 46, 35, 31, 104, 87, 29, 93, 72, 22, 21, 29, 45, 47, 82, 95, 21, 89, 43, 109, 101, 107, 87, 71, 46, 75, 93, 53, 18, 38, 107, 72, 61, 53, 37, 45, 61, 31, 29, 66, 30, 63, 62, 95, 51, 23, 76, 86, 33, 66, 28, 88, 55, 98, 97, 64}
	fmt.Println(len(times))
	sum := 0
	for i := range times {
		sum = sum + times[i]
	}
	t := float64(sum) / float64(len(times))
	fmt.Println(t)
	lateIDs := []int{}
	late := []int{}
	for i := range times {
		if times[i] > 60 {
			late = append(late, times[i])
			lateIDs = append(lateIDs, i+1)
		}
	}
	fmt.Println(late)
	fmt.Println(len(late), float64(len(late))/float64(len(times))*100)
	fmt.Println(slices.Max(late))
	fmt.Println(slices.Index(times, slices.Max(late)) + 1)

	fmt.Println(lateIDs)
}
