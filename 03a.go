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

func main() {
	var fn = "input/03.txt"
	ans := 0

	forrest := fileToArray(fn)

	for n, s := range forrest {
		if s[((n*3)%31)] == '#' {
			ans += 1
		}
	}

	fmt.Printf("Answer: %d\n", ans)
}
