# Capitulo Fundamento
## Aula 5 - Percorrendo Array

Alem dos tipos primitivos como `string`, `int`, `float` existe alguns tipos compostos, nesta aula vamos trabalhar com `array`

`Array` trabalha como uma variavel, ele contem um tamanho fixo e podemos percorrer os valores

```go
package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(meuArray[len(meuArray)-1]) // Valor do ultimo index
	fmt.Println(len(meuArray) -1) //Ultimo Index do meu array
}
```
No exemplo acima a variavel `meuArray` contem 3 posição, isso significa que NÃO conseguimos adicionar mais de 3 atributos dentro do array<br>
Caso tentamos inserir um quarto valor, o compilador apresentara o erro abaixo:
```shell
./main.go:10:10: invalid array index 3 (out of bounds for 3-element array)
```

Podemos percorrer um array utilizando `for`

```go
package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	for index, value := range meuArray {
		fmt.Printf("O valor do indece é %d e o valor é %d\n", index, value)
    }
}
```
output:
```shell
O valor do indece é 0 e o valor é 1
O valor do indece é 1 e o valor é 2
O valor do indece é 2 e o valor é 3
```