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

func abs(in int) int {
	if in < 0 {
		return -in
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

func parseInstruction(in instruction, s *ship, wp *coords) {
	switch in.dir {
	case 'N':
		wp.y += in.dist
	case 'E':
		wp.x += in.dist
	case 'S':
		wp.y -= in.dist
	case 'W':
		wp.x -= in.dist
	case 'R':
		switch in.dist {
		case 90:
			wp.x, wp.y = +wp.y, -wp.x
		case 180:
			wp.x, wp.y = -wp.x, -wp.y
		case 270:
			wp.x, wp.y = -wp.y, +wp.x
		}
	case 'L':
		switch in.dist {
		case 90:
			wp.x, wp.y = -wp.y, +wp.x
		case 180:
			wp.x, wp.y = -wp.x, -wp.y
		case 270:
			wp.x, wp.y = +wp.y, -wp.x
		}
	case 'F':
		s.pos.x += wp.x * in.dist
		s.pos.y += wp.y * in.dist
	default:
		fmt.Println("Unknown direction, exiting.\n")
		os.Exit(2)
	}
}

func parseRoute(in []instruction, s *ship, wp *coords) {
	for _, inst := range in {
		parseInstruction(inst, s, wp)
		fmt.Printf("%s\t%d\t%d   \t%d\n", string(inst.dir), inst.dist, s.pos, wp)
	}
}

func main() {
	fn := "input/12.txt"
	instructions := fileToArray(fn)

	minnow := &ship{
		head: 'E',
		pos:  coords{0, 0},
	}

	wp := &coords{
		x: 10,
		y: 1,
	}

	parseRoute(instructions, minnow, wp)

	fmt.Printf("Minnow Position: %d\n", minnow.pos)
	fmt.Printf("Answer: %d\n", abs(minnow.pos.x)+abs(minnow.pos.y))
}
