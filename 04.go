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
	copies := make([]int, len(lines))
	for i := range copies {
		copies[i] = 1
	}
	for i, line := range lines {
		v := Spac(line, ":", -1)
		v = Spac(v[1], "|", -1)
		winning := Vatoi(Noempty(Spac(v[0], " ", -1)))
		having := Vatoi(Noempty(Spac(v[1], " ", -1)))

		ws := make(Set[int])
		for _, w := range winning {
			ws[w] = true
		}

		pts := 0
		count := 0
		for _, h := range having {
			if ws[h] {
				if pts == 0 {
					pts = 1
				} else {
					pts += pts
				}
				count++
			}
		}
		part1 += pts
		for j := 0; j < count; j++ {
			copies[i+j+1] += copies[i]
		}
	}
	Sol(part1)
	Sol(Sum(copies))

}
