package main

import "fmt"

type Carro struct {
	Marca      string
	Modelo     string
	Ano        int
	Velocidade int
}

func (c *Carro) Acelerar(quantidade int) {
	c.Velocidade += quantidade
}

func (c *Carro) Frear(quantidade int) {
	c.Velocidade -= quantidade
	if c.Velocidade < 0 {
		c.Velocidade = 0
	}
}

func (c Carro) ExibirVelocidade() {
	fmt.Printf("Velocidade atual: %d km/h\n", c.Velocidade)
}

func main() {
	carro := Carro{Marca: "Ford", Modelo: "Mustang", Ano: 2022}
	carro.Acelerar(50)
	carro.ExibirVelocidade()
	carro.Frear(20)
	carro.ExibirVelocidade()
}
