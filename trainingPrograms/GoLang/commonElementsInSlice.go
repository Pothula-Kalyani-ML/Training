package main

import (
	"fmt"
)

func commonelements(list1 []int, list2 []int) []int {
	var commonlist []int
	for _, item1 := range list1 {
		for _, item2 := range list2 {

			if item1 == item2 {
				var addInCommon bool = true
				for _, itemcommon := range commonlist {
					if itemcommon == item1 {
						addInCommon = false
						break
					}

				}
				if addInCommon {
					commonlist = append(commonlist, item1)
					break
				}

			}
		}

	}
	return commonlist
}

func main() {

	list1 := []int{7, 7, 1, 5, 2, 3, 6}
	list2 := []int{3, 8, 6, 20, 7, 7}
	// expected output [7,3,6]
	result := commonelements(list1, list2)
	fmt.Println(result)
}
