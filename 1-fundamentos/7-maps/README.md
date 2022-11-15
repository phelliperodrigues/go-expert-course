# Capitulo Fundamento
## Aula 7 - Maps

Map é uma coleção de chave e valor e não possui uma ordenação por padrão

```go
package main

import "fmt"

func main() {
	salarios := map[string]int{"Phellipe": 1000, "João": 2000, "Pedro": 3000}
	fmt.Println(salarios)
	fmt.Println("Salario:", salarios["Phellipe"])
}
```
output:
```shell
map[João:2000 Pedro:3000 Phellipe:1000]
Salario: 1000
```

Podemos deletar um item o `map` da seguinte forma

```go
delete(salarios, "Phellipe")
fmt.Println(salarios)
```
output:
```shell
    map[João:2000 Pedro:3000]
```
Também podemos adicionar um novo registro ao nosso `map`

```go
salarios["Phellipe"] = 20000
fmt.Println("Salario:", salarios["Phellipe"])
```
output:
```shell
map[João:2000 Pedro:3000 Phellipe:20000]
Salario: 20000
```

Podemos inicializar um `map` utilizando a função `make()`
```go
novoSalarios := make(map[string]int)
novoSalarios["Phellipe"] = 25000
fmt.Println(novoSalarios)
```
output:
```shell
map[Phellipe:25000]
```

Podemos percorrer também o map utilizando `for`

```go
for nome, salario := range salarios {
    fmt.Printf("O salario de %s é de %d\n", nome, salario)
}
```
output:
```shell
O salario de Phellipe é de 20000
O salario de João é de 2000
O salario de Pedro é de 3000
```
No `for` também podemos ignorar utilizando `blank identifier`, para isso basta utilizar um `_` no lugar da variavel, no exemplo abaixo removemos o nome

```go
for _, salario := range salarios {
    fmt.Printf("O salario é de %d\n", salario)
}
```
output:
```shell
O salario é de 20000
O salario é de 2000
O salario é de 3000
```