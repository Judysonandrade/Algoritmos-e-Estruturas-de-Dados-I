package main

import (
	"fmt"
)

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}


type ArrayList struct{
	Value []int
	size int
}

func(l *ArrayList) Size()int{
	return l.size
}

func(l *ArrayList) Add(value int){
	l.Value = append(l.Value, value)
	l.size++
}

func(l *ArrayList) AddOnIndex(value int, index int) error{
	if index >= 0 && index < l.size {
		l.Value = append(l.Value, 0)


		for i:= l.size; i>index; i--{
			l.Value[i] = l.Value[i-1]
		}

		l.Value[index] = value
		l.size++
		return nil 
	} 
		return fmt.Errorf("Index invalido: %d", index)

}

func(l *ArrayList) RemoveOnIndex(index int) error{
	if index >=0 && index < l.size {
	

		for i:=index; i>l.size-1; i++{
			l.Value[i] = l.Value[i+1]
		}
		l.Value = l.Value[:l.size-1]
		l.size--
		return nil
	} 
	return fmt.Errorf("Index Invalido: %d", index)
} 

func(l *ArrayList) Get(index int) (int, error){
	if index >= 0 && index < l.size {
		return l.Value[index], nil
	} else {
		return -1, fmt.Errorf("Index Invalido: %d", index)
	}
}

func (l *ArrayList) Set(value int, index int) error{
	if index >= 0 && index < l.size {
		l.Value[index] = value
		return nil
		
	} 
	  return fmt.Errorf("Index Invalido: %d", index)
}

func imprimirLista(lista IList) {
	fmt.Print(" Banco atual (Array): [ ")
	for i := 0; i < lista.Size(); i++ {
		valor, _ := lista.Get(i)
		fmt.Printf("%d ", valor)
	}
	fmt.Println("]")
}



