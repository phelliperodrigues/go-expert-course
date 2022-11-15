# Capitulo Fundamento
## Aula 17 - Ponteiros e Structs

No exemplo desta aula vamos trabalhar com ponteiro e strucs.

```go
package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c Cliente) andou() {
	c.Nome = "Phellipe Rodrigues"
	fmt.Printf("O cliente %v andou\n", c.Nome)
}
func main() {
	phellipe := Cliente{Nome: "Phellipe"}
	phellipe.andou()
	fmt.Println("Nome:", phellipe.Nome)
}

```
output:
```shell
O cliente Phellipe Rodrigues andou
Nome: Phellipe
```
Podemos observar que mesmo que dentro do metodo `andou()` eu esteja alterando o valor do atributo `Nome` do cliente, ao imprimir o nome dentro da função `main` o nome não sofreu alteração.

Vamos corrigir isso:
```go
package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) andou() {
	c.Nome = "Phellipe Rodrigues"
	fmt.Printf("O cliente %v andou\n", c.Nome)
}
func main() {
	phellipe := Cliente{Nome: "Phellipe"}
	phellipe.andou()
	fmt.Println("Nome:", phellipe.Nome)
}
```
output:
```shell
O cliente Phellipe Rodrigues andou
Nome: Phellipe Rodrigues
```
Agora passamos a receber o ponteiro de memoria do cliente como reciver da função `func (c *Cliente) andou()` assim toda alteração feita no cliente refletira diretamente no espaço de memoria da variavel que executou o metodo


Vamos utilizar em um caso mais real, o exemplo é uma simulação de emprestimo em uma conta, onde a simulação não deve alterar o saldo da conta até ser contemplada.
```go
package main

import "fmt"

type Conta struct {
	saldo float64
}

func (c Conta) simular(valor float64) float64 {
	println("Simulando")
	c.saldo += valor
	return c.saldo
}

func (c *Conta) contemplarSimulacao(valorSimulacao float64) {
	println("Aplicando emprestimo")
	c.saldo = valorSimulacao
}
func main() {
	conta := Conta{saldo: 0.0}
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	valorSimulacao := conta.simular(20000)
	fmt.Printf("O valor total da sua conta com a simulação é %.2f\n", valorSimulacao)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	conta.contemplarSimulacao(valorSimulacao)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)

}
```
output:
```shell
O saldo de sua conta é 0.00
Simulando
O valor total da sua conta com a simulação é 20000.00
O saldo de sua conta é 0.00
Aplicando emprestimo
O saldo de sua conta é 20000.00
```
Vamos que o metodo `simular()`, somente soma o valor do saldo da conta com o valor a emprestar, porem não altera o valor original da conta, mas quando passamos o valor simulado para o metodo `contemplarSimulacao()` o valor do saldo é alterado com sucesso

Vamos criar uma função para criar uma conta sempre vazia.

```go
package main

import "fmt"

type Conta struct {
	saldo float64
}

func NewConta() *Conta {
	return &Conta{saldo: 0.0}
}

func (c Conta) simular(valor float64) float64 {
	println("Simulando")
	c.saldo += valor
	return c.saldo
}

func (c *Conta) contemplarSimulacao(valorSimulacao float64) {
	println("Aplicando emprestimo")
	c.saldo = valorSimulacao
}
func main() {
	conta := NewConta()
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	valor := conta.simular(20000)
	fmt.Printf("O valor total da sua conta com a simulação é %.2f\n", valor)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	conta.contemplarSimulacao(valor)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)

}
```
output:
```shell
O saldo de sua conta é 0.00
Simulando
O valor total da sua conta com a simulação é 20000.00
O saldo de sua conta é 0.00
Aplicando emprestimo
O saldo de sua conta é 20000.00
```
Podemos observar a função criada `NewConta()` a mesma cria uma Conta retornando o endereço de memoria, isso nos ajuda e simplifica muito na criação de structs padronizadas.