package main

import "fmt"

func revertText(s string) []rune {
	runes := []rune(s)

	// start marks the first index of the current word
	start := 0

	// i walks through the string; when we hit a space (or the end), we reverse the word [start, i-1]
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			left, right := start, i-1
			for left < right {
				runes[left], runes[right] = runes[right], runes[left]
				left++
				right--
			}
			// next word starts after the space
			start = i + 1
		}
	}

	return runes
}

func main() {
	text := "Hello Word"
	textRevert := revertText(text)
	fmt.Printf("%c", textRevert)
}
