# Capitulo Fundamento
## Aula 22 - Pacotes e módulos parte 2 

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

Devemos executar dentro da pasta o seguinte comando
```shell
go mod init github.com/phelliperodrigues/22-pacote-2
```
output:
```shell
go: creating new go.mod: module github.com/phelliperodrigues/22-pacote-2                                                                                                                                                                                                                                              ─╯
go: to add module requirements and sums:
        go mod tidy

```
O Comando vai gerar o seguinte arquivo `go.mod`
```go
module github.com/phelliperodrigues/22-pacote-2

go 1.19

```

Após isso conseguimos importar o pacote matematica a partir do modulo que criamos, o arquivo `pacote.go` fica assim:

`pacote.go`
```go
package main

import (
	"fmt"
	"github.com/phelliperodrigues/22-pacote-2/matematica"
)

func main() {
	resultado := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v", resultado)
}
```

`matematica/matematica.go`
```go
package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}
```
output: `go run pacote.go`
```shell
Resultado: 30
```
A partir da criação do modulo com o comando `go mod init ...` conseguimos criar um modulo proprio e acessar os pacotes existentes dentro deste mesmo módumo.
