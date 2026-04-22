package main

type ArrayList struct {
	values   []int
	inserted int
}

func (l *ArrayList) Reverse() {
	inicio := 0

	fim := l.inserted - 1

	for inicio < fim {
		l.values[inicio], l.values[fim] = l.values[fim], l.values[inicio]

		inicio++
		fim--
	}

}
