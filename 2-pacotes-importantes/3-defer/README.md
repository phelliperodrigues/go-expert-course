# Capitulo: Pacotes Importantes
## Aula 3: Defer

Podemos definir `defer` como algo que atrasa a execução de algo.

Vamos utilizar o exemplo da aula 2, requisição http.
```go
package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

	req.Body.Close()
}
```
No código acima é muito facil nós implementarmos uma logica e esquecermos de executar algo no final da função, no nosso exemplo o fechamento da conexão `Body` `req.Body.Close()`

Então para evitarmos o esquecimento da execução de algo no final da função podemos alterar nosso codigo para a seguinte instrução

```go
package main

import (
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
}
```

O `defer` é um statiment, oque ele faz é atrazar a execução, então basicamente, quando nosso programa encontra a notação `defer` antes de uma instrução, nosso programa pula a execução dele e deixa para executar somente no final da função, isso nos ajuda a organizar nosso codigo, no nosso exemplo, nao podemos exequer de fechar a conexão, entao assim que abrimos a mesma, testando se não há erro ja passamos a instrução dizendo pro `GO` fehca minha conexão assim que acabar de executar tudo na função.

Vamos fazer um exemplo mais simples

```go
package main

import "fmt"

func main() {
	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
```
output:
```shell
Segunda linha
Terceira linha
Primeira linha
```

Podemos observar que mesmo que adicionamos um `prinln` no inicio da função ele sera executado somente no final, antes da finalização da função.
