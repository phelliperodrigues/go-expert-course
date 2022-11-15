# Capitulo Fundamento
## Aula 11 - Iniciando com Structs

`GO` não é uma linguagem orientada a objeto, de acordo com o time que criou a linguagem `GO` go tem o jeito `go way` de fazer as coisas

Mas existe um recurso que é bastante simular a algumas coisas da orientação a objeto, uma delas são as `structs`

Podemos fazer uma analogo de uma `struct` como se fosse uma `classe` porem devemos ter bem claro que *NÃO* é uma classe

##### Caso de Uso? 
Vamos imaginar que vamos falar algo de um cliente, então queremos ter dados de um cliente para conseguir desenvolver nosso software, então devemos criar uma estrutura para conseguir adicionar todos estes dados do cliente em um só lugar, a ideia da `struct` é exatamente essa.

Exemplo:
```go
package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	phellipe := Cliente{
		Nome:  "Phellipe",
		Idade: 32,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
}
```
output:
```shell
Nome: Phellipe, Idade: 32, Ativo: true
```

Podemos observar acima que criamos a struct `Cliente` com três atributos, e a nossa função `main` criamos uma variavel do tipo `Cliente` passando os atributos, e para printar os valores eu simplemente utilizo a variavel criara ponto o atributo que desejo (`phellipe.Nome`)

Podemos também mudar valores de atibutos que estão na struct, podemos adicionar as linhas de codigo ao nosso exemplo
```go
phellipe.Ativo = false
fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", phellipe.Nome, phellipe.Idade, phellipe.Ativo)
```
output:
```shell
Nome: Phellipe, Idade: 32, Ativo: false
```