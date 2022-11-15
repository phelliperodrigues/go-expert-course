# Capitulo Fundamento
## Aula 21 - Pacotes e módulos parte 1 

Vamos começar a estudar pacotes em `GO`

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
	"matematica"
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
pacote.go:5:2: package is not in GOROOT (/usr/local/go/src/matematica)
```
Podemos observar no erro acima que o `GO` não encontrou o pacote na pasta do GOROOT

Para resolver isso utilizaremos os módulos do `GO`

> Continua na aula 2....
