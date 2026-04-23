package main

import (
	"errors"
)

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Front() (int, error)
	IsEmpty() bool
	Size() int
}

type ArrayQueue struct{
	value []int
}

func (l *ArrayQueue) Size() int{
	return len(l.value)
}

func (l *ArrayQueue) IsEmpty() bool{
	return len(l.value) == 0
}

func (l *ArrayQueue) Enqueue(value int){
	l.value = append(l.value, value)
}
func (l *ArrayQueue) Dequeue() (int, error){
	if l.IsEmpty(){
		return 0, errors.New("Fila Vazia")
	}
	val := l.value[0]
	l.value = l.value[1:]
	return val, nil

}

func (l *ArrayQueue) Front() (int, error){
	if l.IsEmpty(){
		return 0, errors.New("Fila Vazia")
	}

	return l.value[0], nil
}

