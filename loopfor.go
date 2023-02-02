package main 

import "fmt"

// Brincanndo Loop usando o break com a condição if 
// mostra o intervalo entre as entradas das variáveis.
func main(){
	anonascimento := 1976
	anoatual:= 2050

	for {
			if anonascimento > anoatual{
				break
			}
			fmt.Println(anonascimento)
			anonascimento++
	}
	
	// anonascimento := 1976; anonascimento <= 2050; anonascimento++{
	//	fmt.Println(anonascimento)

	// O códigoo acima faz um loop conforme entradas das variáveis

	}

