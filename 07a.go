package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) map[string]map[string]int {
	var l = map[string]map[string]int{}

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		a := parseSentence(scanner.Text())
		for k, v := range a {
			l[k] = v
		}

	}

	file.Close()
	return l
}

func parseSentence(in string) map[string]map[string]int {
	var ret = map[string]map[string]int{}

	r := regexp.MustCompile(`^(.*)bags contain(.*)$`)
	q := r.FindStringSubmatch(in)
	parent := q[1]
	ret[parent] = map[string]int{}

	if q[2] == " no other bags." {
		return ret
	}

	children := strings.Split(q[2], ",")

	for _, child := range children {
		r1 := regexp.MustCompile(`([0-9]+) ([a-z ]+) bag`)
		q1 := r1.FindStringSubmatch(child)
		n, _ := strconv.Atoi(q1[1])
		ret[parent][q1[2]] = n
	}

	return ret
}

func checkArr(in []string, val string) bool {
	for _, k := range in {
		if k == val {
			return true
		}
	}
	return false
}

func distinctArray(in []string) []string {
	var out = []string{}
	for _, k := range in {
		if !(checkArr(out, k)) {
			out = append(out, k)
		}
	}
	return out
}

func parseBags(bags map[string]map[string]int, val string) []string {
	var pBags []string
	count := 0
	val = val[:(len(val) - 1)]

	for k, v := range bags {
		count += 1
		for k1, _ := range v {
			if val == k1 {
				pBags = append(pBags, k)
				pBags = append(pBags, parseBags(bags, k)...)
			}
		}
	}
	return pBags
}

func main() {
	var fn = "input/07.txt"
	myBag := "shiny gold\000"

	nestedBags := fileToArray(fn)

	ans := parseBags(nestedBags, myBag)
	ans = distinctArray(ans)
	fmt.Printf("Answer: %d\n", len(ans))
}
