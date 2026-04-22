package main

import "fmt"


type DoublyLinkedList struct{
	head *Node2P
	tail *Node2P
	size int
}

type Node2P struct{
	prev *Node2P
	value int
	next *Node2P

}

func (l *DoublyLinkedList) Reverse(){
	current := l.head
	var temp *Node2P = nil

	l.head, l.tail = l.tail, l.head

	for current != nil{
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}
}

func (l *DoublyLinkedList) Imprimir() {
	atual := l.head
	for atual != nil {
		fmt.Printf("[%d] -> ", atual.value)
		atual = atual.next
	}
	fmt.Println("nil (Fim da linha)")
}