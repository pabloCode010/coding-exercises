package main

import "fmt"

func searchInsert(nums []int, target int) int {
	i, n := 0, len(nums)

	for ; i < n && nums[i] < target; i++ {}

	return i
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5

	n := searchInsert(nums, target)
	fmt.Println(n)
}
