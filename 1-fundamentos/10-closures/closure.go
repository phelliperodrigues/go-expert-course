package main

import "fmt"

func main() {
	somaMultiplicada := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println("Total multiplicado:", somaMultiplicada)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
