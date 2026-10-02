package main

import "fmt"

func main() {
	s := make([]int, 3) // uzunligi 3, sig‘imi 10

	fmt.Println(s) 

	a := append(s,4)

	a[2]=2

	a= append(a, 1,2,3,4,5)

	b := a
	b[0]=22

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)	

}
