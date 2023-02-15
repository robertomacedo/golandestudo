package main 

import "fmt" 

func main() {

	c := make(chan int)

	go mloop(10, c)

	for v := range c {
		fmt.Println("Recebido do canal:", v)
	}
}
func mloop(t int, s chan<- int) {
	for i := 0; i < t; i++{
		s <- i
	}
	close(s) // Fecha o canal
}