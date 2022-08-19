package main

import (
	"fmt"
	"os_system/matematica"
	"os_system/teste"
)

func imp_nota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota ", nota)
	} else {
		fmt.Println("Reprovado com nota ", nota)

	}
}

func main() {
	imp_nota(7.3)
	imp_nota(5.1)
	fmt.Println(teste.Somar(2, 5))
	fmt.Println(matematica.Mult(2, 5))
	fmt.Println(matematica.Div(6, 2))

	// limpando e gravando
}
