package main

import "fmt"

// Definição da struct Professor
type Professor struct {
	Nome    string
	Escolas []*Escola // Um professor pode lecionar em várias escolas
}

// Método para adicionar uma escola ao professor
func (p *Professor) AdicionarEscola(e *Escola) {
	p.Escolas = append(p.Escolas, e)
}

// Método para listar as escolas do professor
func (p Professor) ListarEscolas() {
	fmt.Printf("Professor: %s leciona nas seguintes escolas:\n", p.Nome)
	for _, escola := range p.Escolas {
		fmt.Printf("- %s\n", escola.Nome)
	}
}

// Definição da struct Escola
type Escola struct {
	Nome        string
	Professores []*Professor // Uma escola pode ter vários professores
}

// Método para adicionar um professor à escola
func (e *Escola) AdicionarProfessor(p *Professor) {
	e.Professores = append(e.Professores, p)
}

// Método para listar os professores da escola
func (e Escola) ListarProfessores() {
	fmt.Printf("Escola: %s tem os seguintes professores:\n", e.Nome)
	for _, professor := range e.Professores {
		fmt.Printf("- %s\n", professor.Nome)
	}
}

func main() {
	// Criando objetos Professor
	professor1 := &Professor{Nome: "Maria"}
	professor2 := &Professor{Nome: "Carlos"}
	professor3 := &Professor{Nome: "Ana"}

	// Criando objetos Escola
	escola1 := &Escola{Nome: "Escola Primária Central"}
	escola2 := &Escola{Nome: "Colégio Novo Horizonte"}

	// Associação entre professores e escolas
	// Escola 1 com professor1 e professor2
	escola1.AdicionarProfessor(professor1)
	escola1.AdicionarProfessor(professor2)

	// Escola 2 com professor2 e professor3
	escola2.AdicionarProfessor(professor2)
	escola2.AdicionarProfessor(professor3)

	// Associando escolas aos professores
	professor1.AdicionarEscola(escola1)
	professor2.AdicionarEscola(escola1)
	professor2.AdicionarEscola(escola2)
	professor3.AdicionarEscola(escola2)

	// Listando professores por escola
	escola1.ListarProfessores()
	escola2.ListarProfessores()

	// Listando escolas por professor
	professor1.ListarEscolas()
	professor2.ListarEscolas()
	professor3.ListarEscolas()
}
