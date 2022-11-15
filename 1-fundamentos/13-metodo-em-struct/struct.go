package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliene %s foi desativado\n", c.Nome)
}

func (c *Cliente) Ativar() {
	c.Ativo = true
	fmt.Printf("O cliene %s foi ativado\n", c.Nome)

}

func main() {
	phellipe := Cliente{
		Nome:  "Phellipe",
		Idade: 32,
		Ativo: false,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
	phellipe.Ativar()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)

}
