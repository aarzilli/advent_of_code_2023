package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1 := 0
	for _, line := range lines {
		part1 += process(Vatoi(Spac(line, " ", -1)))
	}
	Sol(part1)
	
	part2 := 0
	for _, line := range lines {
		part2 += process2(Vatoi(Spac(line, " ", -1)))
	}
	Sol(part2)
}

func process(v []int) int {
	//pln(v)
	ok := true
	for i := range v {
		if v[i] != 0 {
			ok = false
			break
		}
	}
	if ok {
		return 0
	}
	v2 := []int{}
	for i := 1; i < len(v); i++ {
		v2 = append(v2, v[i] - v[i-1])
	}
	r := process(v2)
	//pln("-> ", v[len(v)-1] + r)
	return v[len(v)-1] + r
	
}

func process2(v []int) int {
	//pln(v)
	ok := true
	for i := range v {
		if v[i] != 0 {
			ok = false
			break
		}
	}
	if ok {
		return 0
	}
	v2 := []int{}
	for i := 1; i < len(v); i++ {
		v2 = append(v2, v[i] - v[i-1])
	}
	r := process2(v2)
	//pln("-> ", v[0] - r)
	return v[0] - r
	
}