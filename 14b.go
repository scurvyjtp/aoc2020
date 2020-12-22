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

func arrVals(in []byte, val byte) []int {
	var out []int
	for n, v := range in {
		if v == val {
			out = append(out, n)
		}
	}
	return out
}

func applyMask(mask []int, addr []byte) []int {
	var out []int
	var n int64

	addr[mask[0]] = '0'

	n, _ = strconv.ParseInt(string(addr), 2, 64)

	out = append(out, int(n))
	if len(mask) > 1 {
		out = append(out, applyMask(mask[1:], addr)...)
	}

	addr[mask[0]] = '1'
	n, _ = strconv.ParseInt(string(addr), 2, 64)
	out = append(out, int(n))
	if len(mask) > 1 {
		out = append(out, applyMask(mask[1:], addr)...)
	}
	return out
}

func maskAddress(address []byte, mask []byte) []int {
	mOne := arrVals(mask, '1')
	mX := arrVals(mask, 'X')

	for _, v := range mOne {
		address[v] = '1'
	}

	allMem := applyMask(mX, address)
	return allMem
}

func main() {
	var fn = "input/14.txt"
	var curMask []byte
	var memVals = make(map[int]int)

	maskRe := regexp.MustCompile(`^mask = (.*)$`)
	memRe := regexp.MustCompile(`^mem\[(.*)\] = (.*)`)

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		isMem, _ := regexp.MatchString("mem", line)

		if isMem {
			l := memRe.FindStringSubmatch(line)
			lt, _ := strconv.Atoi(l[1])
			memBin := lpadZ(strconv.FormatInt(int64(lt), 2), 36)

			memArr := (maskAddress([]byte(memBin), curMask))
			val, _ := strconv.Atoi(l[2])

			for _, v := range memArr {
				memVals[v] = val
			}

		} else {
			l := maskRe.FindStringSubmatch(line)
			curMask = []byte(l[1])
		}
	}

	runningTotal := 0
	for _, v := range memVals {
		runningTotal += v
	}
	fmt.Printf("Answer: %d\n", runningTotal)
}
