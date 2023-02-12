package main 


import "fmt" 
// implementando algoritmo ordenaçãoo
func buble (n []int) []int {
	for i, _ := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}

func main(){
	sort := []int{2,5,4,3,8,9,1,6}
	fmt.Println(buble(sort))
}