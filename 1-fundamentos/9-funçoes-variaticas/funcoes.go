package main

import "fmt"

func main() {
	soma := sum(1, 2, 3, 4, 5)
	fmt.Println("Soma Total:", soma)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
