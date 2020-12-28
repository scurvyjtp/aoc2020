package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := make(map[int]int)
	input := "1,17,0,10,18,11,6"
	turns := 30000000
	var last, next int

	orig := strings.Split(input, ",")

	for n, val := range orig {
		intVal, _ := strconv.Atoi(val)
		nums[intVal] = n + 1
		last = intVal
	}

	for turn := len(nums) + 1; turn <= turns; turn++ {
		if val, ok := nums[last]; ok {
			if turn-1 == val {
				next = 0
			} else {
				next = (turn - 1) - nums[last]
				nums[last] = turn - 1
			}
		} else {
			nums[last] = turn - 1
			next = 0
		}
		last = next
	}

	fmt.Printf("Answer: %d\n", next)
}
