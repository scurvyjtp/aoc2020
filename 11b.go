package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	//"strconv"
)

type coords struct {
	x int
	y int
}

type seats struct {
	w     int
	h     int
	cells map[coords]rune
}

var runeMap = map[rune]int{
	'L': 0,
	'.': 0,
	'#': 1,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var out []string

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}

func initSeats(in []string) *seats {
	first := &seats{
		h:     len(in),
		w:     len(in[0]),
		cells: make(map[coords]rune, len(in)*len(in[0])),
	}

	for y, inLine := range in {
		for x, cell := range inLine {
			first.cells[coords{x, y}] = cell
		}
	}
	return first
}

func nextState(s *seats) (newSeats *seats, change int) {
	newSeats = &seats{
		h:     s.h,
		w:     s.w,
		cells: make(map[coords]rune, s.w*s.h),
	}

	for x := 0; x < s.w; x++ {
		for y := 0; y < s.h; y++ {
			cur := coords{x, y}
			switch s.cells[cur] {
			case '.':
				newSeats.cells[cur] = '.'
			case 'L':
				if occupiedAdjacent(s, cur) == 0 {
					newSeats.cells[cur] = '#'
					change += 1
				} else {
					newSeats.cells[cur] = 'L'
				}
			case '#':
				if occupiedAdjacent(s, cur) >= 5 {
					newSeats.cells[cur] = 'L'
					change += 1
				} else {
					newSeats.cells[cur] = '#'
				}
			default:
				fmt.Printf("Error: unknown value at %d\n", cur)
			}

		}
	}
	return newSeats, change
}

func occupiedAdjacent(s *seats, c coords) int {
	var count int

	count += runeMap[findClosest(s, c, -1, -1)]
	count += runeMap[findClosest(s, c, -1, 0)]
	count += runeMap[findClosest(s, c, -1, 1)]

	count += runeMap[findClosest(s, c, 0, -1)]
	count += runeMap[findClosest(s, c, 0, 1)]

	count += runeMap[findClosest(s, c, 1, -1)]
	count += runeMap[findClosest(s, c, 1, 0)]
	count += runeMap[findClosest(s, c, 1, 1)]

	return count
}

func findClosest(s *seats, start coords, xd int, yd int) rune {
	x := start.x + xd
	y := start.y + yd

	for {
		if x < 0 || x >= s.w || y < 0 || y >= s.h {
			return '0'
		}

		r := s.cells[coords{x, y}]
		if r == '.' {
			x += xd
			y += yd
			continue
		}
		return r
	}
	return '0'
}

func countSeats(s *seats) int {
	var count int
	for x := 0; x < s.w; x++ {
		for y := 0; y < s.h; y++ {
			count += runeMap[s.cells[coords{x, y}]]
		}
	}
	return count
}

func printMap(s map[coords]rune) {
	for y := 0; y < 10; y++ {
		fmt.Printf("\n")
		for x := 0; x < 10; x++ {
			fmt.Printf(string(s[coords{x, y}]))
		}
	}
	fmt.Print("\n")
}

func main() {
	var fn = "input/11.txt"
	dataIn := fileToArray(fn)
	seats := initSeats(dataIn)

	state, change := nextState(seats)
	count := 1

	for {
		state, change = nextState(state)
		//printMap(state.cells)
		if change == 0 {
			fmt.Printf("Answer: %d\n", countSeats(state))
			os.Exit(0)
		}
		count += 1
	}
}
