package main

import "fmt"

type Imprimivel interface {
	Imprimir()
}

type Relatorio struct{}

func (r Relatorio) Imprimir() {
	fmt.Println("Imprimindo relatório...")
}

type Contrato struct{}

func (c Contrato) Imprimir() {
	fmt.Println("Imprimindo contrato...")
}

func main() {
	documentos := []Imprimivel{Relatorio{}, Contrato{}}
	for _, doc := range documentos {
		doc.Imprimir()
	}
}
