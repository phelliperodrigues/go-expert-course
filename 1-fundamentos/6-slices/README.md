# Capitulo Fundamento
## Aula 6 - Slices

Slices são mau compreendidos no `GO`, pois muita gente diz que trabalhamos com array infinitos quando utilizamos slices.

Mas na verdade quando utilizamos slices, estamos trabalhando por "debaixo do capô" com arrays, basicamente os slices são divididos em três partes, um ponteiro para apontar para o array, um ponteiro e a capacidade.

Para diferenciar um slice de um array podemos ver na declaração abaixo
```go
var meuArray = [3]int{}
var meuSlice = []int{}
```
O array por padrão recebe um valor para sua capacidade, enquanto o slice não recebe dizendo que ele pode ser "dinamico"

```go
package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

Com o print abaixo podemos ver que podemos dizer ao slice para que ele "drop" items a partir de uma notação `:`

Podemos observar que ainda temos a capacidade inicial de 5 items mesmo "dropando" todos os itens

```go
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
```
output:
```shell
len=0 cap=5 []
```
Outros exemplos:
```go
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
```
output:
```shell
len=4 cap=5 [2 4 6 8]
```

Agora vamos ver outro caso, vamos utilizar o `index` seguido de `:`

Agora o `GO` vai pegar apenas a partir do segundo index

```go
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
```
output:
```shell
len=3 cap=3 [6 8 10]
```

Um ponto importante para compreendermos, é que como aumento a capacidade do `slice`, pois se a diferença entre um slice e um array é que o slice eu nao preciso definir a capacidade no inicio, como eu aumento a capacidade deste slice, no final das contas o slice sempre vai apontar para um array, e este array sempre vai ter uma capacidade fixa, e se precisarmos de uma capacidade maior ele vai apontar para um array com uma capacidade maior

Vamos adicionar o trecho abaixo em nosso código:

```go
	s = append(s, 12)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
```

```shell
len=2 cap=10 [2 4]
len=6 cap=10 [2 4 6 8 10 12]
```
Podemos ver que a capacidade de nosso slice agora é 10, sendo que no inicio o mesmo slice tinha 5 como capacidade, e no comenta `append` somente adicionamos 1 item apenas

Por baixo dos panos o go, ele vê que toda vez que vc precisa dar um `append` em seu slice e ele nao tem mais capacidade para a adicionar aquele item, ele pega a capacidade inicial e "dobra" no slice, como nosso slice inicial tinha 5 de capacidade e adicionamos um sexto item ele dobrou sua capacidade para 10

Dica: Sempre tente iniciar o slice em um valor proximo ao valor real que você irá trabalhar, por que se vc trabalhar com um valor muito diferente e você precisar redimencionar o slice vai ser um processo muito custoso para o go ficar criando outros arrays.