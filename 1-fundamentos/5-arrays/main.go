package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for index, value := range meuArray {
		fmt.Printf("O valor do indece é %d e o valor é %d\n", index, value)
	}
}
