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

func findExploit(vals []int) {
	endVal := 25
	for i := 0; i < len(vals)-endVal; i++ {
		if !(checkSlice(vals[i:(i+endVal)], vals[i+endVal])) {
			fmt.Printf("Answer: %d\n", vals[i+endVal])
			os.Exit(0)
		}
	}
}

func main() {
	var fn = "input/09.txt"
	vals := fileToArray(fn)
	findExploit(vals)
}
