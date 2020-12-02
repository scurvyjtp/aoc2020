package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type pass struct {
	first    int
	second   int
	checkVal string
	passVal  string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) int {
	var p pass
	ret := 0

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		p = parsePass(scanner.Text())
		if checkPass(p) {
			ret += 1
		}
	}

	file.Close()
	return ret
}

func parsePass(l string) pass {
	var ret pass

	r := regexp.MustCompile(`^(?P<first>\d+)-(?P<second>\d+)\s(?P<checkVal>\w):\s(?P<passVal>\w+)$`)
	q := r.FindStringSubmatch(l)

	ret.first, _ = strconv.Atoi(q[1])
	ret.second, _ = strconv.Atoi(q[2])
	ret.checkVal = q[3]
	ret.passVal = q[4]

	return ret
}

func checkPass(p pass) bool {
	firstBool := string(p.passVal[p.first-1]) == p.checkVal
	secondBool := string(p.passVal[p.second-1]) == p.checkVal

	if (firstBool || secondBool) && !(firstBool && secondBool) {
		return true
	}

	return false
}

func main() {
	var fn = "input/02.txt"
	vals := fileToArray(fn)
	fmt.Println(vals)
}
