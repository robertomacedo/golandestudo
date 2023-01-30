package main

import "fmt"

// Práticando o uso do if de forma bem simples.

func main (){

	d := 3
	if d == 0{
		fmt.Println("Estou muito cansado vou dormir")
	} else if d == 1 {
		fmt.Println("Ainda consido ficar acordado")
	} else {
		fmt.Println("Brincadeira vamos lá!")
	}
}