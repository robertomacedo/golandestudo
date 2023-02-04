package main 

import "fmt"

type comida struct{
	nome	string
	sobrenome string
	sabores []string  //Slice para lista de sabores de comida
	horario int
}

func main(){
		p1 := comida{
			nome: "Roberto",
			sobrenome: "Macedo",
			hora: 19
			sabores: []string{"Peixe frito", "Cuxa", "Saladas"}}

		p2 := comida{ 
			nome: "Emily",
			sobrenome: "de Jesus Santos",
			sabores: []string{"Sorvete", "Frutas"}}

			fmt.Println(p1.nome, p1.sobrenome)
			for _, v := range p1.sabores{
				fmt.Println("\t", v)
			}

			fmt.Println(p2.nome, p2.sobrenome,"\n")
			for _, v := range p2.sabores{
				fmt.Println("\t", v)
			}

		fmt.Println(p1)
		fmt.Println(p2)
}