package main

import (
	"fmt"
)

func main() {
	list := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	number := 20

	baixo := 0
	alto := len(list) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2

		if chute := list[meio]; chute == number {
			fmt.Println(number)
			break
		} else if meio > number {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}

}
