package main

import (
	"bufio"
	"fmt"
	"os"
)

type block struct {
	first int
	last  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var l []string

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}

	file.Close()
	return l
}

func splitTkt(v string, x block) block {
	diff := x.last - x.first
	dist := diff / 2
	if v == "B" || v == "R" {
		x.first = x.first + dist
	} else if v == "F" || v == "L" {
		x.last = x.last - dist
	}
	return x
}

func parseTkt(line string) block {
	rows := block{0, 128}
	cols := block{0, 8}

	r := line[0:7]
	c := line[7:]

	for _, v := range r {
		rows = splitTkt(string(v), rows)
	}
	for _, v := range c {
		cols = splitTkt(string(v), cols)
	}
	return block{rows.first, cols.first}
}

func getTktId(row int, col int) int {
	return (row * 8) + col
}

func getMaxVal(v []int) int {
	ans := 0
	for _, x := range v {
		if x > ans {
			ans = x
		}
	}
	return ans
}

func main() {
	var fn = "input/05.txt"
	vals := fileToArray(fn)
	var seatVals []int

	for _, line := range vals {
		val := parseTkt(line)
		id := getTktId(val.first, val.last)
		seatVals = append(seatVals, id)
	}

	fmt.Println("Answer: ", getMaxVal(seatVals))
}
