package main

import "fmt"

func main(){
	fmt.Println("==== TESTE ====")

	lista := &ArrayList{}

	lista.values = []int{10,20,30,40,50}
	lista.inserted = 5

	fmt.Println("==== ORIGINAL ====")
	lista.Reverse()

	fmt.Println("==== MUDANÇA ====")
	fmt.Println(lista.values[:lista.inserted])
}