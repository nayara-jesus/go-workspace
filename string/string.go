package main

import "fmt"

func main() {
	fmt.Println(len("Hello World")) //conta a qtde de caracteres
	fmt.Println("Hello World"[2])   //le codigo binario
	fmt.Println("Hello" + "World")  //concatenacao
}

//go build string.go
//go run string.go
