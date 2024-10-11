package main

type Matematica struct{}

func (m Matematica) Fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * m.Fatorial(n-1)
}

func main() {
	m := Matematica{}
}
