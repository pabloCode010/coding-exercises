package main

import "fmt"

func solution(input string) string {
	var result string
	stack := make([]string, 0)

	for _, v := range input {
		switch v {
		case '(':
			stack = append(stack, result)
			result = ""

		case ')':
			lastElement := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = lastElement + reverse(result)

		default:
			result += string(v)
		}
	}

	return result
}

func reverse(input string) string {
	var result string

	for i := len(input) - 1; i >= 0; i-- {
		result += string(input[i])
	}

	return result
}

func main() {
	result := solution("Hello, (wo(rl)d)")
	fmt.Println(result)
}
