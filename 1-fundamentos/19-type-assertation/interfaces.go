package main

import "fmt"

func main() {
	var minhaVar interface{} = "Phellipe"

	println(minhaVar.(string))

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é: %v e o valor de ok é: %t", res, ok)
}
