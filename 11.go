package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

type point struct{ i, j int }

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		M = append(M, []byte(line))
	}

	for _, delta := range []int{2, 1000000} {
		points := []point{}

		for i := range M {
			for j := range M[i] {
				if M[i][j] == '#' {
					points = append(points, point{i, j})
				}
			}
		}

		for i := len(M) - 1; i >= 0; i-- {
			if emptyline(M[i]) {
				addline(points, i, delta-1)
			}
		}

		for j := len(M[0]) - 1; j >= 0; j-- {
			if emptycol(j) {
				addcol(points, j, delta-1)
			}
		}

		r := 0
		for i := range points {
			for j := i + 1; j < len(points); j++ {
				r += dist(points[i], points[j])
			}
		}
		Sol(r)
	}
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
