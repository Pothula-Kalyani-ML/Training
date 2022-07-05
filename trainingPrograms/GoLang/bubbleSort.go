package main

import (
	"fmt"
)

func bubbleSort(list []int) []int {
	length := len(list)
	for length > 0 {
		var sequenceChange bool
		for i := range list {
			sequenceChange = false
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sequenceChange = true
			}

		}
		length = length - 1
		if sequenceChange == false {
			break
		}

	}

	return list
}
func main() {
	list := []int{2, 9, 6, 7, 6, 3, 2, 4, 3, 10, 18, 90, 64, 21}

	result := bubbleSort(list)
	fmt.Println(result)
}
