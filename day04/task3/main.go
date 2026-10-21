package main

import (
	"fmt"
	"strconv"
)

func show(s string) {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ошибка - изначальная строка не является числом")
		return
	}
	fmt.Println(result * 2)
}

func main() {
	show("3")
	show("abc")
}
