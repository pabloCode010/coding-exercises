package main

import "fmt"

func solution(matrix [][]int) int {
	sum := 0
	rows := len(matrix)
	columns := len(matrix[0])

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			val := matrix[row][column]
			sum += val

			if val == 0 {
				break
			}
		}
	}

	return sum
}

func main() {
	matrix := [][]int{
		{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3},
	}

	fmt.Println(solution(matrix))
}
