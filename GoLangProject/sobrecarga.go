package main

import "fmt"

type Calculadora struct{}

func (c Calculadora) Somar(a, b int) int {
	return a + b
}

func (c Calculadora) SomarTres(a, b, c int) int {
	return a + b + c
}

func main() {
	calc := Calculadora{}
	fmt.Println(calc.Somar(2, 3))
	fmt.Println(calc.SomarTres(2, 3, 4))
}
