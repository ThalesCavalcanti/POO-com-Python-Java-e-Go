package main

import "fmt"

// Definição da struct Carro
type Carro struct {
	Marca  string
	Modelo string
	Ano    int
}

// Método para exibir detalhes do carro
func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Ano: %d\n", c.Marca, c.Modelo, c.Ano)
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
}
