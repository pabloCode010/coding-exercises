package main

import "fmt"

// Guaranteed constraints:
// 2 ≤ inputArray.length ≤ 10,
// -1000 ≤ inputArray[i] ≤ 1000.

func adjacentElementsProduct(n []int) int {
	max := n[0] * n[1]

	for i := 1; i < len(n)-1; i++ {
		if product := n[i] * n[i+1]; product > max {
			max = product
		}
	}

	return max
}

func main() {
	elements := []int{3, 6, -2, -5, 7, 3}
	maxProduct := adjacentElementsProduct(elements)

	fmt.Println(maxProduct)
}
