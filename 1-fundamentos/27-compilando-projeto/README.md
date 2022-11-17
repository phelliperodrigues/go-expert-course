# Capitulo Fundamento
## Aula 27 - Compulando Projeto

`GO` é uma linguagem compilada

Em `GO` vc tem a possibilidade de escolher qual plataforma você queira gerar um binario

Vamos executar o seguinte comando
```shell
go build main.go

ls
```
output:
```shell
main  main.go  README.md
```

Podemos ver que temos o arquivo binario `main`, se executarmos:
```shell
./main
```
output:
```shell
Hello, World!
```

Como temos diversas plataformas que podem executar nosso programa como `linux`, `windows`, `mac`, `intel`, `arm` etc.

Vamos compilar nosso programa para gerar um executavel para windows.

```shell
GOOS=windows go build main.go
```
output:
```shell
main.exe  main.go  README.md
```

Podemos ver que foi gerado agora um arquivo `main.exe` que é um executavel do windows

Lista de OS suportado pelo go

```shell
"aix", "android", "darwin", "dragonfly", "freebsd", "hurd", "illumos", "ios", "js", "linux", "nacl", "netbsd", "openbsd", "plan9", "solaris", "windows", "zos"
```

Um outro exemplo de compulação é que alem do OS você pode escolhar a arquitetura, no exemplo abaixo o OS é MAC com arquitetura `arm64` que é do chip M1

```shell
GOOS=darwin GOARCH=arm64 go build main.go
```

Existe um comando que exibe todas as versoes de plataformas que podemos compular
```shell
go tool dist list  
```
```shell
OS/ARCH

aix/ppc64                                                                                                                                                                                                                                                                                            ─╯
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
...
```

Quando temos um modulo inicializado em nossa aplicação com o `go mod init`, não é necessario apontar qual arquivo queremos buildar como feito anteriormente com `go build main.go` basta somente executar `go build` que o `GO` vai gerar o binanrio do modulo todo.

