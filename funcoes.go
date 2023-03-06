package main 

import (
	"fmt"
)

func main() {
	c := gen()
	receber(c)

	fmt.Println("Fim da saída")
}
func gen() <-chan int {
	c := make(chan int)
	go func(){
		for i := 0; i < 10; i++{
			c <- i 
		}
		close(c)
	}()
	return c
}
func receber(canal <-chan int) {
	for v := range canal {
		fmt.Println(v)
	}
}