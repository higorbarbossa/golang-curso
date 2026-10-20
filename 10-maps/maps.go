package main

import "fmt"

func main() {
	// Create a map to store student names and their grades
	usuario1 := map[string]string{
		"nome":  "Anakin",
		"email": "anakin@jedi.com",
		"idade": "30",
	}

	fmt.Println("Map usuario1:", usuario1)

	usuario2 := map[string]map[string]string{
		"usuario1": {
			"nome":  "Anakin",
			"email": "anakin@jedi.com",
			"idade": "30",
		},
	}
	fmt.Println("Map usuario2:", usuario2)

	usuario3 := map[string]map[string]string{
		"usuario": usuario1,
		"endereco": {
			"rua":    "Rua dos Jedi",
			"numero": "123",
			"cidade": "Coruscant",
		},
	}

	fmt.Println("Map usuario3:", usuario3)
	delete(usuario3["usuario"], "idade")
	fmt.Println("Map usuario3 depois de deletar 'idade':", usuario3)

}
