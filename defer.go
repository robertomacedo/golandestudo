package main 

import "fmt"

func main(){
	//Entendendo a func defer, deixa para executar a linha por último ou depois como queira entender.


	// Usando o defer para inverter a ordem de execução das linhas abaixo.
	
	defer fmt.Println("Primeiro Lugar.")
	fmt.Println("Segundo Lugar.")
}