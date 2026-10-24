package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
	Sex  string
}

// type Person2 struct {
// 	Name string
// 	Age  int
// 	City string
// 	Sex  string
// }

func main() {
	// p2 := Person2{Name: "Arseniy", Age: 20, City: "Vlg", Sex: "Male"}
	// println(p2.sex)
	p := Person{Name: "Arseniy", Age: 20, City: "Vlg", Sex: "Male"}

	fmt.Println(p)      // {Igor 20 Moscow}
	fmt.Println(p.Name) // Igor

	p.Age = 21         // поменять поле
	fmt.Println(p.Age) // 21
}
