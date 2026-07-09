package main

import "fmt"

/*
	Em GO, o conceito de array estatico e dinamico é abordado de uma outra maneira,
	aqui temos os chamados SLICES, eles sao uma struct leve (ponteiro + len + cap)
	que aponta para um array por baixo dos panos.

	Quando usamos SLICES, temos a sensação de estar usando um array dinamico.
*/

func main() {
	// no GO podemos definir um array "estatico" da seguinte forma
	static := [10]int{}
	fmt.Printf("Tamanho Array: %v \n", len(static))

	// Quando queremos criar um array "dinamico" podemos trabalhar com uma slice
	slice := []int{}
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}

	fmt.Printf("Tamanho Slice: %v \n", len(slice))

	// O GO um slice possui 2 dados importantes:
	// Len, que basicamente é a quantidades de itens presentes no array
	// Cap, essa é a capacidade que aquele array pode receber de itens

	// O GO gerencia a capacidade slice, de forma que, quando sua capacidade esta
	// proxima de ser ultrapassada, ele realoca o array/slice para um novo espaço da memoria
	// onde o mesmo pode pegar a cap atual e multiplicar ela por 1.25x ou 2x

	// A funcao make nos permite definir uma capacidade inicial para o nosso array

	new_slice := make([]int, 0, 10)
	fmt.Printf("Tamanho: %v, Capacidade: %v \n", len(new_slice), cap(new_slice))
	for i := 0; i < 15; i++ {
		new_slice = append(new_slice, i)
	}
	fmt.Printf("Tamanho: %v, Capacidade: %v \n", len(new_slice), cap(new_slice))
}
