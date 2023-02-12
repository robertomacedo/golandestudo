package main 

import "fmt"

func main() {
	// Um pontoreiro faz referência a um valor na memória
	x := 10
	y := &x // Endereço de "x" atribuido à variável "y"

	*y++ //Vai acrescentar uma unidade no valor de "x"

	fmt.Println(&x) // Retorna o endereço na memória da variável x
	fmt.Println("Valor de x acrescidoo de uma unidade:",*y) // Retorna o valor da variável x
	fmt.Println(y) 	// Acrescenta "1" no valor de "x" sem utilizar 
					// o valor de "x" diretamente
}