package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	contador := 0 

	totalgoroutina := 10

	var wg sync.WaitGroup 
	wg.Add(totalgoroutina)

	for i := 0; i < totalgoroutina; i++ {
		go func() {
			v := contador 
			runtime.Gosched() 
			v++
			contador = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(contador)
}