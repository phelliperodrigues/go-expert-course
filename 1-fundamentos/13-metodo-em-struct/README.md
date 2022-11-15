# Capitulo Fundamento
## Aula 13 - Métodos em Structs

Em `GO` uma `struct` também tem metodos "como uma `classe`"

Exemplo: 

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
```
output:
```shell
Nome: Phellipe, Idade: 32, Ativo: false
O cliene Phellipe foi ativado
Nome: Phellipe, Idade: 32, Ativo: true

```
Podemos observar que criamos duas funções `Ativar()`, `Desativar()` a diferença delas para uma função normal são o os recivers adicionados apos a palavra `func` no exemplo `(c *Cliente)` isso caracteriza nossas funções como métodos da `struct` Cliente
