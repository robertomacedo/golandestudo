package main 

import "fmt"

func main(){

	x := 25

	

	func(x int){
		fmt.Println(x, "vezes 5 é: ")
		fmt.Println(x * 5)
		
	}(x)

}