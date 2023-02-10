package main 

import "fmt"

//Usando e entendo o recurso do clouse

func main(){
	a := i()
	fmt.Println(a()) 
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
}
func i() func() int {
	x := 0 
	return func() int {
		x++
		return x
	}
}