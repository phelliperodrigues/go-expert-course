# Capitulo: Pacotes Importantes
## Aula 8: Criando função para BuscaCEP

Nesta aula vamos implementar a função de BuscaCep para devolver o edereço encontrato na API do ViaCep

```go

func BuscaCep(cep string) (*Endereco, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var endereco Endereco

	err = json.Unmarshal(body, &endereco)
	if err != nil {
		return nil, err
	}
	return &endereco, nil
}
```
Toda explicação de como funciona esta função podemos encontrar na *Aula 5: Busca Cep*
