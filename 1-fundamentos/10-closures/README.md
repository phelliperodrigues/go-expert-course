# Capitulo Fundamento
## Aula 10 - Closures

`GO` tem funçoes anonimas, ou closures, isso significa que dentro dentro de uma função você pode ter outra função.

Exemplo:
```go
package main

import "fmt"

func main() {
	somaMultiplicada := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) * 2
	}()

	fmt.Println("Total multiplicado:", somaMultiplicada)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

```
output:
```shell
Total multiplicado: 110
```
Podemos observar que criamos uma função anonima dentro da função `main()` que chama a função `sum()` e multiplica o resuldato da mesma por 2

