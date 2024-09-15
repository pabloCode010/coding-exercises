package main

import (
	"fmt"
)

func solution(s1 string, s2 string) int {
	count := 0
	counterS1 := make(map[rune]int)
	counterS2 := make(map[rune]int)

	for _, char := range s1 {
		counterS1[char] += 1
	}

	for _, char := range s2 {
		counterS2[char] += 1
	}

	for char, value1 := range counterS1 {
		value2 := counterS2[char]
		count += min(value1, value2)
	}

	return count
}

func main() {
	fmt.Println(solution("aabcc", "adcaa"))
}
