package main

import (
	"fmt"
	"sort"
)

func solution(statues []int) (need int) {
	sort.IntSlice.Sort(statues)

	for i := 0; i < len(statues)-1; i++ {
		need += statues[i+1] - statues[i] - 1
	}

	return need
}

func main() {
	need := solution([]int{6, 2, 3, 8})
	fmt.Println(need)
}
