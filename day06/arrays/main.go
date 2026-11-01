package main

import "fmt"

// func main() {
// 	scores := [5]int{90, 75, 60, 85, 100}
// 	// fmt.Println(scores[0]) // 90
// 	// fmt.Println(scores[4]) // 100
// 	// scores[1] = 80
// 	// scores[2] = 70 // поменять вторую ячейку
// 	// fmt.Println(scores[2])
// 	// fmt.Println(scores) // [90 80 60 85 100]
// 	// for i := 0; i < len(scores); i++ {
// 	// 	fmt.Println(i, scores[i])
// 	// }
// 	// fmt.Println(len(scores))
// 	fmt.Println(scores[len(scores)-1])
// 	for i := len(scores) - 1; i >= 0; i-- {
// 		fmt.Println(i, scores[i])
// 	}

// }

func main() {
	scores := [5]int{90, 75, 60, 85, 100}
	// for i, s := range scores {
	// 	fmt.Println("Оценка №" + strconv.Itoa(i+1) + ": " + strconv.Itoa(s))
	// }
	count := 0
	for i := range scores {
		if scores[i]+5 <= 100 {
			scores[i] = scores[i] + 5
		} else {
			scores[i] = 100
		}
		count = count + scores[i]
	}
	average := float64(count) / float64(len(scores))
	fmt.Println(scores)
	fmt.Println(count)
	fmt.Println(average)

}
