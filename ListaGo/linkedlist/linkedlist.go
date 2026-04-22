package main

import (
	"fmt"
)

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}
type LinkedList struct{
	Head *Node
	size int
}
type Node struct{
	Value int 
	Next *Node
}
func (l *LinkedList) Add(value int){
	node := &Node{Value: value} // novo nó

	if l.Head == nil{
		l.Head = node
	} else {
		atual := l.Head

		for atual.Next != nil{
			atual = atual.Next
		}

		atual.Next = node
	}

	l.size++
}

func (l *LinkedList) AddOnIndex(value int, index int) error{
	if index >=0 && index <= l.size{
		newNode := &Node{Value: value}

		if index == 0{
			newNode.Next = l.Head
			l.Head = newNode
		} else {
			aux := l.Head
			for i :=0; i<index-1;i++{
				aux = aux.Next
			}
			newNode.Next = aux.Next
			aux.Next = newNode
		}
		l.size++
		return nil
	} else {
		return fmt.Errorf("Index Invalido: %d", index)
	}
}

func (l *LinkedList) RemoveOnIndex(index int) error{
	if index >= 0 && index < l.size{
		if index == 0{
			l.Head = l.Head.Next
		} else {
			aux := l.Head

			for i := 0; i< index-1; i++{
				aux = aux.Next
			}
				aux.Next = aux.Next.Next
		}
		l.size--
		return nil
	} else {
		return fmt.Errorf("Index Invalido: %d", index)
	}
}
func (l *LinkedList) Get(index int) (int, error){
	if index >= 0 && index < l.size{

	
	aux := l.Head

	for i := 0; i<index;i++{
		aux = aux.Next
	}
	return aux.Value, nil
} else {
		return 0, fmt.Errorf("Index Invalido: %d", index)
	}
}

func (l *LinkedList) Set(value int, index int) error{
	if index >= 0 && index < l.size {
		aux := l.Head

		for i :=0; i<index;i++{
			aux = aux.Next
		}

		aux.Value = value
		return nil
	} else {
		return fmt.Errorf("Index Invalido: %d", index)
	}
}
func (l *LinkedList) Size() int{
	return l.size
}

func imprimirLista(lista IList) {
	fmt.Print("Trem atual: [ ")
	for i := 0; i < lista.Size(); i++ {
		valor, _ := lista.Get(i)
		fmt.Printf("%d ", valor)
	}
	fmt.Println("]")
}