package main

import "fmt"

func isPalindrome(x int) bool {
	var number, reverse, digit int

	for number = x; number > 0; number /= 10 {
		digit = number % 10
		reverse = reverse*10 + digit
	}

	return x == reverse
}

func main() {
	x := 121

	if isPalindrome(x) {
		fmt.Printf("%d is palindrome number\n", x)
	} else {
		fmt.Printf("%d is not palindrome number\n", x)
	}
}
