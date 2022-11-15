# Capitulo Fundamento
## Aula 15 - Ponteiros

Trabalhar com ponteiro no `GO` é bem simples

Quando criamos uma variavel em `GO` como `a := 10` o `GO` abre uma "caixinha" e guarda o valor `10` dentro e a variavel `a` guarda o endereço dessa caixinha.

Veja o exemplo abaixo:

```go
package main

func main() {
	a := 10
	println(a)
	println(&a)
}
```
output:
```shell
10
0xc00003e770
```
Quando executamos a função `println()` somente com a variavel `a` o `GO`pega o valor guardado dentro do endereço de memoria e printa na tela, porem quando colocando a notação `&` antes da variavel `a` estamos solicitando para o `GO` imprimir o endereço de memória onde que esta guardado na variavel `a`.

Então toda vez que criamos um ponteiro, o ponteiro É o endereçamento na memória

```go
package main

func main() {
	a := 10
	println(a)
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)
	*ponteiro = 20
	println(a)

}
```
```shell
10
0xc00003e770
0xc00003e770
20
```
Podemos ver que criamos uma variavel `ponteiro` com a notação `*int` recebendo o valor `&a` esta varaivel é um ponteiro de memória, então quando mudamos o valor da variavel `ponteiro` para 20 o valor de `a` que esta apontamos para o mesmo local da memoria também passa a ser 20.

Vamos criar mais uma variavel `b` apontando para o ponteiro de memoria de `a` e tentar ver o valor da variavel `b`

```go
package main

func main() {
	a := 10
	println(a)
	println(&a)

	b := &a
	println(b)
	println(*b)
}
```
output
```shell
10
0xc00003e770
0xc00003e770
10
```
Observamos que para acessar/alterar o valor do ponteiro `b` precisamos utilizar a notação `*` antes da variavel, isso é chamado de _reference_
