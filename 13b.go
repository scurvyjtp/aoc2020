package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var lines []string

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func parseBuses(line string) map[int]int {
	var out = make(map[int]int)

	for n, bus := range strings.Split(line, ",") {
		if bus != "x" {
			b, _ := strconv.Atoi(bus)
			out[b] = n
		}
	}
	return out
}

func findZero(b map[int]int) int {
	for k, v := range b {
		if v == 0 {
			return k
		}
	}
	os.Exit(2)
	return 0
}

func findLargest(b map[int]int) int {
	o := 0
	for k, _ := range b {
		if k >= o {
			o = k
		}
	}
	return o
}

func findFirst(b map[int]int) int {
	x := findZero(b)
	y := findLargest(b)

	n := (100000000000000 / y) * y
	//n := 0
	for {
		check := true
		n += y
		m := n - b[y]

		if m%x != 0 {
			continue
		}
		//fmt.Println(m)
		for k, v := range b {
			if ((m + v) % k) != 0 {
				check = false
				break
			}
		}

		if check {
			return m
		}

	}
}

func findBus(sched map[int]int, target int) int {
	var out int
	for k, v := range sched {
		if v < target {
			out = k
			target = v
		}
	}
	return out
}

func main() {
	var fn = "input/13.txt"

	lines := fileToArray(fn)

	buses := parseBuses(lines[1])

	//fmt.Println(buses)
	fmt.Printf("Answer: %d\n", findFirst(buses))

}
