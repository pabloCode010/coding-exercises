package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)

	j := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

// func removeElement(nums []int, val int) int {
// 	length := len(nums)
// 	deletes := 0

// 	for i := 0; i < length-deletes; i++ {
// 		if nums[i] == val {
// 			nums[i], nums[length-1-deletes] = nums[length-1-deletes], nums[i]
// 			deletes++
// 			i--
// 		}
// 	}

// 	return length - deletes
// }

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	n := removeElement(nums, 2)

	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}
