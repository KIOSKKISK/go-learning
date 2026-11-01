package main

import "fmt"

func main() {
	type Order struct {
		ID       int
		Customer string
		Amount   int    // сумма в рублях
		Status   string // "paid", "cancelled" или "pending"
	}

	orders := []Order{
		{1, "Igor", 1500, "paid"},
		{2, "Anna", 3200, "paid"},
		{3, "Vasya", 800, "cancelled"},
		{4, "Masha", 4100, "pending"},
		{5, "Anna", 2500, "paid"},
	}

	// total := 0
	// for _, o := range orders {
	// 	if o.Status == "paid" {
	// 		total = total + o.Amount
	// 	}
	// }
	// fmt.Println(total)

	// cancelled := 0
	// for _, o := range orders {
	// 	if o.Status == "cancelled" {
	// 		cancelled++
	// 	}
	// }
	// fmt.Println(cancelled)

}
