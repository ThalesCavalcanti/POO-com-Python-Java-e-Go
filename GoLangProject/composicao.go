package main

import "fmt"

type Motor struct {
	Potencia int
}

type Carro struct {
	Marca  string
	Modelo string
	Motor  Motor
}

func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Motor: %d cavalos\n", c.Marca, c.Modelo, c.Motor.Potencia)
}

func main() {
	motorV8 := Motor{Potencia: 450}
	carro := Carro{Marca: "Ford", Modelo: "Mustang", Motor: motorV8}
	carro.ExibirDetalhes()
}
