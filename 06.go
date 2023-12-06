package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	times := Getints(lines[0], false)
	dists := Getints(lines[1], false)
	pln(times, dists)

	part1 := 1
	for i := range times {
		part1 *= enum(times[i], dists[i])
	}
	Sol(part1)

	// part2
	p2time := Getints(strings.ReplaceAll(lines[0], " ", ""), false)[0]
	p2dist := Getints(strings.ReplaceAll(lines[1], " ", ""), false)[0]

	Sol(enum(p2time, p2dist))
}

func enum(time, dist int) int {
	r := 0
	for tbtn := 1; tbtn < time; tbtn++ {
		tm := dist / tbtn
		if tbtn+tm < time {
			r++
		}
	}
	return r
}
