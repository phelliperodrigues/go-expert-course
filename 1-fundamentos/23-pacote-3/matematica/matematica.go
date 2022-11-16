package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 10

type Carro struct {
	Marca string
}

func (c Carro) andar() {
	fmt.Println("Carro Andando....")
}
