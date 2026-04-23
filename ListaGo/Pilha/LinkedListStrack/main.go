package main

import (
	"fmt"
)

func main() {
    // Criando uma pilha
    stack := &Strack{}

    // Testando IsEmpty
    fmt.Println("A pilha está vazia?", stack.IsEmpty())

    // Inserindo elementos
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    // Verificando tamanho
    fmt.Println("Tamanho da pilha:", stack.Size())

    // Removendo elementos
    valor, err := stack.Pop()
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Pop:", valor)
    }

    valor, err = stack.Pop()
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Pop:", valor)
    }

    // Verificando tamanho novamente
    fmt.Println("Tamanho da pilha:", stack.Size())

    // Testando Pop em pilha vazia
    stack.Pop()
    stack.Pop() // aqui a pilha já estará vazia
    valor, err = stack.Pop()
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Pop:", valor)
    }
}
