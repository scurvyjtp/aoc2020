package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type interval struct {
	min int
	max int
}

type intervalArr []interval

func (in intervalArr) Len() int {
	return len(in)
}

func (in intervalArr) Less(a, b int) bool {
	return in[a].min < in[b].min
}

func (in intervalArr) Swap(a, b int) {
	in[a], in[b] = in[b], in[a]
}

type rule struct {
	label string
	rng   interval
}

func sArrtoi(in []string) []int {
	var out []int

	for _, v := range in {
		t, _ := strconv.Atoi(v)
		out = append(out, t)
	}
	return out
}

func readFile(fn string) ([]interval, [][]int) {
	var rules []interval
	var tickets [][]int

	file, err := os.Open(fn)
	check(err)
	r := regexp.MustCompile(`^(.*): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if r.MatchString(line) {
			m := r.FindStringSubmatch(line)

			m2, _ := strconv.Atoi(m[2])
			m3, _ := strconv.Atoi(m[3])
			m4, _ := strconv.Atoi(m[4])
			m5, _ := strconv.Atoi(m[5])

			rules = append(rules, interval{m2, m3})
			rules = append(rules, interval{m4, m5})

		} else {
			if ok, _ := regexp.MatchString(`[0-9]`, line); ok {
				ticket := sArrtoi(strings.Split(line, ","))
				tickets = append(tickets, ticket)
			}
		}
	}

	return rules, tickets
}

func mergeRanges(in []interval) []interval {
	var r []interval
	sort.Sort(intervalArr(in))

	a := &in[0]

	for i := 1; i < 40; i++ {
		b := &in[i]
		if a.max >= b.min {
			if a.max < b.max {
				a.max = b.max
			}
		} else {
			r = append(r, *a)
			a = b
		}
	}

	r = append(r, *a)
	return r

}

func checkVals(in int, ranges []interval) bool {
	for _, r := range ranges {
		if in >= r.min && in <= r.max {
			return true
		}
	}
	return false
}

func main() {
	fn := "input/16.txt"

	rules, tickets := readFile(fn)
	var eVals []int

	allValid := mergeRanges(rules)

	for _, tkt := range tickets {
		for _, v := range tkt {
			if !(checkVals(v, allValid)) {
				eVals = append(eVals, v)
			}
		}
	}

	runningTotal := 0
	for _, v := range eVals {
		runningTotal += v
	}
	fmt.Printf("Answer: %d\n", runningTotal)
}
