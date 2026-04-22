package main

import "fmt"

func main() {
	fmt.Println("=== TESTANDO A INVERSÃO DO TREM ===")

	lista := &LinkedList{}

	// Criando e engatando os vagões manualmente
	v3 := &Node{Value: 30, Next: nil}
	v2 := &Node{Value: 20, Next: v3}
	v1 := &Node{Value: 10, Next: v2}

	lista.head = v1
	lista.size = 3

	fmt.Print("Trem Original:   ")
	lista.Imprimir()

	// O Go sabe que o Reverse e o Imprimir estão lá no outro arquivo!
	lista.Reverse()

	fmt.Print("Trem Invertido:  ")
	lista.Imprimir()
	
	fmt.Println("===================================")
}