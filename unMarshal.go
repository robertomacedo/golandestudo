package main 

import ("fmt"
		"encoding/json"
)

// Programa pega uma estrutura em json e mostra na tela os dados da estrutura

type dados struct{
	Nome 		string 	`json: "Nome"`
	Sobrenome 	string 	`json: "Sobrenome"`
	Idade 		int 	`json: "Idade"`
	Profissao 	string 	`json: "Profissao"`
	Conta 		float64 `json: "Conta"`
}

func main() {
	sb := []byte(`{"Nome": "Arnaldo", "Sobrenome": "Das Couves", "Idade": 45, "Profissao": "Agente", "Conta": 1320.22}`)

	var p dados
	err := json.Unmarshal(sb, &p) 
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(p)
	fmt.Println(p.Idade)
}