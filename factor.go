package main 

import "fmt"

// Se x é igual a 1 retornar    1
// Se x é igual a 2 retorna     2 * (2-1)
// Se x é igual a 3 retornar    3 * (3-1) = 3*2 = 6 e 6 * (2-1) => 6*1 = 6
// Se x é igual a 4 retornar    4 * (4-1) = 4*3
//								12 * (3-1) = 12*2
//  							24 * (2-1) = 24*1
// 								24 fatorial
var x int = 4
func main() {

	fmt.Println("O fatorial de ", x, "é:", fatorial(x))
}
func fatorial(x int) int{
	if x == 1 { 
		return x
	}
	
	return x * fatorial(x-1)
}