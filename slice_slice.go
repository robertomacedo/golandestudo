package main 

import "fmt"

func main(){

	ws := []string{"maça", "pepino", "graviola", "cajá", "bacuri"}

	fatia := ws[:]

	fmt.Println(fatia,"\n")

	for i := 0; i < len(ws); i++{
		fmt.Println(ws[i])
	}
}