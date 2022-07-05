package main

// func Search(n int, f func(int) bool) int
import (
	"fmt"
	"sort"
)

func binarySearch(list [15]int, searchItem int) int {

	var first, last, div_count, middle, n int
	n = len(list)
	last = n
	div := n / 2
	middle = div
	for div > 0 {
		div = div / 2
		div_count++
	}

	for div_count > 0 {
		if list[middle] == searchItem {
			break
		}
		if searchItem > list[middle] {
			first = middle + 1
			middle = middle + (last-middle)/2
		}
		if searchItem < list[middle] {
			first = middle - 1
			middle = middle - (middle-first)/2
		}
		div_count--
	}
	return middle
}
func main() {
	arraylist := [...]int{23, 46, 56, 76, 89, 87, 60, 55, 45, 32, 12, 22, 21, 43, 78}
	listSlice := arraylist[:]
	sort.Ints(listSlice)
	fmt.Println(arraylist)
	fmt.Println(listSlice)
	searchItem := 87
	index := binarySearch(arraylist, searchItem)
	fmt.Printf("index of %d  is %d ", searchItem, index)

}
