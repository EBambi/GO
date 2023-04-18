package main

import "fmt"

func main() {
	nums := []int{23, 2, 4, 1}
	fmt.Println(nums)

	nums = append(nums, 9)
	fmt.Println(nums)

	for i, nums := range nums {
		fmt.Println(i, nums)
	}
}
