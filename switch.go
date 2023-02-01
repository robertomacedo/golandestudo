package main

import "fmt"

// Práticando o uso do if de forma bem simples.

func main (){

	d := 0

	switch{
	case d == 0:
		fmt.Println("Estou muito cansado vou dormir")
	case d == 1:
		fmt.Println("Vamos continuar")
	case d == 2:
		fmt.Println("Brincadeira, não tem como!")
	}

}