package main

import (
	"fmt"
	"strconv"
)

type Account struct {
	owner   string
	balance int
}

func NewAccount(o string) Account {
	return Account{owner: o}
}

func (a Account) Info() string {
	return a.owner + ": " + strconv.Itoa(a.balance)
}

func main() {
	acc := NewAccount("Igor")
	fmt.Println(acc.Info())
