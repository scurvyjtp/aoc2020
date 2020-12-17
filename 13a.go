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

func parseBuses(line string) []int {
	var out []int

	for _, bus := range strings.Split(line, ",") {
		if bus != "x" {
			b, _ := strconv.Atoi(bus)
			out = append(out, b)
		}
	}
	return out
}

func findFirst(b int, t int) int {
	return t - (t % b) + b
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
	myTs, _ := strconv.Atoi(lines[0])
	buses := parseBuses(lines[1])
	sched := make(map[int]int)

	for _, b := range buses {
		sched[b] = findFirst(b, myTs)
	}

	targetBus := findBus(sched, myTs+1000)

	fmt.Printf("Answer: %d\n", ((sched[targetBus] - myTs) * targetBus))
}
