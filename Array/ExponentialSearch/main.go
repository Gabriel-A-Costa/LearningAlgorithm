package main

import "fmt"

func BinarySeach(arr []int, item int) (int, error) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (right + left) / 2
		value := arr[mid]

		if value == item {
			return mid, nil
		}

		if value < item {
			left = mid + 1
		}

		if value > item {
			right = mid - 1
		}
	}

	return 0, fmt.Errorf("Item não encontrado!")

}

func ExponentialSeach(arr []int, item int) (int, error) {
	if arr[0] == item {
		return 0, nil
	}

	p := 1
	maxIndex := len(arr) - 1

	for p < maxIndex && arr[p] < item {
		p *= 2
	}

	left := p / 2
	right := p
	if right > maxIndex {
		right = maxIndex
	}

	if arr[right] == item {
		return right, nil
	}

	index, err := BinarySeach(arr[left:right+1], item)
	if err != nil {
		return 0, err
	}

	return left + index, nil

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 12, 15, 18, 20, 30, 35, 50, 100}
	var item int

	fmt.Println("Digite um numero:")
	fmt.Scan(&item)

	index, err := ExponentialSeach(arr, item)

	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}

	fmt.Printf("O item %d esta no index %d\n", item, index)
}
