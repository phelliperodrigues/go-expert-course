# Capitulo Fundamento
## Aula 4 - Importando fmt e tipagem

O `GO` tem um pacote de formatação `fmt`

Da mesma forma das variaveis locais, um import se não utilizado da erro de compilação

quando utilizamos a lib `fmt` podemos formatar algumas coisas, no exemplo desta aula vamos utilizar para imprimir o tipo de uma variável

```go
package main

import "fmt"

var e float64 = 1.2

func main() {

	fmt.Printf("O tipo de E é %T", e)
}
```

O exemplo acima deve imprimir o console a seguinte instrução `O tipo de E é float64` 

a instrução `fmt.Printf` nos permite formatar o texto a ser impresso utilizando alguns modificadores, por exemplo

`fmt.Printf("%T¨, e)` imprime o tipo da variavel `e`
`fmt.Printf("%f", e)` imprime o valor da variavel `e` se ela for do tipo `float`
`fmt.Printf("%s", e)` imprime o valor da variavel `e` se ela for do tipo `string`
`fmt.Printf("%v", e)` imprime o valor da variavel `e` independente do tipo