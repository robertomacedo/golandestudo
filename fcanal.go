package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go enviar(10, canal1)

	go f2(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}

}
// canal envia x valores para o primeiro canal
func enviar(n int, canal chan int){
	for i := 0; i < n; i++ {
		canal <- i 
	}
	//fechar o canal
	close(canal) 
}
// função pega cada valor enviado para o primeiro canal
// e gera uma go rotine para cada valor
func f2 (canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- work(x)
			wg.Done()
		}(v)
	}
	// quando o trabalho terminar as rotinas serão enviadas para
	// o canal dois ou no caso a func f2
	wg.Wait()
	close(canal2)
}
func work(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n 
}