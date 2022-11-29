package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Escrevendo alguma coisa no arquivo"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	f.Close()

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			println(".")
			break
		}

		fmt.Print(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
