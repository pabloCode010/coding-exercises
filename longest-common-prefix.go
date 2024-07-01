package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := ""

	for i, char := range strs[0] {
		for _, word := range strs {
			if i >= len(word) || word[i] != byte(char) {
				return prefix
			}
		}

		prefix += string(char)
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)

	fmt.Println(prefix)
}
