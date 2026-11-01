package main

import "fmt"

func main() {
	temps := [7]int{3, -2, 5, 8, -1, 4, 6}
	//1) fmt.Println(len(temps))
	//2) fmt.Println(temps[0], temps[len(temps)-1])
	//3) for i := range temps {
	// 	fmt.Println("День №"+strconv.Itoa(i), ": ", +temps[i])
	// }
	// 	4) cold := 0
	// 	for i := range temps {
	// 		if temps[i] < 0 {
	// 			cold++
	// 		}
	// 	}
	// 	fmt.Println(cold)
	// }
	hottest := temps[0]
	for _, t := range temps {
		if t > hottest {
			hottest = t
		}
	}
	fmt.Println(hottest)
}
