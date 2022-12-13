# Capitulo: Pacotes Importantes
## Aula 9: Finalizando resposta para o servidor HTTP

Nesta aula vamos finalizar nosso serviço de busca de CEP

Nossa alteração sera apenas na função `BuscarCEPHandler`

Antes:
```go
func BuscarCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
```

Vamos remover o trecho de codigo
```go
_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
```

E adicionar os seguinte codigo
```go
endereco, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	
	err = json.NewEncoder(w).Encode(endereco)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
```

Resultado:

```go
func BuscarCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	endereco, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	
	err = json.NewEncoder(w).Encode(endereco)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

```
Nesta implementação estamos chamando nossa função `BuscaCep` passando o parametro que recebemos no query string `cepParam` depois convertemnos o retorno que esta na variavel `endereco` utilizando o `Encoder` na sequinte instrução `json.NewEncoder(w).Encode(endereco)` que por si só ja converte a variavel `endereco` que é do tipo `Endereco` struct e agrela ele no `Writer` do nosso response. 