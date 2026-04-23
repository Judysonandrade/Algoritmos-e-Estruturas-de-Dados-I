package main

import (
	"errors"
)

// A interface fornecida
type IDeque interface {
	EnqueueFront(value int)
	EnqueueRear(value int)
	DequeueFront() (int, error)
	DequeueRear() (int, error)
	Front() (int, error)
	Rear() (int, error)
	IsEmpty() bool
	Size() int
}

// ==========================================
// 1. Implementação: ArrayQueue
// ==========================================

type ArrayQueue struct {
	items []int
}

func (q *ArrayQueue) EnqueueFront(value int) {
	// Prepend no slice (adiciona no início)
	q.items = append([]int{value}, q.items...)
}

func (q *ArrayQueue) EnqueueRear(value int) {
	// Append no slice (adiciona no final)
	q.items = append(q.items, value)
}

func (q *ArrayQueue) DequeueFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	valor := q.items[0]
	q.items = q.items[1:] // Remove o primeiro elemento
	return valor, nil
}

func (q *ArrayQueue) DequeueRear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	ultimoIndice := len(q.items) - 1
	valor := q.items[ultimoIndice]
	q.items = q.items[:ultimoIndice] // Remove o último elemento
	return valor, nil
}

func (q *ArrayQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	return q.items[0], nil
}

func (q *ArrayQueue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	return q.items[len(q.items)-1], nil
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.items)
}