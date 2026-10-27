package main

import (
	"errors"
	"fmt"
	"strconv"
)

func parseAge(str string) (age int, err error) {
	age, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	} else if age < 0 || age > 150 {
		return 0, errors.New("возраст вне диапазона")
	} else {
		return age, nil
	}
}

func show(shh string) {
	age, err := parseAge(shh)
	if err != nil {
		fmt.Println("ошибка: ", err)
		return
	}
	fmt.Println("возраст: ", age)
}

func main() {
	show("25")
	show("-3")
	show("abc")
	show("200")
}
