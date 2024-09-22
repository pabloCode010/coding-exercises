package main

import "fmt"

func solution(sequence []int) bool {
	count := 0

	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			count++
			if i > 0 && i < len(sequence)-2 && sequence[i] >= sequence[i+2] && sequence[i-1] >= sequence[i+1] {
				return false
			}
		}
	}

	return count <= 1
}

func main() {
	result := solution([]int{1, 3, 2, 1})
	fmt.Println(result)
}
