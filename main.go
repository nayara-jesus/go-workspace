// Desafio 01
// Criando um Programa em Go para a Conversão de Escalas Termométricas

package main

import (
	"fmt"
	"math"
)

const ebulicao_K = 373.2

func main() {
	var K = ebulicao_K
	var C = (K - 273)
	fmt.Println("Ponto de ebulição da água em °K é:", K)
	fmt.Printf("Ponto de ebulição da água em °C é: %.1f", math.Round(C))
}

//go build main.go
//go run main.go
