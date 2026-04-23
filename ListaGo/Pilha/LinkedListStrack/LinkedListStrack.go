package main

import "errors"

// A interface que define as regras da nossa Pilha
type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type Node struct{
	Value int
	next *Node
}

type Strack struct{
	size int
	head *Node
}

func (l *Strack) IsEmpty() bool{
	return l.head == nil
}

func (l *Strack) Push(value int){
	NewNode := &Node{Value : value}

	NewNode.next = l.head
	l.head = NewNode

	l.size++
}

func(l *Strack) Pop() (int, error){
	if l.IsEmpty(){
		return 0, errors.New("Pilha Vazia")
	}

	Valor := l.head.Value
	l.head = l.head.next

	l.size--

	return Valor, nil
}

func(l *Strack) Size() int{
	return l.size
}
