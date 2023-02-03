package main 

import "fmt"
								// Make cria mapas, fatias e rotorna um inicializado não zerado
								//make inicializa a estrutura de dados internos e prepara 
								//o valor para us

func main(){
	slice := make([]int, 5, 10)
								// A cada posição do slice é adicionado um elemento
								//no exemplo posicção [0]=1, [2]=3
	slice[0], slice[1] = 1, 3

	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 11)

	fmt.Println(slice, len(slice), cap(slice))


	// Saída do codigo [1 0 3 0 0] 5 10 o número "5" representa a quatidade
	// de elementos do slice e número "10" representa a capacidade
	// quando for aumentanda a capacidade dobra

	// Nova saída [1 0 3 0 0 4 5 7] 8 10 após o append houve o aumentoo
	// de posição no array subiu para "8" e capacidade permanceu "10"

	// Aumentamos o número de posições para "11" [1 0 3 0 0 4 5 7 8 9 11] 11 "20"
	// nesse momento a capacidade que era de "10" passa a ser 20
}