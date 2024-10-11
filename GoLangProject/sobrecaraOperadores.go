package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

func (p Produto) SomarPreco(outro Produto) float64 {
	return p.Preco + outro.Preco
}

func main() {
	produto1 := Produto{"Produto A", 50}
	produto2 := Produto{"Produto B", 70}
	fmt.Println(produto1.SomarPreco(produto2)) // Soma dos preços
}
