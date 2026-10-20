package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	university string
	course     string
}

func main() {
	var student1 Student
	student1.name = "Sheldon Cooper"
	student1.age = 40
	student1.university = "MIT"
	student1.course = "Physics"

	fmt.Println("Student 1:", student1)

	person := Person{name: "Leonard Hofstadter", age: 38}
	student2 := Student{Person: person, university: "Caltech", course: "Physics"}

	fmt.Println("Student 2:", student2)

}
