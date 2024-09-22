package main

import (
	"fmt"
	"math"
)

func solution(number int) bool {
	digits, copy := 0, number
	part1, part2 := 0, 0

	for ; copy != 0; copy /= 10 {
		digits += 1
	}

	for i := 0; i < digits/2; i++ {
		pow := digits - 1 - i*2
		nextDigitPart1 := number / int(math.Pow10(pow))
		nextDigitPart2 := number % 10

		part1 += nextDigitPart1
		part2 += nextDigitPart2

		number -= nextDigitPart1 * int(math.Pow10(pow))
		number /= 10
	}

	return part1 == part2
}

func main() {
	fmt.Println(solution(1230))
}
