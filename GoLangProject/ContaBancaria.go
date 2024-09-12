package main

import "fmt"

// Definição da struct ContaBancaria
type ContaBancaria struct {
	titular string  // Titular da conta (privado)
	saldo   float64 // Saldo da conta (privado)
}

// Função para criar uma nova conta bancária
func NovaContaBancaria(titular string, saldoInicial float64) *ContaBancaria {
	return &ContaBancaria{
		titular: titular,
		saldo:   saldoInicial,
	}
}

// Método para depositar dinheiro na conta
func (c *ContaBancaria) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
		fmt.Printf("Depósito de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
	} else {
		fmt.Println("Valor de depósito inválido.")
	}
}

// Método para sacar dinheiro da conta
func (c *ContaBancaria) Sacar(valor float64) {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		fmt.Printf("Saque de R$%.2f realizado com sucesso. Saldo atual: R$%.2f\n", valor, c.saldo)
	} else {
		fmt.Println("Saque inválido ou saldo insuficiente.")
	}
}

// Método para exibir o saldo atual (somente leitura)
func (c *ContaBancaria) ExibirSaldo() {
	fmt.Printf("O saldo atual da conta é: R$%.2f\n", c.saldo)
}

func main() {
	// Criando uma nova conta bancária
	conta := NovaContaBancaria("João Silva", 1000.0)

	// Exibindo saldo inicial
	conta.ExibirSaldo()

	// Depositando dinheiro
	conta.Depositar(500.0)

	// Sacando dinheiro
	conta.Sacar(300.0)

	// Tentativa de saque maior que o saldo
	conta.Sacar(1500.0)

	// Exibindo saldo final
	conta.ExibirSaldo()
}
