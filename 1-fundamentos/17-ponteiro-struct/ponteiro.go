package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) andou() {
	c.Nome = "Phellipe Rodrigues"
	fmt.Printf("O cliente %v andou\n", c.Nome)
}

type Conta struct {
	saldo float64
}

func NewConta() *Conta {
	return &Conta{saldo: 0.0}
}

func (c Conta) simular(valor float64) float64 {
	println("Simulando")
	c.saldo += valor
	return c.saldo
}

func (c *Conta) contemplarSimulacao(valorSimulacao float64) {
	println("Aplicando emprestimo")
	c.saldo = valorSimulacao
}
func main() {
	phellipe := Cliente{Nome: "Phellipe"}
	phellipe.andou()
	fmt.Println("Nome:", phellipe.Nome)

	conta := Conta{saldo: 0.0}
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	valorSimulacao := conta.simular(20000)
	fmt.Printf("O valor total da sua conta com a simulação é %.2f\n", valorSimulacao)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)
	conta.contemplarSimulacao(valorSimulacao)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta.saldo)

	conta1 := NewConta()
	fmt.Printf("O saldo de sua conta é %.2f\n", conta1.saldo)
	valorSimulacao1 := conta1.simular(20000)
	fmt.Printf("O valor total da sua conta com a simulação é %.2f\n", valorSimulacao1)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta1.saldo)
	conta1.contemplarSimulacao(valorSimulacao1)
	fmt.Printf("O saldo de sua conta é %.2f\n", conta1.saldo)

}
