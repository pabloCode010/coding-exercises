package main

import "fmt"

func solution(inputArray []string) []string {
	longestStrings := make([]string, 0)
	maxLength := 0

	for _, str := range inputArray {
		length := len(str)

		if length < maxLength {
			continue

		} else if length > maxLength {
			maxLength = length
			longestStrings = longestStrings[:0]
		}

		longestStrings = append(longestStrings, str)
	}

	return longestStrings
}

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}
	longestStrings := solution(inputArray)
	fmt.Println(longestStrings)
}
