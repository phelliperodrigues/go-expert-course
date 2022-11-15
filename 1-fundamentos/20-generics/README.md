# Capitulo Fundamento
## Aula 20 - Generics

Vamos observar o seguinte exemplo:

```go
package main

import "fmt"

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	println(SomaInteiro(m))
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	fmt.Printf("%.2f",SomaFloat(m2))

}
```
output:
```shell
31000
31000.00
```
No exemplo acima vamos que para somar 2 tipos numericos diferentes `int` e `float` temos que criar duas funçoes especificas, com generics podemos simplicar isso, veja no exemplos:
```go
package main

import "fmt"

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	fmt.Printf("Soma m: %d\n", Soma(m))
	fmt.Printf("Soma m2: %.2f\n", Soma(m2))
}
```
output:
```shell
Soma m: 31000
Soma m2: 31000.00
```
Podemos observar que a função `Soma` tem uma notação diferente `[T int | float64]` isso indica que os atributos podem ser do tipo `int` ou `float` e o retorno também

Podemos também criar uma `constrant`
```go
package main

import "fmt"

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}

	fmt.Printf("Soma m: %d\n", Soma(m))
	fmt.Printf("Soma m2: %.2f\n", Soma(m2))
}
```
output:
```shell
Soma m: 31000
Soma m2: 31000.00
```
Podemos observar a interface `Number` possui os tipos `int` e `float64` e nosso generics passa a receber essa constrant na assinatura do metodo

Também a uma possibilidade de criarmos nosso proprio tipo de numero, por exemplo o MyNumber
```go
package main

import "fmt"

type MyNumber int 

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m3 := map[string]MyNumber{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}

	fmt.Printf("Soma m: %d\n", Soma(m))
	fmt.Printf("Soma m2: %.2f\n", Soma(m2))
	fmt.Printf("Soma m3: %d\n", Soma(m3))
}
```
output:
```shell
# command-line-arguments
./generics.go:46:36: MyNumber does not implement Number (possibly missing ~ for int in constraint Number)
```
Mas o `GO` por ser fortemente tipado não consegue entender que o MyNumber é um `int` porem ele sugere que vc utilize a notação `~` na frente do int, veja abaixo:
```go
package main

import "fmt"

type MyNumber int 

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}
func main() {
	m := map[string]int{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m2 := map[string]float64{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}
	m3 := map[string]MyNumber{"Phellipe": 10000, "Patricia": 20000, "Ping": 1000}

	fmt.Printf("Soma m: %d\n", Soma(m))
	fmt.Printf("Soma m2: %.2f\n", Soma(m2))
	fmt.Printf("Soma m3: %d\n", Soma(m3))
}
```
output:
```shell
SomaFloat: 31000.00
Soma m: 31000
Soma m2: 31000.00
Soma m3: 31000
```
O `~` significa que o `GO` vai entender a relação entre `MyNumber` e `int` 

Vamos adicionar mais uma logica com generics, criaremos a função `Compara()`
```go
package main

import "fmt"

type Number interface {
	~int | float64
}

func Compara[T Number](a, b T) bool {
	return a == b
}
func main() {
	fmt.Println(Compara(10, 10.0))
}
```
output:
```shell
# command-line-arguments
./generics.go:52:26: default type float64 of 10.0 does not match inferred type int for T
```
Podemos ver que não conseguimos comparar por os tipos de informações são diferentes, estou estou tentando caparar `int` com `float`

Para corrigir isso podemos utilizar uma notação `any`, isso significa que pode ser qualquer tipo. Porem isso ainda não resolver nosso problema, por que o compilador vai dar o seguinte erro
```shell
# command-line-arguments
./generics.go:39:9: invalid operation: a == b (incomparable types in type set)
./generics.go:52:26: default type float64 of 10.0 does not match inferred type int for T
```
Então podemos utilizar a seguinte contrant `comparable` porem não conseguimos ainda comprar tipos diferentes
```shell
# command-line-arguments
./generics.go:52:26: default type float64 of 10.0 does not match inferred type int for T
```
Mas como resolver o problema de comparação? 

Existe um pacote chamado constraints, para utilizar isso vamos ver na aula de pacotes. ......


#### _Recurso implemetado a partidar 1.18 do `GO`_ 