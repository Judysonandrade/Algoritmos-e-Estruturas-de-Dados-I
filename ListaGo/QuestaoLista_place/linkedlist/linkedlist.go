package main

import "fmt"

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	Next  *Node
	Value int
}

func (l *LinkedList) Reverse() {
	var prev *Node = nil
	current := l.head
	var Next *Node = nil

	for current != nil {
		Next = current.Next
		current.Next = prev

		prev = current
		current = Next
	}
	l.head = prev
}
func (l *LinkedList) Imprimir() {
	atual := l.head
	for atual != nil {
		fmt.Printf("[%d] -> ", atual.Value)
		atual = atual.Next
	}
	fmt.Println("nil (Fim da linha)")
}