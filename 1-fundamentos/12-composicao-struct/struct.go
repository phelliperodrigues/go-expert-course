package main

import "fmt"

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

func main() {
	phellipe := Cliente{
		Nome:  "Phellipe",
		Idade: 32,
		Ativo: true,
	}
	phellipe.Address.Cidade = "São Paulo"
	phellipe.Address.Estado = "SP"
	phellipe.Address.Numero = 10
	phellipe.Address.Logadouro = "Rua do Endereço que é um logadouro"
	fmt.Println(phellipe)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
}
