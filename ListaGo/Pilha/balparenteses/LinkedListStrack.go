package main

import (
	"errors"
)

// Interface da pilha
type IStack interface {
    Push(value rune)
    Pop() (rune, error)
    Peek() (rune, error)
    IsEmpty() bool
    Size() int
}

type Node struct {
    Value rune
    next  *Node
}

type Strack struct {
    size int
    head *Node
}

func (l *Strack) IsEmpty() bool {
    return l.head == nil
}

func (l *Strack) Push(value rune) {
    NewNode := &Node{Value: value}
    NewNode.next = l.head
    l.head = NewNode
    l.size++
}

func (l *Strack) Pop() (rune, error) {
    if l.IsEmpty() {
        return 0, errors.New("Pilha Vazia")
    }
    valor := l.head.Value
    l.head = l.head.next
    l.size--
    return valor, nil
}

func (l *Strack) Peek() (rune, error) {
    if l.IsEmpty() {
        return 0, errors.New("Pilha Vazia")
    }
    return l.head.Value, nil
}

func (l *Strack) Size() int {
    return l.size
}

// Função que verifica se os parênteses estão balanceados usando a pilha Strack
func balparenteses(par string) bool {
    stack := &Strack{}

	for _, ch := range par {
		switch ch {
		case '(':
			stack.Push(ch)
		case ')':
			_, err := stack.Pop()
			if err != nil {
				return false
			}
		}
	}

    // Se a pilha estiver vazia, está balanceado
    return stack.IsEmpty()
}