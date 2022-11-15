# Capitulo Fundamento
## Aula 16 - Quando usar ponteiro

Quando passamos uma variavel como um atributo de uma função, temos que ter em mente de que o `GO` ira clonar o valor da minha variavel dentro da função.

Vejamos o exemplo abaixo:
```go
package main

import "fmt"

func soma(a, b int) int {
	a = 50
	return a + b
}
func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println("Soma:", soma(minhaVar1, minhaVar2))
	fmt.Println("minhaVar1:",minhaVar1)
}
```
output:
```shell
Soma: 70
minhaVar1: 10
```
Podemos observar que mesmo alterando o valor da variavel `a` dentro da função `soma()` o valor da minha variavel `minhaVar1` continua sem alteração

Para validar isso vamos imprimir o endereço de memoria das variaveis `a` e `minhaVar1`

```go
package main

import "fmt"

func soma(a, b int) int {
	println("a:", &a)
	a = 50
	return a + b
}
func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println("Soma:", soma(minhaVar1, minhaVar2))
	fmt.Println("minhaVar1:", minhaVar1)
	println("minhaVar1:", &minhaVar1)

}
```
output:
```shell
a: 0xc00009af30
Soma: 70
minhaVar1: 10
minhaVar1: 0xc00009af28
```
Agora vamos ver como alterar o valor da variavel `minhaVar1` dentro da função `soma()`
```go
package main

import "fmt"

func soma(a, b *int) int {
	println("a:", a)
	*a = 50
	return *a + *b
}
func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println("Soma:", soma(&minhaVar1, &minhaVar2))
	fmt.Println("minhaVar1:", minhaVar1)
	println("minhaVar1:", &minhaVar1)

}
```
output:
```shell
a: 0xc00006ef30
Soma: 70
minhaVar1: 50
minhaVar1: 0xc00006ef30
```
Para alterar o valor da variavel passada como atributo de função, devemos receber como atributo ponteiros, no codigo acima podemos ver isso em `func soma(a, b *int) int` esta notação `*int` dentro da assinatura da função siginifica que agora não vou receber apenas o valor mas sim o endereço de memoria da variavel que vou receber como atributo, e quando eu altero este valor da variavel `a` estou automaticamente alterando o valor da `minhaVar1` pois as duas variaveis estao apontando para o mesmo endereço de memoria.

##### Quando não utilizar ponteiro?
Quando queremos somente fazer uma copia dos dados, processar e retornar algum valor sem alterar o valor inicial da variavel que passei como parametro.