package main

import (
	"fmt"
)

func main() {
    q := &Queue{}

    fmt.Println("Fila vazia?", q.IsEmpty())

    q.Enqueue(10)
    q.Enqueue(20)
    q.Enqueue(30)

    fmt.Println("Tamanho:", q.Size())

    val, _ := q.Front()
    fmt.Println("Primeiro da fila:", val)

    val, _ = q.Dequeue()
    fmt.Println("Removido:", val)

    fmt.Println("Tamanho atual:", q.Size())
}