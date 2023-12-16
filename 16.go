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

const (
	west = iota
	south
	east
	north
)

type point struct{ i, j int }
type state struct {
	point
	dir int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	Sol(part1(state{point{0, 0}, east}))

	var max int = -1
	domax := func(n int) {
		if n > max {
			max = n
		}
	}
	for i := range M {
		domax(part1(state{point{i, 0}, east}))
		domax(part1(state{point{i, len(M[i]) - 1}, west}))
	}
	for j := range M[0] {
		domax(part1(state{point{0, j}, south}))
		domax(part1(state{point{len(M) - 1, j}, north}))
	}
	Sol(max)
}

func part1(start state) int {
	seen := make(Set[state])
	fringe := make(Set[state])
	fringe[start] = true

	for len(fringe) > 0 {
		cur := OneKey(fringe)
		delete(fringe, cur)

		if cur.i < 0 || cur.i >= len(M) || cur.j < 0 || cur.j >= len(M[cur.i]) {
			continue
		}

		seen[cur] = true

		//printmatrix(seen)
		//pln()

		add := func(dir int) {
			var di, dj int
			switch dir {
			case north:
				di = -1
				dj = 0
			case south:
				di = +1
				dj = 0
			case west:
				di = 0
				dj = -1
			case east:
				di = 0
				dj = +1
			default:
				panic("blah")
			}

			next := state{point{cur.i + di, cur.j + dj}, dir}
			if seen[next] {
				return
			}
			fringe[next] = true
		}

		switch M[cur.i][cur.j] {
		case '.':
			add(cur.dir)
		case '/':
			switch cur.dir {
			case north:
				add(east)
			case south:
				add(west)
			case west:
				add(south)
			case east:
				add(north)
			default:
				panic("blah")
			}
		case '\\':
			switch cur.dir {
			case north:
				add(west)
			case south:
				add(east)
			case west:
				add(north)
			case east:
				add(south)
			default:
				panic("blah")
			}
		case '|':
			switch cur.dir {
			case north:
				add(north)
			case south:
				add(south)
			case west:
				add(north)
				add(south)
			case east:
				add(north)
				add(south)
			default:
				panic("blah")
			}
		case '-':
			switch cur.dir {
			case north:
				add(east)
				add(west)
			case south:
				add(east)
				add(west)
			case west:
				add(west)
			case east:
				add(east)
			default:
				panic("blah")
			}
		default:
			panic("blah")
		}
	}

	pseen := make(Set[point])
	for s := range seen {
		pseen[s.point] = true
	}
	return len(pseen)
}

func (s *state) String() string {
	return fmt.Sprintf("%v %s", s.point, dir2str(s.dir))
}

func dir2str(dir int) string {
	switch dir {
	case north:
		return "^"
	case south:
		return "v"
	case east:
		return ">"
	case west:
		return "<"
	default:
		panic("Blah")
	}
}

func printmatrix(seen Set[state]) {
	for i := range M {
		for j := range M[i] {
			if M[i][j] != '.' {
				pf("%c", M[i][j])
			} else {
				found := false
				for s := range seen {
					if s.point == (point{i, j}) {
						found = true
						pf(dir2str(s.dir))
						break
					}
				}
				if !found {
					pf(".")
				}
			}
		}
		pln()
	}
}
