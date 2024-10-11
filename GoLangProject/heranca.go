package main

import "fmt"

type Animal interface {
	Som()
}

type Cachorro struct{}

func (c Cachorro) Som() {
	fmt.Println("O cachorro faz: Au au!")
}

type Gato struct{}

func (g Gato) Som() {
	fmt.Println("O gato faz: Miau!")
}

func main() {
	cachorro := Cachorro{}
	gato := Gato{}

	cachorro.Som()
	gato.Som()
}
