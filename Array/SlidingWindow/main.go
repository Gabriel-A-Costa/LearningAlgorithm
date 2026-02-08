package main

import "fmt"

func SlidingWindow(arr []string) (int, []string) {
	_max := 1
	subArray := []string{}
	info := make(map[string]int)
	l, start, end := 0, 0, 0
	for r := 0; r <= len(arr)-1; r++ {
		info[arr[r]]++
		for info[arr[r]] >= 3 {
			info[arr[l]]--
			l++
		}

		if cur := r - l + 1; cur > _max {
			start = l
			end = r
			_max = cur
			subArray = arr[start : end+1]
		}

	}

	return _max, subArray
}

func main() {
	arr := []string{"a", "v", "b", "b", "a", "c", "b", "b"}

	max, subArray := SlidingWindow(arr)

	fmt.Println(max)
	fmt.Println(subArray)
}
