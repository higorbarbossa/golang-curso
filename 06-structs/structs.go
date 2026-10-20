package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func main() {
	// Criando uma instância de Person usando a sintaxe de struct literal
	var p0 Person
	fmt.Println("Person 0:", p0)

	// Setando valores para os campos da struct
	p0.Name = "Daenerys"
	p0.Age = 23
	fmt.Println("Person 0 depois de setar valores:", p0)

	// Criando instâncias de Person usando a sintaxe literal
	p1 := Person{Name: "John", Age: 30}
	p2 := Person{Name: "Arya", Age: 25}

	fmt.Println("Person 1:", p1)
	fmt.Println("Person 2:", p2)

	// Criando uma instância de Person usando o construtor
	p3 := Person{"Tyrion", 35, Address{"Rua dos bobos, 0", "Casterly Rock", "Seven Kingdoms", "12345"}}
	fmt.Println("Person 3:", p3)
}
