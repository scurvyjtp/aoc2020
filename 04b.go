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
	v1, e1 := p["byr"]
	v2, e2 := p["iyr"]
	v3, e3 := p["eyr"]
	v4, e4 := p["hgt"]
	v5, e5 := p["hcl"]
	v6, e6 := p["ecl"]
	v7, e7 := p["pid"]
	//v8,e8 := p['cid']
	if e1 && e2 && e3 && e4 && e5 && e6 && e7 {
		if checkByr(v1) && checkIyr(v2) && checkEyr(v3) && checkHgt(v4) && checkHcl(v5) && checkEcl(v6) && checkPid(v7) {
			return true
		}
	}

	return false
}

func checkByr(v1 string) bool {
	v, e := strconv.Atoi(v1)
	if e != nil {
		return false
	}
	if v >= 1920 && v <= 2002 {
		return true
	}
	return false
}

func checkIyr(v2 string) bool {
	v, e := strconv.Atoi(v2)
	if e != nil {
		return false
	}
	if v >= 2010 && v <= 2020 {
		return true
	}
	return false
}

func checkEyr(v3 string) bool {
	v, e := strconv.Atoi(v3)
	if e != nil {
		return false
	}
	if v >= 2020 && v <= 2030 {
		return true
	}
	return false
}

func checkHgt(v4 string) bool {
	r := regexp.MustCompile(`^(\d+)(cm|in)$`)
	q := r.FindStringSubmatch(v4)
	if q == nil {
		return false
	}
	h, _ := strconv.Atoi(q[1])
	if q[2] == "cm" && (h >= 150 && h <= 193) {
		return true
	} else if q[2] == "in" && (h >= 59 && h <= 76) {
		return true
	}
	return false
}

func checkHcl(v5 string) bool {
	if v5[0] == '#' && len(v5) == 7 {
		m, _ := regexp.MatchString("[0-9a-f]+", v5[1:])
		return m
	}

	return false
}

func checkEcl(v6 string) bool {
	if v6 == "amb" || v6 == "blu" || v6 == "brn" || v6 == "gry" || v6 == "grn" || v6 == "hzl" || v6 == "oth" {
		return true
	}
	return false
}

func checkPid(v7 string) bool {
	if len(v7) != 9 {
		return false
	}
	_, e := strconv.Atoi(v7)
	if e == nil {
		return true
	}
	return false
}

func main() {
	var fn = "input/04.txt"
	val := parseFile(fn)
	fmt.Printf("Answer: %d\n", val)

}
