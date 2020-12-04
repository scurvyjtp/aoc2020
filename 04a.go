package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile(fn string) int {
	file, err := os.Open(fn)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	recCount := 0
	ans := 0

	p := make(map[string]string)
	for scanner.Scan() {
		s := scanner.Text()

		if len(s) == 0 {
			if checkPassport(p) {
				ans += 1
			}
			fmt.Println(recCount, p, checkPassport(p))

			recCount += 1
			p = make(map[string]string)
			continue
		}

		for _, kv := range strings.Split(s, " ") {
			kvpair := strings.Split(kv, ":")
			p[kvpair[0]] = kvpair[1]
		}

	}

	// one more time for the last line
	if checkPassport(p) {
		ans += 1
	}
	fmt.Println(recCount, p, checkPassport(p))

	file.Close()
	return ans
}

func checkPassport(p map[string]string) bool {
	_, e1 := p["byr"]
	_, e2 := p["iyr"]
	_, e3 := p["eyr"]
	_, e4 := p["hgt"]
	_, e5 := p["hcl"]
	_, e6 := p["ecl"]
	_, e7 := p["pid"]
	//v8,e8 := p['cid']

	if e1 && e2 && e3 && e4 && e5 && e6 && e7 {
		return true
	}

	return false
}

func main() {
	var fn = "input/04.txt"
	val := parseFile(fn)
	fmt.Printf("Answer: %d\n", val)

}
