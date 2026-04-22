package main

import "fmt"

func main() {
	fmt.Println("=== TESTANDO A INVERSÃO DO TREM ===")

	lista := &DoublyLinkedList{}

	// Criando e engatando os vagões manualmente
	v3 := &Node2P{value: 30, next: nil}
	v2 := &Node2P{value: 20, next: v3}
	v1 := &Node2P{value: 10, next: v2}

	v3.prev = v2
	v2.prev = v1
	v1.prev = nil

	lista.head = v1
	lista.tail = v3
	lista.size = 3

	fmt.Print("Trem Original:   ")
	lista.Imprimir()

	// O Go sabe que o Reverse e o Imprimir estão lá no outro arquivo!
	lista.Reverse()

	fmt.Print("Trem Invertido:  ")
	lista.Imprimir()
	
	fmt.Println("===================================")
}