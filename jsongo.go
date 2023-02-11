package main 

import ("fmt"
		"encoding/json"
)

type pessoa struct{
	Nome string
	Profissao string
	Salario float64 
	Idade int

}
func main() {
	james := pessoa{
		Nome: "James",
		Profissao: "Vilão",
		Salario: 40000000.45,
		Idade: 54,
	}
	conan := pessoa{
		Nome: "Conan o Barbaro",
		Profissao: "Herói",
		Salario: 10000.15,
		Idade: 56,
	}

	j, err := json.Marshal(james)
	if err != nil {
		fmt.Println(err)
	}
	c, err := json.Marshal(conan)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	fmt.Println(string(c))

}