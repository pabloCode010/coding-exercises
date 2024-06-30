package main

import "fmt"

var values = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	result, prev := 0, 0

	for i := 0; i < len(s); i++ {
		current := values[s[i]]

		if prev < current {
			result -= prev * 2
		}

		result += current
		prev = current
	}

	return result
}

func main() {
	romanNum := "XIV"

	fmt.Printf("%v = %v\n", romanNum, romanToInt(romanNum))
}
