package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero1 int64 = 1000000000000000000
	fmt.Println(numero1)

	var numero2 uint32 = 1000000000
	fmt.Println(numero2)

	// alias
	// rune == int32
	var numero3 rune = 1000000000
	fmt.Println(numero3)

	//byte == uint8
	var numero4 byte = 255
	fmt.Println(numero4)

	// NUMEROS REAIS

	// float32
	var numeroReal1 float32 = 3.141592653589793238462643383279502884197169399375105820974944
	fmt.Println(numeroReal1)

	// float64
	var numeroReal2 float64 = 3.141592653589793238462643383279502884197169399375105820974944
	fmt.Println(numeroReal2)

	numeroReal3 := 3.1415
	fmt.Println(numeroReal3)

	// NUMEROS REAIS FIM

	// variáveis do tipo string

	var string1 string = "Olá"
	fmt.Println(string1)

	string2 := "tchau"
	fmt.Println(string2)

	// variáveis do tipo char
	// char é representado por aspas simples
	char := 'A'
	fmt.Println("Char: ", char)

	// FIM STRINGS

	// VARIÁVEIS VAZIAS
	// em go, quando declaramos uma variável sem atribuir um valor,
	// ela recebe o valor padrão do tipo de dado.
	// Por exemplo, para tipos numéricos, o valor padrão é 0;
	// para strings, é uma string vazia;
	// e para booleanos, é false.
	var (
		intVazio    int32
		floatVazio  float64
		stringVazio string
	)

	fmt.Println("Inteiro vazio:", intVazio)
	fmt.Println("Float vazio:", floatVazio)
	fmt.Println("String vazio:", stringVazio)

	// FIM VARIÁVEIS VAZIAS

	// TIPO ERRO

	var erro error = errors.New("Erro interno do sistema")
	fmt.Println(erro)

}
