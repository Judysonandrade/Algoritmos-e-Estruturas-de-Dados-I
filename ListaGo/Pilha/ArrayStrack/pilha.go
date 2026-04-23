package main

import "errors"

type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
	Size() int
}

type Strack struct {
	items []int
}

func (l *Strack) Push(value int) {
	l.items = append(l.items, value)
}

func (l *Strack) Pop() (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("A pilha ta vazia")
	}
	ultimo := len(l.items) - 1
	value := l.items[ultimo]

	l.items = l.items[:ultimo]

	return value, nil
}

func(l *Strack) Peek()(int, error){
	if l.IsEmpty(){
		return 0, errors.New("A pilha ta vazia")

	}
	return l.items[len(l.items)-1],nil
}

func(l *Strack) IsEmpty() bool{
	return len(l.items) == 0
}

func(l *Strack) Size() int{
	return len(l.items)
}