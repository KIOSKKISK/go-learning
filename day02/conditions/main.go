package main

import "fmt"

func main() {
	score := 150
	if score < 0 || score > 100 {
		fmt.Println("ошибка")
	} else if score >= 90 {
		fmt.Println("отлично")
	} else if score <= 89 && score >= 75 {
		fmt.Println("хорошо")
	} else if score >= 50 && score <= 74 {
		fmt.Println("удовлетворительно")
	} else if score < 50 {
		fmt.Println("не сдал")
	}
}
