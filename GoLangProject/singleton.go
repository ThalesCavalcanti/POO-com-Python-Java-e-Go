package main

import (
	"fmt"
	"sync"
)

type Configuracao struct{}

var instancia *Configuracao
var once sync.Once

func GetInstancia() *Configuracao {
	once.Do(func() {
		instancia = &Configuracao{}
	})
	return instancia
}

func main() {
	config1 := GetInstancia()
	config2 := GetInstancia()
	fmt.Println(config1 == config2)
}
