# Capitulo Fundamento
## Aula 2 - Declaração e atribuição

Quando declaramos uma variavel do tipo constante(`const`) siginifica que o valor da variavel não pode ser alterado.

Podemos declarar variaveis de varias forma

```go
package main

const a = "Hello, World!"

var b bool
var c = "c"
var (
	d string
	e int
)

func main() {
	println(a)
	println(b)
	b = true
	println(b)
	println(c)
	println(d)
	println(e)
	var f string
	println(f)
}
```

Variaveis de `a` até `e` são variaveis de escopo global
Variavel `f` é de escopo local

Lembrando que `var` eu consigo alterar o valor e `const` é um valor constante

Tudo no `GO` é fortemente tipado

Não conseguimos pegar a variavel `b` e atribuir o valor `10` por exemplo, erro:

```bash
cannot use 10 (type untyped int) as type bool in assignment
```

Em `GO` toda variavel declarada dentro de um escopo local deve ser utilizada, exemplo:

```go
package main


func main() {
	var helloWorld string = "Hello, World!"
	var hello string = "Hello, Phellipe!"
	println(hello)

}
```
Erro:
```bash
./main.go:22:2: helloWorld declared but not used
```

Existe uma forma reduzida de declarar variavel _short hand_ 

Utilizando a notação `:=` já conseguimos criar a variavel e atribuir o valor da mesma, e o `GO` por sua vez declara a variável com o tipo correto

```go
package main


func main() {
	hello := "Hello, Phellipe!"
	println(hello)

}
```
Uma atenção do exemplo acima o `:=` é utilizado somente na primeira vez declaração/atribuição da variavel, depois que vc declarou a variavel, para mudar o valora é só utilizar o `=`