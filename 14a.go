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

func lpadZ(in string, length int) string {
	r := length - len(in)
	if r <= 0 {
		return in
	}
	return strings.Repeat("0", r) + in
}

func main() {
	var fn = "input/14.txt"

	memVals := make(map[int]string)
	maskRe := regexp.MustCompile(`^mask = (.*)$`)
	memRe := regexp.MustCompile(`^mem\[(.*)\] = (.*)`)

	var maskOne []int
	var maskZero []int
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		match, _ := regexp.MatchString("mem", line)

		if match {
			var ret []byte
			l := memRe.FindStringSubmatch(line)
			mem, _ := strconv.Atoi(l[1])
			val, _ := strconv.Atoi(l[2])
			valS := lpadZ(strconv.FormatInt(int64(val), 2), 36)

			ret = []byte(valS)

			for _, v := range maskOne {
				ret[v] = '1'
			}

			for _, v := range maskZero {
				ret[v] = '0'
			}

			memVals[mem] = string(ret)

		} else {
			maskOne = nil
			maskZero = nil

			l := maskRe.FindStringSubmatch(line)
			for num, zo := range l[1] {
				if zo == '1' {
					maskOne = append(maskOne, num)
				} else if zo == '0' {
					maskZero = append(maskZero, num)
				}
			}

			//fmt.Println(maskVal, len(maskVal))
			//fmt.Println("MaskOne: ", maskOne)
			//fmt.Println("MaskZero: ", maskZero)
		}
	}

	runningTotal := int64(0)
	for _, v := range memVals {
		i, _ := strconv.ParseInt(v, 2, 64)
		runningTotal += i
	}
	fmt.Printf("Answer: %d\n", runningTotal)
}
