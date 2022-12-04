# Capitulo: Pacotes Importantes
## Aula 4: Trabalhando com json

Vamos entender o processo de serialização de dezerialização de JSON com `GO`.

Para trabalhar com JSON, `GO` tras consigo uma pacote `json` para serializar e deserializar arquivos jsons


#### Gerando JSON
```go
package main

import (
	"encoding/json"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))
}
```
output:
```shell
{"numero":1,"saldo":100}
```
No exemplo acima estamos utilizando a função `Marshal()` para serializar uma struct em um json, esta função retornar um slice de bytes e um erro, como podemos observar no codigo, precisamos converter esse `[]bytes` para `string`

Uma observação importante no código assima são as anotação na strunct como ```json:"numero"``` isso nos permite padronizar os nomes dos atributos do json, também podemos criar validacoes

Também há outra forma de serializar, utilizando a função `Encode()`, veja o exemplo abaixo

```go
package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(conta)
	if err != nil {
		println(err)
	}
}
```
output:
```shell
{"numero":1,"saldo":100}
```
A diferenção entre `Marshal()` e `Encode()` é que a função encode trabalhamos quando queremos passar a serialização para uma outra ação, no nosso exmplo estamos passando para a imprissão no console `os.Stdout` mas poderia ser para uma resposta de um webserver ou uma variável, ou qualquer coisa que faça sentido ser passado o json serializado.

#### Lendo JSON
Também podemos ler um Json e transforma-lo em uma struct, veja o exemplo abaixo
```go
package main

import (
	"encoding/json"
	"fmt"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {

	var contaX Conta
	jsonPuro := []byte(`{"numero":2,"saldo":200}`)
	err := json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	fmt.Printf("%#v\n", contaX)

}
```
Todos json em `go` é um `[]byte` e o pacote `json` nos fornece uma função `Unmarshal` que basicamente converte um `[]byte` em uma struct, mapeando os valores dos atributos para os atrubutos da struct.