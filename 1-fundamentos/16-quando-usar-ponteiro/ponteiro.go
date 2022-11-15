package main

import "fmt"

func soma(a, b *int) int {
	println("a:", a)
	*a = 50
	return *a + *b
}
func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println("Soma:", soma(&minhaVar1, &minhaVar2))
	fmt.Println("minhaVar1:", minhaVar1)
	println("minhaVar1:", &minhaVar1)

}
