package main

import "fmt"

func findIndex(haystack string, needle string) int {
	length := len(haystack)
	n := len(needle)

	for i := 0; i <= length-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "sadbutsad"
	needle := "sad"

	fmt.Println(findIndex(haystack, needle))
}
