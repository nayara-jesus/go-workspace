package main

import "fmt"

func main() {
	var nome string = "Ana"
	var idade int = 20
	var versao float32 = 3.2
	fmt.Println("Nome:", nome, "Idade:", idade)
	fmt.Println("Versao do programa:", versao)
}

//go build inferencia.go
//go run inferencia.go
