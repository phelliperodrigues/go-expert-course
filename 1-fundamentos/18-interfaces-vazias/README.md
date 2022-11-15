# Capitulo Fundamento
## Aula 18 - Interfaces Vazias

Podemos trabalhar com interfaces vazias para trabalhar com mais liberdade referente aos tipos

Porem devemos trabalhar com muita cautela ao utilizar isso, por o `GO` tem como objetivo ser fortemente tipado, mas quando utilizamos interfaces vazias quebramos isso em `GO`

Exemplo:
```go
package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
```
output:
```shell
O tipo da variavel é int e o valor é 10
O tipo da variavel é string e o valor é Hello, World!
```
Podemos observar que independente do valor que passamos para a variaveld o tipo `interface{}` ela é aceita