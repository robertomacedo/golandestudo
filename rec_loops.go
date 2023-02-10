package main 

import "fmt"

func main() {
var x int = 7

	fmt.Println("Fatorial de ", x, "igual a:",loop(x))
}

func loop(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}