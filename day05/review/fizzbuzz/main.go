// package main

// import "fmt"

// func fizzbuzz () {

// 	if n %3 == 0 && n%5==0 {
// 		fmt.Println("FizzBuzz")
// 	} else if n%3 == 0 {
// 		fmt.Println ("Fizz")
// 	} else if n%5 == 0 {
// 		fmt.Println ("Buzz")
// 	} fmt.Println(n(strconv.Itoa))
// }

// func main (n, int) (string) {
// 	for n := 1; n <= 15; i++

// }

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func main() {
	for n := 1; n <= 15; n++ {
		fmt.Println(fizzBuzz(n))
	}
}
