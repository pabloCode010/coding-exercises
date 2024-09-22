package main

import (
	"fmt"
	"slices"
)

func solution(a []int) []int {
	heights := make([]int, 0)

	for _, v := range a {
		if v != -1 {
			heights = append(heights, v)
		}
	}

	slices.Sort(heights)

	for i, _ := range a {
		if a[i] != -1 {
			a[i] = heights[0]
			heights = heights[1:]
		}
	}

	return a
}

func main() {
	a := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	sorted := solution(a)

	fmt.Println(sorted)
}
