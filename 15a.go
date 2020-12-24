package main

import (
	//"bufio"
	"fmt"
	//"os"
	"strconv"
    "strings"
)

func sToi (in []string) []int {
    var out []int
    for _, v := range in {
        t, _ := strconv.Atoi(v)
        out = append(out,t)
    }
    return out
}

func getAge(inArr []int ) int {
    next := inArr[len(inArr)-1]
    for i := len(inArr)-2; i >=0; i-- {
        if inArr[i] == next {
            return (len(inArr) - (i+1))
        }
    }
    return 0
}


func main() {
	input := "1,17,0,10,18,11,6"
    turns := 2020
    orig := strings.Split(input, ",")
    origI := sToi(orig)

    for count := len(origI); count != turns; count++ {
        origI = append(origI, getAge(origI))
    }

    fmt.Printf("Answer: %d\n", origI[len(origI)-1])
}

