package main

import "fmt"

type Movie struct {
	name       string
	year       int
	rating     float64
	waswatched bool
}

func NewMovie(n string, y int, r float64, w bool) Movie {
	return Movie{name: n, year: y, rating: r, waswatched: w}
}

func main() {
	movie1 := NewMovie("Starwars", 2000, 7.5, true)
	movie2 := NewMovie("Avengers", 2009, 8.3, true)
	fmt.Println(movie1, movie2)
}
