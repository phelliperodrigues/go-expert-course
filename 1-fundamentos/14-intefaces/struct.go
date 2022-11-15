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

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
	fmt.Println("Desativado")
}

func (e Empresa) Ativar() {
	fmt.Println("Ativado")
}

type Fornecedor struct {
	Nome string
}

func (f Fornecedor) Desativar() {
	fmt.Println("Desativado")
}

type Pessoa interface {
	Desativar()
	Ativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func Ativacao(pessoa Pessoa) {
	pessoa.Ativar()
}

func main() {
	phellipe := Cliente{
		Nome:  "Phellipe",
		Idade: 32,
		Ativo: false,
	}
	Desativacao(&phellipe)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
	Ativacao(&phellipe)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)

	minhaEmpresa := Empresa{Nome: "Minha Empresa"}
	Ativacao(minhaEmpresa)
	Desativacao(minhaEmpresa)

	//fornecedor := Fornecedor{Nome: "Fornecedor"}
	//Desativacao(fornecedor)

}
