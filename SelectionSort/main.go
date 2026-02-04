package main

import "fmt"

/*------------FIRST VERSION (out-of-place)---------------*/

// var maior int

// func main() {
// 	list := []int{10, 23, 1, 215, 1612, 23, 15, 612, 123, 12}
// 	newList := []int{}

// 	for len(list) > 0 {
// 		maiorInx := 0

// 		for i := 1; i < len(list); i++ {
// 			if list[i] > list[maiorInx] {
// 				maiorInx = i
// 			}
// 		}

// 		newList = append(newList, list[maiorInx])

// 		list = append(list[:maiorInx], list[maiorInx+1:]...)
// 	}

// 	fmt.Println(newList)
// }

/*------------SECOUND VERSION (out-of-place)---------------*/

// func buscarMenor(arr []int) int {
// 	menor := arr[0]
// 	menore_indice := 0

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < menor {
// 			menor = arr[i]
// 			menore_indice = i
// 		}
// 	}

// 	return menore_indice
// }

// func main() {

// 	arr := []int{10, 23, 1, 215, 1612, 23, 15, 612, 123, 12}
// 	novoArr := []int{}

// 	for len(arr) > 0 {
// 		menor := buscarMenor(arr)
// 		novoArr = append(novoArr, arr[menor])

// 		arr = append(arr[:menor], arr[menor+1:]...)
// 	}

// 	fmt.Println(novoArr)
// }

/*------------RECURSIVA (out-of-place)---------------*/

// func minIndexRec(arr []int, i int, minIdx int) int {
// 	if i >= len(arr) {
// 		return minIdx
// 	}

// 	if arr[i] < arr[minIdx] {
// 		minIdx = i
// 	}

// 	return minIndexRec(arr, i+1, minIdx)
// }

// func orderArrByMinToMax(arr *[]int, newArr *[]int) []int {
// 	if len(*arr) != 0 {
// 		minIdx := minIndexRec(*arr, 1, 0)

// 		*newArr = append(*newArr, (*arr)[minIdx])

// 		*arr = append((*arr)[:minIdx], (*arr)[minIdx+1:]...)

// 		orderArrByMinToMax(arr, newArr)
// 	}

// 	return *newArr
// }

// func main() {
// 	arr := []int{10, 23, 1, 215, 1612, 23, 15, 612, 123, 12}
// 	novoArr := []int{}

// 	arr = orderArrByMinToMax(&arr, &novoArr)

// 	fmt.Println("novoArr:", novoArr)
// 	fmt.Println("arr (restou):", arr)
// }

/*------------FIRST VERSION---------------*/

func main() {
	arr := []int{12, 213, 5, 5134, 34, 231, 2, 13, 13}

	pointer := 0
	for pointer < len(arr) {
		menor := pointer

		for i := pointer + 1; i < len(arr); i++ {
			if arr[i] < arr[menor] {
				menor = i
			}
		}

		if arr[pointer] != arr[menor] {
			arr[pointer], arr[menor] = arr[menor], arr[pointer]
		}
		pointer++
	}

	fmt.Print(arr)
}
