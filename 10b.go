package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []int {
	var l []int

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		l = append(l, n)
	}

	file.Close()
	return l
}

func findNext(vals []int, a int) int {
	return 0
}

func findMax(in []int) int {
	x := 0
	for _, y := range in {
		if y > x {
			x = y
		}
	}
	return x
}

func findAll(in []int, start int, end int) int {
	ret := 0
	for i := 0; i < len(in); i++ {
		if in[i] == start+1 || in[i] == start+2 || in[i] == start+3 {
			if in[i] == end {
				return (1)
				//fmt.Println("True")
			}
			ret += findAll(in, in[i], end)
		}
	}
	return ret
}

func main() {
	var fn = "input/10.txt"
	vals := fileToArray(fn)
	end := findMax(vals) + 3

	vals = append(vals, 0, end)

	fmt.Println(findAll(vals, 0, end))
}
