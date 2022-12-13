# Capitulo: Pacotes Importantes
## Aula 6: Iniciando com HTTP

Nesta aula vamos começar a trabalhar com serviços HTTP, Vamos começar a preparar uma API para realizar a busca de CEP via HTTP.

Vamos criar uma função que devolve um `Hello, World` 

```go
package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscarCEP)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func BuscarCEP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
```

Criamos uma função `BuscaCEP` que tem como assinatura dois parametros `http.ResponseWriter` e `*http.Request` quando fazemos isso estamos implementando uma interface, nessa função por enquanto estamos somente dizendo para o response devolver uma mensagem `Hello World`

Dentro da função `main` temos a chamada da função `http.HandleFunc` no primeiro parametro passamos `"/"` isso é o mapeamento da roda raiz e o segundo parametro é a nossa função `BuscaCEP`, então o que significa que toda requisição que bater na rota `/` irá chamar a função `BuscarCEP`

Também temos a chamada da função `http.ListenAndServe` esta função segura nossa aplicação de pé levantando um servidor escutando na porta `8080`