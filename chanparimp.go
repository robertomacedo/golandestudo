package main 

import "fmt"

// usando canal para mandar números pares e impares para 
// canais diferentes e mostra tudo na tela 
// sinalizando se o número é par ou ímpar

func main () {
	par := make(chan int)
	impa := make(chan int) 
	quit := make(chan bool)

	go enviaNumber(par, impa, quit)

	receber(par, impa, quit)
}
func enviaNumber (par, impa chan int, quit chan bool){
	total := 50 
	for i := 0; i < total; i++{
		if i % 2 == 0 {
			par <- i
		}else{
			impa <- i 
		}
	}
	close(par)
	close(impa)
	quit<-true 
}
func receber(par, impa chan int, quit chan bool){
	for {
		select{
		case v := <-par:
			fmt.Println("O número", v, " é par") // mostra na tela o valor de v sendo par
		case v := <-impa:
			fmt.Println("O número", v, "é impar") // mostra o valor de v sendo impar
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra!", v)
			}else {
				fmt.Println("Encerrando. Tudo certo!", v)
			}
			return
		}
	}
}