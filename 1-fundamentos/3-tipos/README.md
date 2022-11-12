# Capitulo Fundamento
## Aula 3 - Criação de tipos

Em `GO` podemos criar o nosso tipo

```go
package main

type ID int

var id ID
func main() {
    println(id)
}
```

No exemplo acima por uma questão semantica criamos nosso tipo `ID` e declaramos uma variavel `id` do tipo `ID`