package main 

import "fmt"

func main(){
	z := [][]int {
		
		[]int{0, 1, 2, 3},			//0
		[]int{4,5,6},				//1
		[]int{7,8},					//2
		[]int{9, 10, 11, 12, 13},	//3
	}
	fmt.Println(z[2][1])			//posição do sice [2] na posição [1] valor = 8
	

	fmt.Println(z[1])      			//Valores do slice "z" no seu ídice [1] igual a [4 5 6]


}