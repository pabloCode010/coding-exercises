package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j, y := range nums {
			if i == j {
				continue
			}

			if x+y == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func main() {
	target := 16
	nums := []int{1, 3, 8, 7, 9, 10}

	result := twoSum(nums, target)

	if len(result) > 0 {
		fmt.Printf("%v + %v = %v\n", nums[result[0]], nums[result[1]], target)
	}
}
