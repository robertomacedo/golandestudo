package main 

import "fmt"

// Duas go func enviam x valores cada, para apenas um canal
// usando um loop for de x valores e um select case <-x

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++{
			a <- i 
		}
	}(x/2)
	
	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i 
		}
	}(x/2)

	for i := 0; i < x; i++{
		select {
		case v := <-a: 
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
}