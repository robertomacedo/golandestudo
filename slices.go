package main 

import "fmt"

func main(){
	slice := []string{"banana", "pera", "uva", "melão"}
	slice = append(slice, "cajú", "Articum")

	for indice, valor := range slice{
		fmt.Println("No indice: ",indice, "temos o valor:", valor)
	}
	// fmt.Println(slice)

	//for _, valor := range slice{
	//	fmt.Printf("Um dos valores do slice é: %v \n", valor)
	}

