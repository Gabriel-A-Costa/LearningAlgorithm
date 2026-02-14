package main

import (
	"fmt"
)

func main() {
	list := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	number := 201

	baixo := 0
	alto := len(list)
	steps := 0

	for baixo < alto {
		meio := (baixo + alto) / 2
		steps++

		if list[meio] == number {
			fmt.Println("Steps:", steps)
			fmt.Println(number)
			return
		} else if meio > number {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}

	fmt.Println("Numero nao encontrado")
}
