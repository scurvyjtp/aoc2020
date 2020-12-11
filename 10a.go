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
	for i := 1; i <= 999; i++ {
		for _, k := range vals {
			//fmt.Println(k, a+i)
			if k == a+i {
				return k
			}
		}
	}
	return 999
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

func countDiff(list []int) {
	counts := make(map[int]int)

	for n, k := range list {
		if n != len(list)-1 {
			counts[list[n+1]-k] += 1
		}
	}

	fmt.Printf("Answer: %d (1cnt)* %d (3cnt) = %d\n", counts[1], counts[3], counts[1]*counts[3])
}

func main() {
	var fn = "input/10.txt"
	vals := fileToArray(fn)
	myC := findMax(vals) + 3
	vals = append(vals, myC)
	var ordered []int
	a := 0
	ordered = append(ordered, a)
	for a < myC {
		a = findNext(vals, a)
		ordered = append(ordered, a)
	}
	countDiff(ordered)
}
