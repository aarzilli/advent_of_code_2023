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

var M [][]byte

type point struct{ i, j int }

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		M = append(M, []byte(line))
	}

	/*M2 := [][]byte{}

	for i := range M {
		if emptyline(M[i]) {
			M2 = append(M2, M[i], M[i])
		} else {
			M2 = append(M2, M[i])
		}
	}

	M = M2
	M2 = make([][]byte, len(M))

	for j := range M[0] {
		if emptycol(j) {
			appendcol(M2, j)
			appendcol(M2, j)
		} else {
			appendcol(M2, j)
		}
	}

	M = M2

	for i := range M {
		pln(string(M[i]))
	}*/

	points := []point{}

	for i := range M {
		for j := range M[i] {
			if M[i][j] == '#' {
				points = append(points, point{i, j})
			}
		}
	}

	N := 1000000 - 1
	//N := 100-1

	for i := len(M) - 1; i >= 0; i-- {
		if emptyline(M[i]) {
			addline(points, i, N)
		}
	}

	for j := len(M[0]) - 1; j >= 0; j-- {
		if emptycol(j) {
			addcol(points, j, N)
		}
	}

	pln(points)

	part1 := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			part1 += dist(points[i], points[j])
		}
	}
	Sol(part1)
}

func emptyline(line []byte) bool {
	for i := range line {
		if line[i] != '.' {
			return false
		}
	}
	return true
}

func emptycol(j int) bool {
	for i := range M {
		if M[i][j] != '.' {
			return false
		}
	}
	return true
}

func appendcol(M2 [][]byte, j int) {
	for i := range M2 {
		M2[i] = append(M2[i], M[i][j])
	}
}

func addline(points []point, starti int, d int) {
	for k := range points {
		p := &points[k]
		if p.i > starti {
			p.i += d
		}
	}
}

func addcol(points []point, startj int, d int) {
	for k := range points {
		p := &points[k]
		if p.j > startj {
			p.j += d
		}
	}
}

func dist(p1, p2 point) int {
	return Abs(p1.i-p2.i) + Abs(p1.j-p2.j)
}

// [{0 4} {1 9} {2 0} {5 8} {6 1} {7 12} {10 9} {11 0} {11 5}]
