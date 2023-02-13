package main 

import "fmt"

type garoto struct{
	nome 		string
	sobrenome 	string
	idade 		int 
}

func main() {

	p := garoto{"Davi", "Lucas da Silva", 15} // Atribuindo dados à variável p
	fmt.Println(p) // Mosntrando na tela o valor de "p"

	mudar(&p) //Chamando a função mudar
	fmt.Println(p) // Monstra na tela o valor alterado da variável "p" utilizando ponterio

}
func mudar(g *garoto) {
	(*g).nome = "Daniel" // Mudar Davi por Daniel
	g.sobrenome = "Silva Souza" // Mudar Lucas da Silva por Silva Souza

}