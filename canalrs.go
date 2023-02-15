package main 

import "fmt" 

// canais bidirecionais
func main() {
	canal := make(chan int)

	go send(canal) //função send vai enviar um valor
	receber(canal)
}
func send(s chan<- int) {
	s <- 49
}
// Função recebe o valor da função send
func receber(r <-chan int) {
	fmt.Println("O valor recebido foi:", <-r)
}