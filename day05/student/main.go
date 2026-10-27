package main

import "fmt"

type Student struct {
	name     string
	course   int
	avgScore float32
}

func NewStudent(n string, c int, a float32) Student {
	return Student{name: n, course: c, avgScore: a}
}

func main() {
	Student1 := NewStudent("Igor", 3, 3.8)
	Student2 := NewStudent("Vasya", 2, 4.5)
	fmt.Println(Student1, Student2)
}
