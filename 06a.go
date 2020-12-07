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
	var x string
	//count := 0
	tString := ""

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		x = scanner.Text()

		if x == "" {
			ret = append(ret, tString)
			//count += 1
			tString = x
			continue
		}
		tString += x

	}
	ret = append(ret, tString)

	file.Close()
	return ret
}

func scoreGroup(in string) int {
	t := make(map[string]bool)
	for _, v := range in {
		t[string(v)] = true
	}

	return len(t)
}

func addArray(arr []int) int {
	var a int
	for _, d := range arr {
		a += d
	}
	return a
}

func main() {
	var fn = "input/06.txt"
	var scores []int

	groups := fileToArray(fn)

	for _, g := range groups {
		scores = append(scores, scoreGroup(g))
	}
	fmt.Printf("Answer: %d\n", addArray(scores))
}
