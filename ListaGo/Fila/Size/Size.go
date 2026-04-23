package main

import (
	"errors"
	"fmt"
)

type IQueue interface {
    Enqueue(value int)
    Dequeue() (int, error)
    Front() (int, error)
    IsEmpty() bool
    Size() int
}

type ArrayQueue struct {
    values []int
    front  int
    rear   int
}

// Cria uma nova fila com capacidade fixa
func NewArrayQueue(capacity int) *ArrayQueue {
    return &ArrayQueue{
        values: make([]int, capacity),
        front:  -1,
        rear:   -1,
    }
}

// Insere um elemento na fila (caráter circular)
func (queue *ArrayQueue) Enqueue(value int) {
    if (queue.rear+1)%len(queue.values) == queue.front {
        fmt.Println("Fila cheia!")
        return
    }

    if queue.front == -1 {
        queue.front = 0
    }
    queue.rear = (queue.rear + 1) % len(queue.values)
    queue.values[queue.rear] = value
}

// Remove um elemento da fila
func (queue *ArrayQueue) Dequeue() (int, error) {
    if queue.IsEmpty() {
        return 0, errors.New("Fila vazia")
    }

    val := queue.values[queue.front]

    if queue.front == queue.rear {
        queue.front = -1
        queue.rear = -1
    } else {
        queue.front = (queue.front + 1) % len(queue.values)
    }

    return val, nil
}

// Retorna o primeiro elemento da fila
func (queue *ArrayQueue) Front() (int, error) {
    if queue.IsEmpty() {
        return 0, errors.New("Fila vazia")
    }
    return queue.values[queue.front], nil
}

// Verifica se a fila está vazia
func (queue *ArrayQueue) IsEmpty() bool {
    return queue.front == -1
}

// Calcula o tamanho da fila sem variável size
func (queue *ArrayQueue) Size() int {
    if queue.front == -1 {
        return 0
    }
    if queue.rear >= queue.front {
        return queue.rear - queue.front + 1
    }
    return len(queue.values) - queue.front + queue.rear + 1
}