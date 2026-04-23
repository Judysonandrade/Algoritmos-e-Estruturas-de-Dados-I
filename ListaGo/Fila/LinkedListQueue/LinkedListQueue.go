package main

import "errors"

// Interface da fila
type IQueue interface {
    Enqueue(value int)
    Dequeue() (int, error)
    Front() (int, error)
    IsEmpty() bool
    Size() int
}

type Queue struct{
	size int
	head *Node
}

type Node struct{
	value int
	Next *Node
}

func (l *Queue) Size() int{
	return l.size
}

func(l *Queue) IsEmpty() bool{
	return l.head == nil
}

func (l *Queue) Enqueue(value int){
	NewNode := &Node{value: value}

	if l.head == nil{
		l.head = NewNode
	} else {
		current := l.head
		for current.Next != nil{
			current = current.Next
		}
		current.Next = NewNode
	}
	l.size++
}
func(l *Queue) Dequeue() (int, error){
	if l.IsEmpty(){
		return 0, errors.New("Fila Vazia")

	}
	val := l.head.value
	l.head = l.head.Next

	l.size--
	
	return val,nil
}

func(l *Queue) Front() (int,error){
	if l.IsEmpty(){
		return 0, errors.New("Fila Vazia")

	}
	return l.head.value, nil
}