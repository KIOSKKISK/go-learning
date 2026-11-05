package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
)

var ErrInsufficientFunds = errors.New("недостаточно средств")
var ErrInvalidAmount = errors.New("неверная сумма")

type Transaction struct {
	ID     int
	Type   string // "deposit" или "withdraw"
	Amount int
}

type Account struct {
	Owner   string
	Balance int
	History []Transaction // только успешные операции
}

func NewAccount(owner string, balance int) Account {
	return Account{Owner: owner, Balance: balance}
}

func (acc *Account) Deposit(amount int, id int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	} else {
		acc.Balance += amount
		tr := Transaction{ID: id, Type: "deposit", Amount: amount}
		acc.History = append(acc.History, tr)
		return nil
	}
}

func (acc *Account) Withdraw(amount int, id int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	} else if acc.Balance < amount {
		return ErrInsufficientFunds
	} else {
		acc.Balance -= amount
		tr := Transaction{ID: id, Type: "withdraw", Amount: amount}
		acc.History = append(acc.History, tr)
		return nil
	}
}

func (t Transaction) TransInfo() string {
	return ("ID: " + strconv.Itoa(t.ID) + "| Тип: " + (t.Type) + "| Сумма операции: " + strconv.Itoa(t.Amount))
}

func (acc Account) Info() string {
	return acc.Owner + ": " + strconv.Itoa(acc.Balance) + " руб."
}

func main() {
	acc1 := NewAccount("Игорь", 100000)
	count_of_cancels := 0
	count_of_deps := 0
	count_of_withs := 0
	maxWith := 0
	amount_of_deps := []int{}

	for i := 1; i < 16; i++ {
		if rand.IntN(2) == 0 {
			b := rand.IntN(49901) + 100
			acc1.Deposit(b, i)
			count_of_deps += 1
			amount_of_deps = append(amount_of_deps, b)
		} else {
			a := rand.IntN(49901) + 100
			err := acc1.Withdraw(a, i)
			if errors.Is(err, ErrInsufficientFunds) {
				count_of_cancels += 1
			} else {
				count_of_withs += 1
				if a > maxWith {
					maxWith = a
				}
			}
		}
	}

	// [{1 deposit 34545} {2 deposit 47740} {4 deposit 7770} {5 deposit 4706} {7 deposit 18498} {8 deposit 32845} {10 deposit 7478}]
	hist := acc1.History
	sumDep := 0
	sumWith := 0
	for i := 0; i < len(hist); i++ {
		if hist[i].Type == "deposit" {
			sumDep += hist[i].Amount
		} else {
			sumWith += hist[i].Amount
		}

	}

	average_sum_deps := float64(sumDep) / float64(count_of_deps)

	fmt.Println("1)", acc1.Info(), "\n")
	fmt.Println("2)", "успешных операций: ", len(acc1.History), "отклонённых операций: ", count_of_cancels, "\n")
	fmt.Println(acc1.History, "\n")
	fmt.Println("3)", "сумма пополнений:", sumDep, "сумма снятий:", sumWith, "\n")
	fmt.Println("4)", "средняя сумма пополнений:", average_sum_deps, "\n")
	fmt.Println("5)", "максимальное списание:", maxWith, "\n")
	fmt.Println("6)", "первые пять транзацкий: ")
	for _, t := range hist[0:5] {
		fmt.Println(t.TransInfo())
	}
	fmt.Println("\n", "последние пять транзацкий: ")
	for _, t := range hist[len(hist)-5 : len(hist)] {
		fmt.Println(t.TransInfo())
	}
	sorted := slices.Clone(amount_of_deps)
	slices.Sort(sorted)
	fmt.Println("\n", "7)", "топ 5 пополнений:", sorted)

}
