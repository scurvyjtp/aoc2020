package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var ret []string
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	file.Close()
	return ret
}

func rideSlope(f []string, s []int) int {
	ans := 0

	for line, trees := range f {
		mod := len(trees)
		if line%s[1] == 0 {
			if trees[((s[0]*(line/s[1]))%mod)] == '#' {
				ans += 1
			}
		}
	}
	return ans
}

func mltArray(in []int) int {
	a := 1

	for _, d := range in {
		a *= d
	}
	return a
}

func main() {
	var fn = "input/03.txt"
	var ans []int
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	forrest := fileToArray(fn)

	for _, s := range slopes {
		ans = append(ans, rideSlope(forrest, s))
	}

	fmt.Printf("Answer: %d\n", mltArray(ans))
}
