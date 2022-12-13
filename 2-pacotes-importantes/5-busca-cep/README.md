# Capitulo: Pacotes Importantes
## Aula 5: Busca CEP

Nesta aula vamos implementar um buscador de cep de linha de comando, onde executaremos a aplicação, digitaremos o cep e imprimiremos no terminal as informações do endereço encontrado.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ddd         string `json:"ddd"`
}

func main() {

	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(req.Body)
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler retorno: %v\n", err)
		}

		var endereco Endereco

		err = json.Unmarshal(res, &endereco)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse da resposta: %v\n", err)
		}

		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, " Erro ao criar arquivo: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("%#v", endereco))
		if err != nil {
			fmt.Fprintf(os.Stderr, " Erro gravar arquivo: %v\n", err)
		}

		fmt.Printf("%#v", endereco)
	}

}
```
input:
```bash
./cep 04686001
```
output:
```shell
main.Endereco{Cep:"04686-001", Logradouro:"Avenida Nossa Senhora do Sabará", Complemento:"de 768 a 1630 - lado par", Bairro:"Vila Isa", Localidade:"São Paulo", Uf:"SP", Ddd:"11"}
```

Este é o primeiro programa que fizemos que é realmente util, que tem uma aplicação em um cenario real, vou explicar um pouco do que foi feito.

No trecho de código
```go
for _, cep := range os.Args[1:] {
....
}
```
Estamos passando uma instrução para o `go` percorrer todos os parametros que passamos ao executar a aplicação no caso quando executamos `go run main.go ola` o valor da variavel `cep` será `ola`