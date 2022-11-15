# Capitulo Fundamento
## Aula 9 - Funções Variáticas

Em `GO` podemos ter uma função do tipo variática, que é muito utilizada, no exemplo abaixo queremos somar o total de todos os numeros que passamos para a função `sum()`

```go
package main

import "fmt"

func main() {
	soma := sum(1, 2, 3, 4, 5)
	fmt.Println(soma)
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
Soma Total: 15
```
