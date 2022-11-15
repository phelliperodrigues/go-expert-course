package main

import "fmt"

func main() {
	salarios := map[string]int{"Phellipe": 1000, "João": 2000, "Pedro": 3000}
	fmt.Println(salarios)
	fmt.Println("Salario:", salarios["Phellipe"])

	delete(salarios, "Phellipe")
	fmt.Println(salarios)

	salarios["Phellipe"] = 20000
	fmt.Println(salarios)
	fmt.Println("Salario:", salarios["Phellipe"])

	novoSalarios := make(map[string]int)
	novoSalarios["Phellipe"] = 25000
	fmt.Println(novoSalarios)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é de %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salario é de %d\n", salario)
	}
}
