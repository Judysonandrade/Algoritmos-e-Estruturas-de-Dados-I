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

type Node struct{
	Next  *Node
	prev *Node
	Value int
}

type DoublyLinkedList struct{
	Head *Node
	Tail *Node
	size int
}

func(l *DoublyLinkedList) Add(value int){
	node := &Node{Value:value}
	if l.Head == nil{
		l.Head = node

	} else {
		atual := l.Head
		
		for atual.Next != nil{
			atual = atual.Next
		}
			atual.Next = node
			node.prev = atual
	}
	l.size++
}

func (l *DoublyLinkedList) AddOnIndex(value int, index int) error {
	// Tem que ser <= l.size para permitir furar a fila no último lugar
	if index >= 0 && index <= l.size {
		newNode := &Node{Value: value}

		// CASO 1: Inserindo no começo (Index 0)
		if index == 0 {
			newNode.Next = l.Head
			if l.Head != nil {
				l.Head.prev = newNode // usando o prev minúsculo
			}
			l.Head = newNode
			
		// CASO 2: Inserindo no meio ou no final
		} else {
			aux := l.Head

			// Paramos exatamente UM VAGÃO ANTES da posição (index - 1)
			for i := 0; i < index-1; i++ {
				aux = aux.Next
			}

			// MÁGICA DOS ENGATES:
			newNode.Next = aux.Next
			newNode.prev = aux

			// Se não estivermos inserindo no exato final, avisamos o vagão da frente
			if aux.Next != nil {
				aux.Next.prev = newNode
			}

			aux.Next = newNode
		}
		
		l.size++
		return nil
	}
	
	// Caso o index seja inválido (negativo ou maior que a lista)
	return fmt.Errorf("Index Invalido: %d", index)
}

func(l *DoublyLinkedList)RemoveOnIndex(index int) error{
	if index>=0 && index < l.size{
		if index == 0{
			l.Head = l.Head.Next
		} else {
			aux := l.Head
			for i :=0; i<index; i++{
				aux = aux.Next
			}

			aux.prev.Next = aux.Next

			if aux.Next != nil{
				aux.Next.prev = aux.prev
			}
		}
		l.size--
		return nil
	}
	return fmt.Errorf("Index Invalido : %d", index)

}

func (l *DoublyLinkedList) Get(index int) (int, error){
	if index >=0 && index < l.size{
		aux := l.Head

		for i:=0;i<index;i++{
			aux = aux.Next
		}
		return aux.Value, nil
	} 
	
		return 0, fmt.Errorf("Index Invalido: %d", index)
}

func (l *DoublyLinkedList) Set(value int, index int) error{
	if index >=0 && index < l.size{
		aux := l.Head

		for i:= 0;i<index;i++{
			aux = aux.Next
		}
		 aux.Value = value
		 return nil
	} 
	return fmt.Errorf("Index Invalido: %d", index)
}

func(l *DoublyLinkedList) Size()int{
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