# Capitulo Fundamento
## Aula 19 - Type assertation

Quando executamos algo com uma interface vazia, no final das cotas o `GO` vai querer saber qual foi o tipo que vc atribuiu 

Veja o exemplo abaixo:
```go
package main

func main() {
	var minhaVar interface{} = "Phellipe"

	println(minhaVar)
}
```
output:
```shell
(0x466460,0x48a278)
```
Se observarmos a saida do console, vemos que o `GO` imprimiu uma referencia mas não o valor da variavel `minhaVar` isso por que não inferimos o tipo da variavel.

Para corrigir isso precisamos informar ao `GO` que a variável `minhaVar` é do tipo string.
```go
package main

func main() {
	var minhaVar interface{} = "Phellipe"

	println(minhaVar.(string))
}
```
output:
```shell
Phellipe
```
Isso funciona perfeitamente, mas temos que ter cautela caso a aferição do tipo esteja errada, como se no print eu tentasse colocar como `int64`

Ocasionaria o seguinte erro:
```shell
panic: interface conversion: interface {} is string, not int64

goroutine 1 [running]:
main.main()
        ~/go-expert-course/1-fundamentos/19-type-assertation/interfaces.go:6 +0x45
```
Porem o `GO` tras uma forma de validarmos se a aferição de tipo deu certo ou não

Exemplo
```go
package main

import "fmt"

func main() {
	var minhaVar interface{} = "Phellipe"

	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é: %v e o valor de ok é: %t", res, ok)

}
```
output:
```shell
O valor de res é: 0 e o valor de ok é: false
```
Com isso conseguimos validar se a conversao foi feita com sucesso.