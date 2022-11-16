package main

import (
	"fmt"
	"github.com/phelliperodrigues/23-pacote-3/matematica"
)

func main() {
	resultado := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", resultado)
	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Printf("%#v", carro)
	carro.andar()
}
