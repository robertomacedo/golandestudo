package main

import (
	"fmt"
	"sort"
)

func main() {
	st := []string{"Eu irei", "Acho", "Bem melhor", "Criar", "Depois de amanhã,"}

	fmt.Println(st,"\n") // Mostra da forma como estar na linha da string
	sort.Strings(st) //Vai organizar as frases 
	fmt.Println(st) // Mostra as frases em ordem alfabética

}