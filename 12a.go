package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type coords struct {
	x int
	y int
}

type instruction struct {
	dir  rune
	dist int
}

type ship struct {
	head rune
	pos  coords
}

var hToDeg = map[rune]int{
	'N': 0,
	'E': 90,
	'S': 180,
	'W': 270,
}

var dToHead = map[int]rune{
	0:   'N',
	90:  'E',
	180: 'S',
	270: 'W',
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func absCompass(in int) int {
	if in < 0 {
		return 360 + in
	}
	return in
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(line string) instruction {
	var out instruction
	out.dir = rune(line[0])

	n, e := strconv.Atoi(line[1:])
	check(e)
	out.dist = n

	return out
}

func fileToArray(fn string) []instruction {
	var out []instruction

	f, e := os.Open(fn)
	check(e)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out = append(out, parseLine(scanner.Text()))
	}
	return out
}

func parseInstruction(in instruction, s *ship) {
	switch in.dir {
	case 'N':
		s.pos.y += in.dist
	case 'E':
		s.pos.x += in.dist
	case 'S':
		s.pos.y -= in.dist
	case 'W':
		s.pos.x -= in.dist
	case 'R':
		s.head = dToHead[(absCompass(hToDeg[s.head]+in.dist) % 360)]
	case 'L':
		s.head = dToHead[(absCompass(hToDeg[s.head]-in.dist) % 360)]
	case 'F':
		parseInstruction(instruction{s.head, in.dist}, s)
	default:
		fmt.Println("Unknown direction, exiting.\n")
		os.Exit(2)
	}
}

func parseRoute(in []instruction, s *ship) {
	for _, inst := range in {
		parseInstruction(inst, s)
	}
}

func main() {
	fn := "input/12.txt"
	instructions := fileToArray(fn)

	minnow := &ship{
		head: 'E',
		pos:  coords{0, 0},
	}

	parseRoute(instructions, minnow)

	fmt.Printf("Minnow Position: %d\n", minnow.pos)
	fmt.Printf("Answer: %d\n", abs(minnow.pos.x)+abs(minnow.pos.y))
}
