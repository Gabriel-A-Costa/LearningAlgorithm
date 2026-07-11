package main

import (
	"fmt"
)

func frequency_map(s string) map[string]int {
	mapa := make(map[string]int)

	for _, ch := range s {
		mapa[string(ch)] += 1
	}

	return mapa
}

func main() {
	mapa := frequency_map("Hello World!")

	fmt.Println(mapa)
}
