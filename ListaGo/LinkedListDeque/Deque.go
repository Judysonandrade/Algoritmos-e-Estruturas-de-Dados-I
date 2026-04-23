package main

import (
	"errors"
)
type DoubleNode struct {
	Value int
	next  *DoubleNode
	prev  *DoubleNode
}

type LinkedListQueue struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

func (q *LinkedListQueue) EnqueueFront(value int) {
	newNode := &DoubleNode{Value: value}

	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		newNode.next = q.head
		q.head.prev = newNode
		q.head = newNode
	}
	q.size++
}

func (q *LinkedListQueue) EnqueueRear(value int) {
	newNode := &DoubleNode{Value: value}

	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		newNode.prev = q.tail
		q.tail.next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *LinkedListQueue) DequeueFront() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}

	valor := q.head.Value
	if q.head == q.tail { // Apenas um elemento
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.head.prev = nil
	}

	q.size--
	return valor, nil
}

func (q *LinkedListQueue) DequeueRear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}

	valor := q.tail.Value
	if q.head == q.tail { // Apenas um elemento
		q.head = nil
		q.tail = nil
	} else {
		q.tail = q.tail.prev
		q.tail.next = nil
	}

	q.size--
	return valor, nil
}

func (q *LinkedListQueue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	return q.head.Value, nil
}

func (q *LinkedListQueue) Rear() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Deque Vazio")
	}
	return q.tail.Value, nil
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LinkedListQueue) Size() int {
	return q.size
}