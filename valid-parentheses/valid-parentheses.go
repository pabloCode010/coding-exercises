package main

import (
	"fmt"
)

func validParentheses(s string) bool {
	stack := make([]byte, 0, 5)
	brackets := map[byte]byte{'(': ')', '{': '}', '[': ']'}
	var char, isClosed, isOpen = byte(0), false, false

	for i := 0; i < len(s); i++ {
		char = s[i]
		isClosed = char == ')' || char == ']' || char == '}'
		_, isOpen = brackets[char]

		if isOpen {
			stack = append(stack, char)

		} else if isClosed && len(stack) == 0 {
			return false

		} else if isClosed && brackets[stack[len(stack)-1]] != char {
			return false

		} else if isClosed {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s := "([{ hello world! }])"
	valid := validParentheses(s)

	fmt.Println(valid)
}
