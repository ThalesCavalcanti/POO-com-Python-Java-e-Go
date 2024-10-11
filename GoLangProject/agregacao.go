package main

import "fmt"

type Empregado struct {
	Nome    string
	Cargo   string
	Salario float64
}

type Empresa struct {
	Nome       string
	Empregados []Empregado
}

func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
	e.Empregados = append(e.Empregados, empregado)
}

func (e Empresa) ListarEmpregados() {
	for _, empregado := range e.Empregados {
		fmt.Printf("%s, %s, R$%.2f\n", empregado.Nome, empregado.Cargo, empregado.Salario)
	}
}

func main() {
	empresa := Empresa{Nome: "Tech Solutions"}
	empresa.AdicionarEmpregado(Empregado{"Ana", "Gerente", 8000})
	empresa.AdicionarEmpregado(Empregado{"João", "Desenvolvedor", 5000})

	empresa.ListarEmpregados()
}
