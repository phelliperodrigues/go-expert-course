package main

import (
	"errors"
	"fmt"
)

func main() {
	soma := sum(1, 2)
	fmt.Println(soma)
	soma1, verdadeiro := sumWithBool(10, 50)
	fmt.Printf("O valor somando é maior ou igual 50? %v, o valor é: %d\n", verdadeiro, soma1)
	soma1, verdadeiro = sumWithBool(10, 10)
	fmt.Printf("O valor somando é maior ou igual 50? %v, o valor é: %d\n", verdadeiro, soma1)

	soma2, err := sumWithError(10, 50)
	if err != nil {
		fmt.Printf("Algo deu errado: \"%v\", valor da soma: %d\n", err, soma2)
	}

	soma2, _ = sumWithError(10, 10)
	fmt.Printf("Valor da soma: %d\n", soma2)

}

func sum(a, b int) int {
	return a + b
}

func sumWithBool(a, b int) (int, bool) {
	soma := a + b
	return soma, soma >= 50

}

func sumWithError(a, b int) (int, error) {
	soma := a + b
	if soma >= 50 {
		return 0, errors.New("Soma é maior que 50")
	}
	return soma, nil
}
