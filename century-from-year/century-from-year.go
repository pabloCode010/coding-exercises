package main

import "fmt"

func centuryFromYear(year int) int {
	return (year-1)/100 + 1
}

func main() {
	years := []int{1900, 2001, 1850, 2000, 1799}

	for _, y := range years {
		fmt.Printf("year: %v century: %v\n", y, centuryFromYear(y))
	}
}
