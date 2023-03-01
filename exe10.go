package main 

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func () {
		c <- 65
	}()
	fmt.Println(<-c)
}