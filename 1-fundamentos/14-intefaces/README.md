# Capitulo Fundamento
## Aula 14 - Interfaces

Interface em `GO` diferentemente de outras linguagens, funciona da seguinte maneira, quando tempo uma interface e QUALQUER struct que implete o metodo que tem declarado na interface

Deferentemente de outras linguages que utiliza a notação `implements` por exemplo, `GO` não tem a necessidade/forma de fazer isso.

```go
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

}
```
output:
```shell
O cliene Phellipe foi desativado
Nome: Phellipe, Idade: 32, Ativo: false
O cliene Phellipe foi ativado
Nome: Phellipe, Idade: 32, Ativo: true
```

Podemos observar que mesmo sem nenhum tipo de referência na struct `Cliente` nossa struct simplemente passou a ser do tipo da interface `Pessoa`

Isso acontece por que o Cliente implementa os metodos `Ativar()` e `Desativar`

Vamos criar outra struct para ver tudo funcionando

```go
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
    minhaEmpresa := Empresa{Nome: "Minha Empresa"}
	Ativacao(minhaEmpresa)
	Desativacao(minhaEmpresa)

}
```
output:
```shell
Ativado
Desativado
```
Podemos ver que a struct `Empresa` também é do interface `Pessoa` por que a mesma implementa os dois metodos, mesmo que o método só tem um print como implementação.

Vale ressaltar que a implementação de uma interface, para ocorrer com sucesso dese ser implementatos TODOS os métodos que contem na interface, vamos criar uma outra struct `Fornecedor` que implementa somente o metodo `Desativar()` e NÃO implementa o método `Ativar()`

```go
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
	
	fornecedor := Fornecedor{Nome: "Fornecedor"}
	Desativacao(fornecedor)

}
```
Error output:
```shell
# command-line-arguments
./struct.go:71:13: cannot use fornecedor (type Fornecedor) as type Pessoa in argument to Desativacao:
	Fornecedor does not implement Pessoa (missing Ativar method)
```

_OBS: Interfaes só aceitam assinaturas de métodos e não aceitam atributos._
