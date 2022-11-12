# Capitulo Fundamento
## Aula 1 - Entendendo a primeira linha

- Todo arquivo `.go` por padrão vai ter um `package` no inicio do arquivo.
- O nome do package deve ser sempre o nome do diretório onde se encontra o arquivo `.go` com exceção do package `main`.
- O `package main` é o ponto de entrada da nossa aplicação.
- A função `main()` é a função que inicia a aplicação go

- Quando temos mais arquivos dentro do mesmo diretorio, os arquivos precisão ter a mesma `package`

Exemplo de erro quando criar no mesmo diretorio arquivos com pacotes diferentes:

```bash
found packages soma (a.go) and main (main.go) in /go-expert-course/1-fundamentos/1-estrutura 
```
Quando utilizamos o mesmo pacote em arquivos diferentes, podemos "compartilhar" variaveis e funções entre arquivos.