# Capitulo: Pacotes Importantes
## Aula 7: Manipulando Headers

Nesta aula vamos aproveitar parte do codigo da aula 6 e vamos adicionar algumas manipulações nos HEADERS do nosso serviço.

Basicamente vamos adicionar duas validações, uma para validar se algo é diferente na rota `/` vamos devolver um status code `404` e uma validação de query string que vai validar se existe o o parametro `cep` caso exista retornará `Hello World` caso não exista retornará `400 - Bad Request`

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

Nossa alteração ficou na função `BuscarCEPHandler` 
- Adicionamos um `if` para validar se o `PATH` da url requisitada é diferente de `/` caso seja adicionamos uma manipulação no header com a dunção `w.WriteHeader(http.StatusNotFound)` e retornando o serviço com status `404`.
- Adicionamos também a busca da quey string `cepParam := r.URL.Query().Get("cep")` caso ele não encontre devolvemos status `400` caso contrario ele continua o processo.
- Adicionamos um header de `Content-Type` `w.Header().Set("Content-Type", "application/json")` e também adicionamos o status code 200 `w.WriteHeader(http.StatusOK)`

