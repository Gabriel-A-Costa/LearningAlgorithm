package main

import "fmt"

func main() {
	x := "saca"
	y := "cta"

	if len(x) != len(y) {
		fmt.Println("X não é anagrama de Y")
		return
	}

	arr := [26]int{}

	for i := 0; i < len(x); i++ {
		arr[x[i]-'a']++
		arr[y[i]-'a']--
	}

	for _, val := range arr {
		if val != 0 {
			fmt.Println("X não é anagrama de Y")
			return
		}
	}

	fmt.Println("X é anagrama de Y")
}
