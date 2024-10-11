package main

import "fmt"

type ContaBancaria struct {
	saldo   float64
	titular string
}

func (c *ContaBancaria) Depositar(valor float64) {
	c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) {
	if valor <= c.saldo {
		c.saldo -= valor
	} else {
		fmt.Println("Saldo insuficiente.")
	}
}

func (c ContaBancaria) ExibirSaldo() {
	fmt.Printf("Saldo atual: R$%.2f\n", c.saldo)
}

func main() {
	conta := ContaBancaria{titular: "João", saldo: 1000}
	conta.Depositar(500)
	conta.Sacar(200)
	conta.ExibirSaldo()
}
