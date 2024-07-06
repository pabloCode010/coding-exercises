package main

import "fmt"

func shapeArea(n int) int {
	area := 1

	for i := 1; i < n; i++ {
		area += 4 * i
	}

	return area
}

func main() {
	n := 3
	area := shapeArea(n)

	fmt.Println(area)
}
