package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	name   string
	level  int
	health int
}

func NewPlayer(n string) Player {
	return Player{name: n, level: 1, health: 100}
}

func (p Player) IsAlive() bool {
	return p.health > 0
}

func (p Player) IsBeginner() bool {
	return p.level < 5
}

func (p Player) Info() string {
	return p.name + ": уровень " + strconv.Itoa(p.level) + "; здоровье: " + strconv.Itoa(p.health)

}

func main() {
	player1 := NewPlayer("Igor")
	player2 := NewPlayer("Anna")
	fmt.Println(player1, player2)

	fmt.Println(player1.IsAlive())
	fmt.Println(player2.IsAlive())

	fmt.Println(player1.IsBeginner())
	fmt.Println(player2.IsBeginner())

	fmt.Println(player1.Info())
	fmt.Println(player2.Info())
}
