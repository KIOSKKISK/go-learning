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

func show(age int) {
	result, err := checkAge(age)
	if err != nil {
		fmt.Println(age, "→ ошибка:", err)
		return
	}
	fmt.Println(age, "→", result)
}

func main() {
	show(-5)
	show(10)
	show(50)
}
