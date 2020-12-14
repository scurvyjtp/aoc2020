package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	"strconv"
)

func main() {
	var fn = "input/10.txt"
	var adapters = make(map[int]bool)
	var max = 0

	file, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		adapters[n] = true

		if n > max {
			max = n
		}
	}

	adapters[max+3] = true
	adapters[0] = true

	last := 0
	var d [4]int
	var c = [200]int{0}
	c[max] = 1

	d[1] += 1 // need to initialize count of 1
	for i := max; i >= 0; i -= 1 {

		if adapters[i] {
			c[i] += c[i+1]
			c[i] += c[i+2]
			c[i] += c[i+3]

			if last != 0 {
				diff := last - i
				d[diff] += 1
			}
			last = i
		}
	}
	fmt.Printf("Answer Part 1: %d * %d = %d\n", d[1], d[3], (d[1] * d[3]))
	fmt.Printf("Answer Part 2: %d\n", c[0])
}
