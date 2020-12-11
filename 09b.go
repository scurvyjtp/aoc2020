package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToArray(fn string) []int {
	var l []int

	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		l = append(l, n)
	}

	file.Close()
	return l
}

func checkSlice(arr []int, val int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if val == arr[i]+arr[j] {
				//fmt.Printf("%d + %d = %d\n", arr[i], arr[j], val)
				return true
			}
		}
	}
	return false
}

func findExploit(vals []int) int {
	endVal := 25
	for i := 0; i < len(vals)-endVal; i++ {
		if !(checkSlice(vals[i:(i+endVal)], vals[i+endVal])) {
			return vals[i+endVal]
		}
	}
	return 0
}

func findMax(in []int) int {
	x := 0
	for _, y := range in {
		if y > x {
			x = y
		}
	}
	return x
}

func findMin(in []int, x int) int {
	for _, y := range in {
		if y < x {
			x = y
		}
	}
	return x
}

func findWeakness(vals []int, ans int) {
	running_total := 0
	for i := 0; i < len(vals); i++ {
		running_total += vals[i]
		if running_total == ans {
			max := findMax(vals[0:i])
			min := findMin(vals[0:i], ans)
			fmt.Printf("Answer: %d + %d = %d\n", min, max, min+max)
			os.Exit(0)
		}
	}
}

func main() {
	var fn = "input/09.txt"
	vals := fileToArray(fn)
	ans := findExploit(vals)
	for i := 0; i < len(vals); i++ {
		findWeakness(vals[i:], ans)
	}
	//fmt.Printf("Answer: %d\n", ans)
}
