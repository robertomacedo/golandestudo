package main

import (
	"fmt"
	"strconv"
)
func hello(nome string){      //Função com argumento nome 
	fmt.Println("Olá", nome)

}

func soma(a, b int) int{   //Função soma
	return a + b
	return a + b

}

func conv(a int, b string) int{  // Função conver o argumento "b" no caso string para inteiro
	i, _ := strconv.Atoi(b)   	 // convete e soma os valores	
	return a + i
}

func main(){
	hello("Roberto Macedo")
	fmt.Println("Total:", soma(25, 4))
	fmt.Println("Total:", conv(25, "10")) //Saídas do programa
}