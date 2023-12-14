package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

type state struct {
	cycle int
	score int
}

var C = map[string]*state{}

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}

	rollVert(0, +1)
	Sol(score(), 136, 106997)

	offset := 0
	length := 0

	for cycle := 1; cycle <= 10000; cycle++ {
		rollVert(0, +1)
		rollHoriz(0, +1)
		rollVert(len(M)-1, -1)
		rollHoriz(len(M[0])-1, -1)

		s := printmatrix()
		if C[s] != nil {
			offset = C[s].cycle
			length = cycle - C[s].cycle
			break
		}
		C[s] = &state{cycle, score()}
	}

	predicted := byid((1000000000-offset)%length + offset)
	Sol(predicted.score, 64, 99641)
}

func byid(cycle int) *state {
	for _, s := range C {
		if s.cycle == cycle {
			return s
		}
	}
	panic("blah")
}

func printmatrix() string {
	s := ""
	for i := range M {
		s += string(M[i]) + "\n"
	}
	return s[:len(s)-1]
}

func rollVert(starti, delta int) {
	for i := starti; i >= 0 && i < len(M); i += delta {
		for j := range M[i] {
			if M[i][j] == 'O' {
				moveVert(i, j, -delta)
			}
		}
	}
}

func rollHoriz(startj, delta int) {
	for j := startj; j >= 0 && j < len(M[0]); j += delta {
		for i := range M {
			if M[i][j] == 'O' {
				moveHoriz(i, j, -delta)
			}
		}
	}
}

func moveVert(i, j, delta int) {
	for i+delta >= 0 && i+delta < len(M) {
		if M[i+delta][j] != '.' {
			break
		}
		M[i][j] = '.'
		M[i+delta][j] = 'O'
		i += delta
	}
}

func moveHoriz(i, j, delta int) {
	for j+delta >= 0 && j+delta < len(M[0]) {
		if M[i][j+delta] != '.' {
			break
		}
		M[i][j] = '.'
		M[i][j+delta] = 'O'
		j += delta
	}
}

func score() int {
	r := 0
	for i := range M {
		s := len(M) - i
		for j := range M[i] {
			if M[i][j] == 'O' {
				r += s
			}
		}
	}
	return r
}
