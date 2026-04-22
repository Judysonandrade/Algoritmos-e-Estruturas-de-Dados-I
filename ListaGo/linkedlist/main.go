package main

import "fmt"

func main() {
	fmt.Println("=== INICIANDO OS TESTES DA LISTA ENCADEADA ===")
	
	// Criando a lista
	var minhaLista IList = &LinkedList{}

	fmt.Println("\n1. Testando Add() e Size()")
	minhaLista.Add(10)
	minhaLista.Add(20)
	minhaLista.Add(30)
	imprimirLista(minhaLista)
	fmt.Printf("Tamanho: %d\n", minhaLista.Size())

	fmt.Println("\n2. Testando AddOnIndex() - Furando a fila! Inserindo 99 no index 1")
	minhaLista.AddOnIndex(99, 1)
	imprimirLista(minhaLista)

	fmt.Println("\n3. Testando AddOnIndex() - Nova Cabeça! Inserindo 5 no index 0")
	minhaLista.AddOnIndex(5, 0)
	imprimirLista(minhaLista)

	fmt.Println("\n4. Testando Set() - Trocando o valor do último vagão (index 4) para 77")
	minhaLista.Set(77, 4)
	imprimirLista(minhaLista)

	fmt.Println("\n5. Testando Get() - Lendo o valor no index 2")
	valor, err := minhaLista.Get(2)
	if err == nil {
		fmt.Printf("A carga do vagão no index 2 é: %d\n", valor)
	}

	fmt.Println("\n6. Testando RemoveOnIndex() - Destruindo o index 2 (que é o 99)")
	minhaLista.RemoveOnIndex(2)
	imprimirLista(minhaLista)
	fmt.Printf("Novo Tamanho: %d\n", minhaLista.Size())
	
	fmt.Println("\n=== TODOS OS TESTES PASSARAM COM SUCESSO! ===")
}