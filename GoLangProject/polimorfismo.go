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

func emitirSom(animais []Animal) {
	for _, animal := range animais {
		animal.Som()
	}
}

func main() {
	animais := []Animal{Cachorro{}, Gato{}}
	emitirSom(animais)
}
