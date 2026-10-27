package main

import "fmt"

type Order struct {
	number  int
	product string
	status  string
}

func NewOrder(n int, p string) Order {
	return Order{number: n, product: p, status: "new"}
}

func main() {
	order1 := NewOrder(1, "phone")
	order2 := NewOrder(2, "tv")
	fmt.Println(order1, order2)
}
