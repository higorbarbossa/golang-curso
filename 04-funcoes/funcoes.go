package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func subtrair(n1 int8, n2 int8) int8 {
	return n1 - n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	return somar(n1, n2), subtrair(n1, n2)
}

func main() {
	fmt.Println(" - - Funções em Go - - ")

	soma := somar(10, 20)
	fmt.Println(soma)

	// funcao anônima
	soma2 := func(n1 int8, n2 int8) int8 {
		return n1 + n2
	}

	fmt.Println(soma2(7, 25))

	fmt.Println(" - - Funções com múltiplos retornos - - ")

	// chamando a função e recebendo ambos os retornos
	resultadoSoma, resultadoSubtracao := calculosMatematicos(9, 21)
	fmt.Println("soma:", resultadoSoma)
	fmt.Println("subtração:", resultadoSubtracao)

	// ignorando o primeiro retorno da função
	_, resultadoSubtracao2 := calculosMatematicos(9, 27)
	fmt.Println("subtração2:", resultadoSubtracao2)
}
