package main

import "fmt"

func revertText(text string) {
	//Primeiro vc transforma a string em um rune "OLA" = ["O", "L", "A"]
	runes := []rune(text)

	//Definir o index do primeiro ponto
	start := 0

	//Fazer loop, para mover o segundo ponto para o final da palavra
	for i := 0; i <= len(runes); i++ {
		//Se o segundo ponto chegar ao final do texto ou a um espaço ele iniciar a troca de letras
		if i == len(runes) || runes[i] == ' ' {
			//Define a posição da primeira e ultima letra
			left, right := start, i-1
			//Enquanto as posições não forem iguais faz as trocas delas
			for left < right {
				runes[left], runes[right] = runes[right], runes[left]
				left++
				right--
			}
			//Define o novo ponto inicia na proxima palavra
			start = i + 1
		}
	}

	fmt.Println(string(runes))
}

func main() {
	text := "Hello World "
	revertText(text)
}
