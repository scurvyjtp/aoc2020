package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "regexp"
)

type pass struct {
    min         int
    max         int
    checkVal    string
    passVal     string
}


func check (e error) {
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
        if (checkPass(p)) {
            ret += 1
        }
    }

    file.Close()
    return ret
}

func parsePass(l string) pass {
    var ret pass

    r:= regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+)\s(?P<checkVal>\w):\s(?P<passVal>\w+)$`)
    q:= r.FindStringSubmatch(l)

    ret.min, _ = strconv.Atoi(q[1])
    ret.max, _ = strconv.Atoi(q[2])
    ret.checkVal = q[3]
    ret.passVal = q[4]

    return ret
}

func checkPass(p pass) bool {
    count := 0

    for _, w := range p.passVal {
        if string(w) == p.checkVal {
            count += 1
        }
    }

    if count >= p.min && count <= p.max {
        return true
    }

    return false
}

func main() {
    var fn = "input/02.txt"
    vals := fileToArray(fn)
    fmt.Println(vals)
}
