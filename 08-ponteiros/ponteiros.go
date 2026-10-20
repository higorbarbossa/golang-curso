package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1 // b é um ponteiro para a variável variavel1

	fmt.Println("Variável 1:", variavel1)
	fmt.Println("Variável 2:", variavel2)

	variavel1++

	fmt.Println("Variável 1 pós incremento:", variavel1)
	fmt.Println("Variável 2 pós incremento da variável 1:", variavel2)

	// Ponteiro é uma referência de memória,
	// ou seja, um ponteiro aponta para o endereço de memória de uma variável.

	variavel3 := 100
	var ponteiro *int = &variavel3 // o ponteiro recebe o endereço de memória da variável variavel3

	fmt.Println("Variável 3:", variavel3)
	// o ponteiro é uma referência de memória, ou seja, ele aponta para o endereço de memória da variável variavel3
	fmt.Println("Endereço de memória da variável 3:", ponteiro)
	fmt.Println("Endereço de memória da variável 3:", &variavel3)
	fmt.Println("Endereço de memória do ponteiro:", &ponteiro)

	// desreferenciação do ponteiro, ou seja, acessa o valor da variável apontada pelo ponteiro
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiro)

	variavel3++

	fmt.Println("Variável 3 pós incremento:", variavel3)
	fmt.Println("Valor apontado pelo ponteiro pós incremento da variável 3:", *ponteiro)
}
