package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && n != 0 {
			return n
		} else if s[i] != ' ' {
			fmt.Printf("%c\n", s[i])
			n += 1
		}
	}

	return n
}

func main() {
	text := "Hello World  "
	n := lengthOfLastWord(text)

	fmt.Println("n =", n)
}
