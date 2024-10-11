package main

import "fmt"

type Professor struct {
	Nome    string
	Escolas []*Escola
}

type Escola struct {
	Nome        string
	Professores []*Professor
}

func (p *Professor) AdicionarEscola(e *Escola) {
	p.Escolas = append(p.Escolas, e)
	e.Professores = append(e.Professores, p)
}

func main() {
	escola := Escola{Nome: "Escola Primária"}
	professor := Professor{Nome: "Maria"}
	professor.AdicionarEscola(&escola)

	fmt.Printf("%s leciona na %s\n", professor.Nome, escola.Nome)
}
