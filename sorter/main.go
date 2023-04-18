package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arrayToSortS := os.Args[1:]
	var arrayToSort []int

	for _, i := range arrayToSortS {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arrayToSort = append(arrayToSort, j)
	}

	fmt.Println("Select sorting algorithm:\n1. Bubble Sort \n2. Insertion Sort\n3. Selection Sort")
	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		BubbleSort(arrayToSort)
	case 2:
		InsertionSort(arrayToSort)
	case 3:
		SelectionSort(arrayToSort)
	}
}
