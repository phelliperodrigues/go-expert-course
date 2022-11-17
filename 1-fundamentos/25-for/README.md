# Capitulo Fundamento
## Aula 25 - For

`GO` para sinplificar a forma de trabalhar com loop adicionou somente o `for`, diferente de outras linguagens que tem `while`, `do while`, `foreach`, `maps`, `interations`etc.

Temos o for interativo como o exemplo abaixo:
```go
package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
```
output:
```shell
0
1
2
3
4
5
6
7
8
9
```

Também temos o `for` com `range`, muito usado para percorrer `array` e `slice`
```go
package main

func main() {
	numeros := []string{"um", "dois", "três"}
	for index, value := range numeros {
		println(index, value)
	}
}
```
output:
```shell
0 um
1 dois
2 três
```

Podemos ignorar o `index` ou até mesmo o `valor` no range utiliando o `blank identifier`
```go
package main

func main() {
	numeros := []string{"um", "dois", "três"}

	//ignorando index
	for _, value := range numeros {
		println(value)
	}
	
	//ignorando valor
	for index, _ := range numeros {
		println(index)
	}
}

```
output:
```shell
um
dois
três
0
1
2
```

Tem a possibilidade de criar um `for` condicional (Como se fosse o `while`)
```go
package main

func main() {
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
```
output:
```shell
0
1
2
3
4
5
6
7
8
9
```

Também tem a possibilidade de um `for` infinito
```go
package main

func main() {
	
	for {
		println("Hello, World!!")
	}
}
```
output:
```shell
Hello, World!!
Hello, World!!
Hello, World!!
Hello, World!!
Hello, World!!
Hello, World!!
Hello, World!!
........ INFINITO
```