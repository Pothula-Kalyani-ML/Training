package main

import (
	"fmt"
)

func insertionSort(listString []string) []string {
	i := 1
	var equalIndex []string
	for i < len(listString) {
		if listString[i] < listString[i-1] {
			listString[i], listString[i-1] = listString[i-1], listString[i]
			for j := i - 1; j > 0; j-- {
				if listString[j] < listString[j-1] {
					listString[j], listString[j-1] = listString[j-1], listString[j]
				} else if listString[j] == listString[j-1] {
					equalIndex = append(equalIndex, listString[j-1])
					break
				} else {
					break
				}
			}
		} else if listString[i-1] == listString[i] {
			equalIndex = append(equalIndex, listString[i-1])
		}
		i += 1

	}
	for j := 0; j < len(equalIndex); j++ {
		for i := len(listString) - 1; i >= 0; i-- {
			if listString[i] == equalIndex[j] {
				listString = append(listString[:i], listString[i+1:]...)
				break
			}
		}
	}
	return listString
}
func main() {
	listString := []string{"d2", "d2", "c", "b", "a2", "c", "d", "d2", "l", "l", "a2", "d2", "a", "a"}
	result := insertionSort(listString)
	fmt.Println(result)
}
