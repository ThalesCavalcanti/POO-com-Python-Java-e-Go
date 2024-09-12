package main

import "fmt"

// Definição da struct Empregado
type Empregado struct {
	Nome    string
	Cargo   string
	Salario float64
}

// Método para exibir detalhes do empregado
func (e Empregado) ExibirDetalhes() {
	fmt.Printf("Nome: %s, Cargo: %s, Salário: R$%.2f\n", e.Nome, e.Cargo, e.Salario)
}

// Definição da struct Empresa
type Empresa struct {
	Nome       string
	Empregados []Empregado // Uma empresa pode ter vários empregados
}

// Método para adicionar um empregado à empresa
func (e *Empresa) AdicionarEmpregado(empregado Empregado) {
	e.Empregados = append(e.Empregados, empregado)
}

// Método para listar todos os empregados da empresa
func (e Empresa) ListarEmpregados() {
	fmt.Printf("Lista de empregados da empresa %s:\n", e.Nome)
	for _, empregado := range e.Empregados {
		empregado.ExibirDetalhes()
	}
}

func main() {
	// Criando uma empresa
	empresa := Empresa{Nome: "Tech Solutions"}

	// Criando objetos Empregado
	empregado1 := Empregado{
		Nome:    "João Silva",
		Cargo:   "Desenvolvedor",
		Salario: 5000.00,
	}

	empregado2 := Empregado{
		Nome:    "Ana Costa",
		Cargo:   "Gerente de Projetos",
		Salario: 8500.00,
	}

	empregado3 := Empregado{
		Nome:    "Carlos Souza",
		Cargo:   "Analista de Sistemas",
		Salario: 6000.00,
	}

	// Adicionando empregados à empresa
	empresa.AdicionarEmpregado(empregado1)
	empresa.AdicionarEmpregado(empregado2)
	empresa.AdicionarEmpregado(empregado3)

	// Listando os empregados da empresa
	empresa.ListarEmpregados()
}
