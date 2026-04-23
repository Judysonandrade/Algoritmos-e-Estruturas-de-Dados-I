package main

import (
	"fmt"
)


func main() {
    q := NewArrayQueue(5)

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

    q.Enqueue(40)
    q.Enqueue(50)
    q.Enqueue(60) // deve mostrar "Fila cheia!"
    fmt.Println("Tamanho final:", q.Size())

    // Testando circularidade
    q.Dequeue()
    q.Enqueue(70)
    fmt.Println("Tamanho após circularidade:", q.Size())

    fmt.Println("Elementos na fila:")
    for i := 0; i < q.Size(); i++ {
        fmt.Println(q.values[(q.front+i)%len(q.values)])
    }
}