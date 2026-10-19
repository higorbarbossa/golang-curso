package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func Escrever() {
	fmt.Println("Imprimindo do arquivo auxiliar.go do pacote auxiliar")
	escrever2()

	isValid := checkmail.ValidateFormat("higor.s009@gmail.com")

	fmt.Println("Email válido: ", isValid)
}
