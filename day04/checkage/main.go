package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) (string, error) {
	if age < 0 || age > 150 {
		return "", errors.New("неверный возраст")
	} else if age < 18 {
		return "child", nil
	} else {
		return "adult", nil
	}
}

func main() {
	fmt.Println(checkAge(-5))
	fmt.Println(checkAge(10))
	fmt.Println(checkAge(50))
}
