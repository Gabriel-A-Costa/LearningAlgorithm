package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	arr = append(arr[:1], arr[1+1:]...)
	fmt.Println(arr)
}
