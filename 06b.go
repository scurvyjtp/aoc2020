package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []string {
	var ret, group []string
	var x string

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		x = scanner.Text()

		if x == "" {
			ret = append(ret, mergeGroup(group))
			group = nil
			continue
		}
		group = append(group, x)
	}

	ret = append(ret, mergeGroup(group))
	file.Close()
	return ret
}

func mergeGroup(in []string) string {
	first := true
	t := make(map[string]bool)
	ret := ""

	for _, v := range in {
		if first {
			for _, w := range v {
				t[string(w)] = true
			}
			first = false
		} else {
			for key, val := range t {
				if val {
					match, _ := regexp.MatchString(key, v)
					if !match {
						t[key] = false
					}
				}
			}
		}
	}

	for key, val := range t {
		if val {
			ret += key
		}
	}
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
