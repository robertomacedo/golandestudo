package main 

import "fmt"

func main(){

	x := 256

	y := func (x int) int{
		//fmt.Println(x, "vezes 4 é: ")
		return x * 4
	}
	fmt.Println(x, "vezes 4 é: ", y(x))
	
}