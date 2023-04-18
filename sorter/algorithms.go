package main

import "fmt"

func BubbleSort(n []int) {
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}
	fmt.Println(n)
}

func InsertionSort(n []int) {
	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}

		i++
	}

	fmt.Println(n)
}

func SelectionSort(n []int) {
	var i = 1
	for i < len(n)-1 {
		var j = i + 1
		var minIndex = i

		if j < len(n) {
			if n[j] < n[minIndex] {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			var temp = n[i]
			n[i] = n[minIndex]
			n[minIndex] = temp
		}

		i++
	}

	fmt.Println(n)
}
