package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
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
	//y := x
	diff := x.last - x.first
	dist := diff / 2
	if v == "B" || v == "R" {
		x.first = x.first + dist
	} else if v == "F" || v == "L" {
		x.last = x.last - dist
	}
	//fmt.Printf("\t%d\t%s\t%d\n", y, v, x)
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

func findSeat(seats [128][8]int) int {
	ans := 0
	for i := 0; i < 128; i++ {
		for j := 1; j < 7; j++ { // cheated here (not a window seat)
			if seats[i][j] == 0 && seats[i][j+1] != 0 && seats[i][j-1] != 0 {
				ans = getTktId(i, j)
			}
		}
	}
	return ans
}

func main() {
	//128 rows
	//8 cols
	var fn = "input/05.txt"
	vals := fileToArray(fn)
	var seats [128][8]int

	for _, line := range vals {
		val := parseTkt(line)
		id := getTktId(val.first, val.last)
		seats[val.first][val.last] = id
	}

	fmt.Println("Answer: ", findSeat(seats))
}
