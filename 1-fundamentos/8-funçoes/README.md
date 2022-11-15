# Capitulo Fundamento
## Aula 8 - Funções

Funções são extramamante importante e estão presente nas maiorias das linguagens de programação

No exemplo abaixo podemos ver que recebemos 2 atributos do tipo inteiro, somamos os dois e retornamos um valor do tipo inteiro também

```go
package main

import "fmt"

func main() {
	soma := sum(1, 2)
	fmt.Println(soma)
}

func sum(a int, b int) int {
	return a + b
}
```
output:
```shell
3
```
Podemos observar que tanto o valor do atributo `a` e o atributo `b` da função tem o mesmo tipo, podemos refatorar a função para:

```go
func sum(a , b int) int {
	return a + b
}
```

Em `GO` podemos retornar mais de 1 valor na função, no exemplo abaixo vamos retornar junto a soma um valor boolean dizendo se o valor da somna é maior ou igual a 50

```go
package main

import "fmt"

func main() {
	soma, verdadeiro := sum(10, 50)
	fmt.Printf("O valor somando é maior ou igual 50? %v, o valor é: %d\n", verdadeiro, soma)
	soma, verdadeiro = sum(10, 10)
	fmt.Printf("O valor somando é maior ou igual 50? %v, o valor é: %d\n", verdadeiro, soma)
}

func sum(a, b int) (int, bool) {
	soma := a + b
	return soma, soma >= 50

}
```
output:
```shell
O valor somando é maior ou igual 50? true, o valor é: 60
O valor somando é maior ou igual 50? false, o valor é: 20
```

Por que utilizaremos funções com mais de 1 retorno no `GO`?

Utilizamos bastante funções com mais de 2 tipos de retorno para retornar ERROS 

Exemplo:
```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	soma, err := sum(10, 50)
	if err != nil {
		fmt.Printf("Algo deu errado: \"%v\", valor: %d\n", err, soma)
	}

	soma, _ = sum(10, 10)
	fmt.Printf("Valor da soma: %d\n", soma)
}

func sum(a, b int) (int, error) {
	soma := a + b
	if soma >= 50 {
		return 0, errors.New("Soma é maior que 50")
	}
	return soma, nil
}
```
output:
```shell
Algo deu errado: "Soma é maior que 50", valor: 0
Valor da soma: 20
```
