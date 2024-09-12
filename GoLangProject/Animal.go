package main

import "fmt"

// Definição da struct base Animal
type Animal struct {
	Nome string
}

// Método Som para a struct Animal (pode ser sobrescrito)
func (a Animal) Som() {
	fmt.Println("O animal emite um som.")
}

// Definição da struct Cachorro, que "herda" de Animal através da composição
type Cachorro struct {
	Animal // Composição: Cachorro tem um Animal
}

// Método Som para Cachorro (sobrescreve o comportamento do método da base Animal)
func (c Cachorro) Som() {
	fmt.Println("O cachorro faz: Au Au!")
}

// Definição da struct Gato, que "herda" de Animal através da composição
type Gato struct {
	Animal // Composição: Gato tem um Animal
}

// Método Som para Gato (sobrescreve o comportamento do método da base Animal)
func (g Gato) Som() {
	fmt.Println("O gato faz: Miau!")
}

func main() {
	// Criando um objeto Cachorro
	cachorro := Cachorro{
		Animal: Animal{
			Nome: "Rex",
		},
	}

	// Criando um objeto Gato
	gato := Gato{
		Animal: Animal{
			Nome: "Mimi",
		},
	}

	// Exibindo o som de cada animal
	fmt.Printf("%s: ", cachorro.Nome)
	cachorro.Som()

	fmt.Printf("%s: ", gato.Nome)
	gato.Som()
}
