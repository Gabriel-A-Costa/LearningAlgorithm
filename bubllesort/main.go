package main

import "fmt"

func main() {
	arr := []int{123, 32, 1, 35, 2, 432, 321, 21}

	//NAVEGAR POR TODO O ARRAY
	for i := 0; i < len(arr)-1; i++ {
		//DEFINE COMO TERMINADO
		swapp := false
		//NAVEGA NOVAMENTE PELO ARRAY ATE O ITEM QUE TEM O MAIOR VALOR
		for j := 0; j < len(arr)-1-i; j++ {
			//VERIFICANDO SE O ITEM ATUAL É MAIOR QUE O PROXIMO
			if arr[j] > arr[j+1] {
				//SE FOR MAIOR, FAZ A TROCA
				arr[j], arr[j+1] = arr[j+1], arr[j]
				//DEFINE QUE AINDA TEM TROCA A SE FAZER
				swapp = true
			}
		}

		/*
			APOS IR VERIFICANDO SE O ITEM ATUAL É O MAIO, AO CHEGAR NO FINAL, VC VAI TER DEFINIDO QUE O ULTIMO ITEM
			É O MAIOR IR TEM DO ARRAY, POR ISSO QUE AO INICIAR UMA NOVA NAVEGACAO PELO ARRY SECUNDARIO, FAZEMOS:
			len(arr)-1 [PARA NAO ESTORA] -i [PARA DESCARTAR OS ITEMS QUE JA FORAM DEFINIDO COMO MAIOR]
		*/

		if !swapp {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(len(arr))

	fmt.Println(arr)
}
