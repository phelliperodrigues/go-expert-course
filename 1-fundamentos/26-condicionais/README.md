# Capitulo Fundamento
## Aula 26 - Condicionais

Como toda linguagem `GO` implementa condicionais

### Condições
- `==` Igual
- `!=` Diferente
- `>` Maior
- `<` Menor
- `<=` Menor ou Igual
- `>=` Maior ou Igual
- `||` Ou
- `&&` E

### IF
```go
package main

func main() {
	a := 1
	b := 2
	if a > b {
		println(a)
	} else {
		println(b)
	}
}
````
Em `GO` não tem `else if` somente `if` e `else`

### Switch

```go
package main

func main() {
	a := 1

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	default:
		println("nenhum")

	}
}
```