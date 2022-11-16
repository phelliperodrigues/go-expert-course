# Capitulo Fundamento
## Aula 23 - Pacotes e módulos parte 3 

Veja a seguinte estrutura

```
.
project 
│
└───1-fundamentos
│   │
│   └─── ...
│   │
│   └───21-pacote-1
│   │   │   pacotes.go
│   │   └───matematica
│   │       │   matematica.go
│   │ ...
```

`pacote.go`
```go
package main

import (
	"fmt"
	"github.com/phelliperodrigues/23-pacote-3/matematica"
)

func main() {
	resultado := matematica.soma(10, 20)
	fmt.Printf("Resultado: %v\n", resultado)
}
```

Vamos fazer uma alteração no arquivo `matematica.go` alterando a função de `Soma()` para `soma()`

`matematica/matematica.go`
```go
package matematica

func soma[T int | float64](a, b T) T {
	return a + b
}
```
output: `go run pacote.go`
```shell
# github.com/phelliperodrigues/23-pacote-3
./pacotes.go:9:26: undefined: matematica.Soma
```
Quando executamos o programa acontece o erro acima, isso acontece por que em `GO` uma função/variavel/metodo para exter exportado para fora do pacote que ele pertence o nome precisa começar com uma letra maiuscula, caso contrario será vizivel somente dentro do mesmo pacote.

Vamos voltar a função para `Soma()` para deixa-la publica novamente

Vamos realizar mais alguns testes, ainda no arquivo `matematica.go` vamos criar uma variavel
```go
var A int = 10
```
E vamos adicionar dentro na função `main()` do arquivo `pacote.go` a seguinte linha
```go
fmt.Println(matematica.A)
```
outout:
```shell
Resultado: 30
10
```

Agora vamos alterar a variavel `A` para `a` e vamos altrar o print no arquivo `pacote.go` para `matematica.a`
output:
```shell
# github.com/phelliperodrigues/23-pacote-3
./pacotes.go:11:25: undefined: matematica.a
```

Vamor fazer um exemplo com uma struct também

No arquivo `matemarica.go` vamos criar uma struct `Carro` com um atributo `Marca`, e vamos criar este carro no `pacote.go`
`matematica.go`
```go
type Carro struct {
	Marca string
}
```
`pacote.go` 
```go
carro := matematica.Carro{Marca: "Fiat"}
fmt.Printf("%#v\n", carro)
```
output:
```shell
Resultado: 30
10
matematica.Carro{Marca:"Fiat"}
```
Agora vamos altrar a struct para `carro`

output:
```shell
# github.com/phelliperodrigues/23-pacote-3
./pacotes.go:13:22: undefined: matematica.carro
```
Agora vamos voltar a struct para `Carro` porem vamos alterar o atributo de `Marca` para `marcar`

output:
```shell
# github.com/phelliperodrigues/23-pacote-3
./pacotes.go:13:28: unknown field 'marca' in struct literal of type matematica.Carro
```
Vamos agora testar com uma método para a struct `Carro`

`matematica.go`
```go
func (c Carro) Andar() {
	fmt.Println("Carro Andando....")
}
```
`pacote.go`
```go
	carro.Andar()
```
output:
```shell
Resultado: 30
10
matematica.Carro{Marca:"Fiat"}
Carro Andando....
```
Agora vamos alterar o metodo para `andar()`

output:
```shell
# github.com/phelliperodrigues/23-pacote-3
./pacotes.go:15:8: carro.andar undefined (type matematica.Carro has no field or method andar)
```

#### RESUMINDO: Atributos/Funções/Metodos com o início MAIUSCULO é publico, com MINUSCULO é vizivel somente dentro do mesmo pacote
