package main

import "fmt"

type Phone struct {
	model string
	price int
	has5g bool
}

func NewPhone(m string, p int, h bool) Phone {
	return Phone{model: m, price: p, has5g: h}
}

func main() {
	Phone1 := NewPhone("iphone", 1000, true)
	Phone2 := NewPhone("samsung", 500, false)
	fmt.Println(Phone1, Phone2)
}
