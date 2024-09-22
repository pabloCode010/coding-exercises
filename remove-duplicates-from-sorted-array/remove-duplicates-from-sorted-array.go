package main

import "fmt"

func removeDuplicatesFromSortedArray(nums []int) int {
	length := len(nums)

	j := 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

func main() {
	nums := []int{1, 1, 2, 2, 2, 2, 3}
	n := removeDuplicatesFromSortedArray(nums)

	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}
