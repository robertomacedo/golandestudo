package main 

import "fmt"

func main() {
	canal := make(chan int)

	go func() {  //Transformar em gorotine
	canal <- 54
	}()
	fmt.Println(<-canal)
}