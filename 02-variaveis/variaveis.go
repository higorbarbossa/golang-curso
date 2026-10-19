package main

import "fmt"

func main() {
	// declarando uma variavel com tipo explicito
	var variavel1 string = "variavel1 string"
	fmt.Println(variavel1)

	// declarando com inferencia de tipo
	variavel2 := "variavel2 inferida"
	fmt.Println(variavel2)

	// declarando multiplas variaveis com tipo explicito
	var (
		variavel3 string = "variavel3 string"
		variavel4 string = "variavel4 string"
	)

	fmt.Println(variavel3, variavel4)

	// declarando multiplas variaveis com inferencia de tipo
	variavel5, variavel6 := "variavel5 string", "variavel6 string"
	fmt.Println(variavel5, variavel6)

	// declarando uma constante
	const constante1 string = "constante1 string"
	fmt.Println(constante1)
}
