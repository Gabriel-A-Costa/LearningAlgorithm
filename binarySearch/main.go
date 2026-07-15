package main

import "fmt"

func main() {
	arr := []int{2, 5, 7, 9, 20, 24, 125}
	var numero int

	max := len(arr) - 1
	min := 0
	fmt.Scanln(&numero)

	for min <= max {
		meio := (min + max) / 2
		chute := arr[meio]

		if chute == numero {
			fmt.Println(meio)
			return
		} else if chute < numero {
			min = meio + 1
		} else {
			max = meio - 1
		}
	}

	fmt.Println("Numero nao encontrado")
}
