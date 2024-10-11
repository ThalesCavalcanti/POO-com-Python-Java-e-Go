package main

import "fmt"

type Funcionario interface {
	CalcularSalario() float64
}

type FuncionarioHorista struct {
	HorasTrabalhadas int
	ValorHora        float64
}

func (f FuncionarioHorista) CalcularSalario() float64 {
	return float64(f.HorasTrabalhadas) * f.ValorHora
}

type FuncionarioAssalariado struct {
	SalarioMensal float64
}

func (f FuncionarioAssalariado) CalcularSalario() float64 {
	return f.SalarioMensal
}

func main() {
	horista := FuncionarioHorista{160, 25}
	assalariado := FuncionarioAssalariado{5000}

	fmt.Printf("Salário horista: R$%.2f\n", horista.CalcularSalario())
	fmt.Printf("Salário assalariado: R$%.2f\n", assalariado.CalcularSalario())
}
