package main

import (
	"errors"
	"fmt"
)

type ContaBancaria struct {
	saldo   float64
	titular string
}

func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.saldo {
		return errors.New("saldo insuficiente")
	}
	c.saldo -= valor
	return nil
}

func main() {
	conta := ContaBancaria{titular: "João", saldo: 1000}
	err := conta.Sacar(1500)
	if err != nil {
		fmt.Println(err)
	}
}
