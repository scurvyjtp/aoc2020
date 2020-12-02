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

func checkArray(vals []int) {
	base := vals[0]
	l := len(vals)

	for i := 1; i < l; i++ {
		s := base + vals[i]
		p := base * vals[i]
		checkArrayV(vals[1:], s, p)
	}
}

func checkArrayV(vals []int, sum int, prod int) {
	l := len(vals)

	for i := 0; i < l; i++ {
		if sum+vals[i] == 2020 {
			fmt.Printf("%d\n", (prod * vals[i]))
			os.Exit(1)
		}
	}
}

func main() {
	var fn = "input/01.txt"
	vals := fileToArray(fn)

	for n := range vals {
		checkArray(vals[n:])
	}
}
