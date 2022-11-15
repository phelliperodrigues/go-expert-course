# Capitulo Fundamento
## Aula 12 - Composição de Structs

Em `GO` podemos compor as `structs` para organizarmos nossos códigos, no nosso exemplo passado criamos uma struct `Cliente` que possuia algumas informações, porem gostariamos de adicionar os endereço a este cliente, então vamos ver o exemplo abaixo:

```go
package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
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
	phellipe.Endereco.Cidade = "São Paulo"
	phellipe.Endereco.Estado = "SP"
	phellipe.Numero = 10
	phellipe.Logadouro = "Rua do Endereço que é um logadouro"
	fmt.Println(phellipe)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
}
```
output:
```shell
{Phellipe 32 true {Rua do Endereço que é um logadouro 10 São Paulo SP}}
Nome: Phellipe, Idade: 32, Ativo: true
```
Criamos uma struct `Endereco` e referenciamos ela dentro da struct `Cliente` assim podemos acessar e atribuir os valores de endereço para nossa variavel `phellipe`

Podemos também dar um alias para o endereço 

```go
package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
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
```
output:
```shell
{Phellipe 32 true {Rua do Endereço que é um logadouro 10 São Paulo SP}}
Nome: Phellipe, Idade: 32, Ativo: true
```
