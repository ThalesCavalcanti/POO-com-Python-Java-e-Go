package main

import "fmt"

// Definição da struct Carro
type Carro struct {
	Marca      string
	Modelo     string
	Ano        int
	Velocidade int
}

// Método para exibir detalhes do carro
func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
}

// Método para acelerar o carro (aumenta a velocidade)
func (c *Carro) Acelerar(quantidade int) {
	c.Velocidade += quantidade
	fmt.Printf("O carro acelerou para %d km/h.\n", c.Velocidade)
}

// Método para frear o carro (diminui a velocidade)
func (c *Carro) Frear(quantidade int) {
	c.Velocidade -= quantidade
	if c.Velocidade < 0 {
		c.Velocidade = 0
	}
	fmt.Printf("O carro desacelerou para %d km/h.\n", c.Velocidade)
}

// Método para exibir a velocidade atual
func (c Carro) ExibirVelocidade() {
	fmt.Printf("A velocidade atual do carro é %d km/h.\n", c.Velocidade)
}

func main() {
	// Instanciando três objetos do tipo Carro
	carro1 := Carro{
		Marca:  "Toyota",
		Modelo: "Corolla",
		Ano:    2020,
	}

	carro2 := Carro{
		Marca:  "Honda",
		Modelo: "Civic",
		Ano:    2018,
	}

	carro3 := Carro{
		Marca:  "Ford",
		Modelo: "Mustang",
		Ano:    2022,
	}

	// Exibindo os detalhes de cada carro
	carro1.ExibirDetalhes()
	carro2.ExibirDetalhes()
	carro3.ExibirDetalhes()

	// Exibindo a velocidade inicial
	carro1.ExibirVelocidade()

	// Acelerando o carro
	carro1.Acelerar(30)
	carro1.Acelerar(20)

	// Freando o carro
	carro1.Frear(15)

	// Exibindo a velocidade atual
	carro1.ExibirVelocidade()
}
