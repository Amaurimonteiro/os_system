package main

import (
	"fmt"
	"sistema_novo/matematica"
	"sistema_novo/multiplica"
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
	fmt.Println(matematica.Somar(2, 5))
	fmt.Println(multiplica.Mult(2, 5))
	fmt.Println(matematica.Divide(6, 2))
}
