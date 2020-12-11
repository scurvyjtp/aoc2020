package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type command struct {
	inst string
	op   string
	val  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []command {
	var l []command
	var t command

	file, err := os.Open(fn)
	check(err)

	r := regexp.MustCompile(`^(.*) ([+|-])(.*)$`)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		q := r.FindStringSubmatch(scanner.Text())

		t.inst = q[1]
		t.op = q[2]
		n, _ := strconv.Atoi(q[3])
		t.val = n

		l = append(l, t)
	}

	file.Close()
	return l
}

func parseMath(acc int, op string, val int) int {
	if op == "+" {
		return (acc + val)
	}
	return (acc - val)
}

func checkArr(in []int, val int) bool {
	for _, v := range in {
		if v == val {
			return true
		}
	}
	return false
}

func runBoot(list []command) {
	acc := 0

	next := 0
	var visits []int

	for {
		if list[next].inst == "nop" {
			next += 1
		} else if list[next].inst == "jmp" {
			next = parseMath(next, list[next].op, list[next].val)
		} else if list[next].inst == "acc" {

			acc = parseMath(acc, list[next].op, list[next].val)
			next += 1
		}

		if checkArr(visits, next) {
			fmt.Printf("Answer: %d\n", acc)
			os.Exit(0)
		} else {
			visits = append(visits, next)
		}
	}
}

func main() {
	var fn = "input/08.txt"
	instructions := fileToArray(fn)
	runBoot(instructions)
}
