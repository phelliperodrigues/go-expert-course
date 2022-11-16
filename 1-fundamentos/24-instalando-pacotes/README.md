# Capitulo Fundamento
## Aula 24 - Instalando Pacotes

Algumas versões anteriores do `GO` quando era preciso uma dependencia externa era preciso executar o comando `go get <dependencia>` e ele baixava para poder usar no projeto, porem não havia como fixar uma versão ou algo pareciso no projeto, ele sempre pegar a ultima versão e deixava disponivel para usar.

Recentemento quando executamos o `go get ...` sem um modulo criado acontece o seguinte.
```shell
go get golang.org/x/exp/constraints
```
```shell
go: go.mod file not found in current directory or any parent directory.                                                                                                                                                                                                                                               ─╯
        'go get' is no longer supported outside a module.
        To build and install a command, use 'go install' with a version,
        like 'go install example.com/cmd@latest'
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.

```
Vamos criar um modulo

```shell
go mod init curso-go
```
output:
```shell
go: creating new go.mod: module curso-go
```
Percebemos que foi criado o arquivo `go.mod`

```go
module curso-go

go 1.19
```
Agora vamos executar novamente o `get`
```shell
go get golang.org/x/exp/constraints
```

output:
```shell
go: downloading golang.org/x/exp v0.0.0-20221114191408-850992195362                                                                                                                                                                                                                                                   ─╯
go: added golang.org/x/exp v0.0.0-20221114191408-850992195362
```

Podemos observar que nosso arquivo `go.mod` ficou da seguinte maneira
```go
module curso-go

go 1.19

require golang.org/x/exp v0.0.0-20221114191408-850992195362 // indirect
```
E também podemos ver que o go gerou um arquivo `go.sum`
```go
golang.org/x/exp v0.0.0-20221114191408-850992195362 h1:NoHlPRbyl1VFI6FjwHtPQCN7wAMXI6cKcqrmXhOOfBQ=
golang.org/x/exp v0.0.0-20221114191408-850992195362/go.mod h1:CxIveKay+FTh1D0yPZemJVgC/95VzuuOLq5Qi4xnoYc=
```
o `go.sum` é como se fosse um `package.lock` do node, tem como função garantir que sempre vai estar com a mesma versão.

Existe um comando muito usado também que é `go mod tidy` este comando é utilizado para organizar nossas dependencias, caso vc tenha baixado o projeto mas não tenha baixado a dependencia o mesmo executará esta ação e caso tenha alguma dependencia no `go.mod` que não esta em uso no projeto ele removerá

Vamos realizar um teste com uma outra dependencia no caso o pacote de `uuid` do google [link](https://github.com/google/uuid)

Vamos criar o arquivo `main.go` e neste arquivo, vamos colocar o import da dependencia, sem intala-la com `go get...`

```go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Meu UUID", uuid.New())
}
```
output:
```shell
main.go:5:2: no required module provides package github.com/google/uuid; to add it:
	go get github.com/google/uuid
```

Podemos observar que o `GO` não conseguiu encontrar a dependencia, vamos executar agora o comando
```shell
go mod tidy
```
output:
```shell
go: finding module for package github.com/google/uuid
go: found github.com/google/uuid in github.com/google/uuid v1.3.0
```

Vamos executar novamente o programa:
output:
```shell
Meu UUID 57fb638c-3b23-454a-ac95-4cc8fca77ce7
```

### RESUMO
`go get ...` Adiciona no `go.mod` a dependencia

`go mod tidy` Otimiza os modulos usados no projeto, baixando ou removendo do `go.mod`