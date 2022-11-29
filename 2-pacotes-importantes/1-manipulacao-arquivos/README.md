# Capitulo: Pacotes Importantes
## Aula 1: Manipulação de arquivos


#### Criando um arquivo

Para criarmos um arquivo vamos importar a lib `os` veja o exemplo abaixo:
```go
package main

import "os"

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	f.Close()
}
```
- Criamos o arquivo com a função `os.Create(nome_do_arquivo.extenção)` passando um nome do arquivo na função
- Validamos se não existe nenhum erro com o `if err ≃ nil`
- Fechamos a geração do arquivo "salvamos" com `f.Close()`

Vamos agora inserir algum informação em nosso arquivo e imprimir o tamanho do arquivo gerado

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho %d bytes\n", tamanho)
	f.Close()

}
```

Incluimos algumas instruções ao nosso programa por exemplo `f.WriteString("Hello, World!")` esta função nos da a possibilidade de gravar string dentro do arquivo `f` que criamos

E também inferimos uma instrução para imprimir o tamanho do arquivo salvo, executando este programa teremos a seguinte saída.

output:
```shell
Arquivo criado com sucesso! Tamanho 13 bytes
```

Se executarmos agora `cat arquivo.txt` vamos ter o seguinte resultado:

```shell
Hello, World
```

Na maioria das vezes em um programa nao queremos somente gravar `strings` em um arquivo, e sim dados (bytes).

Para isso vamos vamos alterar a chamada da função `f.WriteString("Hello, World!")` para `tamanho, err := f.Write([]byte("Escrevendo alguma coisa no arquivo"))`

output:
```shell
Arquivo criado com sucesso! Tamanho 34 bytes
```

Se executarmos agora `cat arquivo.txt` vamos ter o seguinte resultado:

```shell
Escrevendo alguma coisa no arquivo
```

#### Lendo um arquivo

Muitas vezes precisamos ler um arquivo, para este exemplo vamos ler o arquivo gerado anteriormente.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))
}
```
output:
```shell
Escrevendo alguma coisa no arquivo
```

Ao executar a função `os.ReadFile(nome_do_arquivo)` o programa "lê" o arquivo e no nosso caso salva os valores na variavel `arquivo`

Para imprimir este valor convertemos a variavel `arquivo` em `string` com `string(arquivo)` pois toda vez que lemos arquivo os valores são slice de bytes `[]byte` então precisamos converter

Agora vamos pensar o seguinte cenario, temos uma maquina ou container docker com apenas 100mb de memoria, porem gostariamos de ler um arquivo com 1GB, tendo essas informações em mão conseguimos ver que teremos uma limitação de hardware para leitura deste arquivo, para isso podemos implementar da seguinte maneira.

Para isso vamos ler particionado este arquivo, e para isso vamos utilizar um pacote do `GO` que é o `bufio` esse pacote trabalha com buffers, e no nosso caso, vamos criar um `reader` onde o conteudo é `bufferizado` e nao lê o arquivo tudo de uma vez.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)

	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			println(".")
			break
		}

		fmt.Print(string(buffer[:n]))

	}
}
```
output:
```shell
Escrevendo alguma coisa no arquivo.
```
No exemplo acima podemos ver que
- Abrimos o arquivo `os.Open("arquivo.txt")`
- Criamos um leitor com `bufio.NewReader(arquivo2)`
- Criammos um buffer de 10 bytes com `make([]byte, 10)`
- Percorremos o arquivo dividindo-o pelo buffer na leitura `reader.Read(buffer)`

Se mudamor a implementação dentro do `for` um pouco podemos ver como o sistema imprimirá linha a linha de cada 10 bytes

```go
for {
    n, err := reader.Read(buffer)
    if err != nil {
        break
    }

    fmt.Println(string(buffer[:n]))
}
```
output
```shell
Escrevendo
 alguma co
isa no arq
uivo
```

#### Removendo um arquivo

Para remover um arquivo segue o exemplo abaixo


```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
```