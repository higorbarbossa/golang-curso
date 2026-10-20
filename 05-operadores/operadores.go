package main

import "fmt"

func main() {
	fmt.Println(" --- Operadores Aritméticos ---")

	// ARITMÉTICOS
	soma := 5 + 3
	subtracao := 1 - 3
	multiplicacao := 5 * 3
	divisao := 21 / 3
	modulo := 5 % 3

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Módulo:", modulo)

	var numero1 int32 = 10
	var numero2 int8 = 3

	// para somar tipos diferentes, é necessário fazer a conversão de tipo
	soma2 := numero1 + int32(numero2)
	fmt.Println("Soma2:", soma2)

	// FIM ARITMÉTICOS

	// ATRIBUIÇÃO
	fmt.Println(" --- Operadores de Atribuição ---")

	var variavel1 string = "String 1"
	variavel2 := "String 2"

	fmt.Println("Variável 1:", variavel1)
	fmt.Println("Variável 2:", variavel2)

	// FIM ATRIBUIÇÃO

	// RELACIONAIS
	fmt.Println(" --- Operadores Relacionais ---")

	fmt.Println("10 == 10:", 10 == 10)
	fmt.Println("10 != 10:", 10 != 10)
	fmt.Println("10 < 10:", 10 < 10)
	fmt.Println("10 > 10:", 10 > 10)
	fmt.Println("10 <= 10:", 10 <= 10)
	fmt.Println("10 >= 10:", 10 >= 10)

	// FIM RELACIONAIS

	// LÓGICOS
	fmt.Println(" --- Operadores Lógicos ---")

	verdadeiro, falso := true, false

	fmt.Println("Verdadeiro && Falso:", verdadeiro && falso)
	fmt.Println("Verdadeiro || Falso:", verdadeiro || falso)
	fmt.Println("!Verdadeiro:", !verdadeiro)
	// FIM LÓGICOS

	// UNÁRIOS
	fmt.Println(" --- Operadores Unários ---")

	numero := 10
	fmt.Println("Número:", numero)
	numero++
	fmt.Println("Pós-incremento:", numero)

	numero += 9
	fmt.Println("Número após incremento de 9:", numero)

	// FIM UNÁRIOS

	// TERNÁRIO
	fmt.Println(" --- Operadores Ternários ---")

	// O operador ternário em Go é representado pela estrutura if-else
	// mas não tem uma sintaxe específica como em outras linguagens
	idade := 18
	status := "maior de idade"
	if idade < 18 {
		status = "menor de idade"
	}
	fmt.Println("Status:", status)
}
