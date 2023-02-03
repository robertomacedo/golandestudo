package main 

import "fmt"

func main(){

	x := []int{1, 2, 3, 4, 5, 6}

	fmt.Print(x, "\n") // primeiroo resultado [1 2 3 4 5 6]
	x2 := append(x[:2], x[4:]...)
	fmt.Print(x2, "\n") // Após o append resultado [1 2 5 6]
	fmt.Print(x, "\n")  // Resultado do slice x após append [1 2 5 6 5 6] o terceiro 
						// elemento foi alterado devido ao append de x2

	
}