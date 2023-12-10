package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var M [][]byte

type point struct{ i, j int }

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	var start point
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'S' {
				start = point{i, j}
				break
			}
		}
	}
	switch start {
	case point{1, 1}:
		switch len(M) {
		case 5:
			M[start.i][start.j] = 'F'
		case 9:
			M[start.i][start.j] = 'F'
		default:
			panic(fmt.Errorf("unknown input %d", len(M)))
		}
	case point{2, 0}:
		M[start.i][start.j] = 'F'
	case point{61, 10}:
		M[start.i][start.j] = 'J'
	case point{4, 12}:
		M[start.i][start.j] = 'F'

	case point{0, 4}:
		M[start.i][start.j] = '7'
	default:
		panic("unknown input")
	}

	dist := make(map[point]int)
	dist[start] = 0
	fringe := []point{start}
	for len(fringe) > 0 {
		cur := fringe[0]
		fringe = fringe[1:]

		add := func(p point) {
			if _, ok := dist[p]; !ok {
				fringe = append(fringe, p)
				dist[p] = dist[cur] + 1
			}
		}

		switch M[cur.i][cur.j] {
		case '|':
			add(point{cur.i - 1, cur.j})
			add(point{cur.i + 1, cur.j})
		case '-':
			add(point{cur.i, cur.j + 1})
			add(point{cur.i, cur.j - 1})
		case 'L':
			add(point{cur.i - 1, cur.j})
			add(point{cur.i, cur.j + 1})
		case 'F':
			add(point{cur.i + 1, cur.j})
			add(point{cur.i, cur.j + 1})
		case 'J':
			add(point{cur.i - 1, cur.j})
			add(point{cur.i, cur.j - 1})
		case '7':
			add(point{cur.i + 1, cur.j})
			add(point{cur.i, cur.j - 1})
		case '.':
			panic("bad")
		}
	}
	maxv := 0
	for _, v := range dist {
		if v > maxv {
			maxv = v
		}
	}
	Sol(maxv)

	E := make([][]byte, len(M)*3)
	for i := range E {
		E[i] = make([]byte, len(M[0])*3)
	}

	for i := range M {
		for j := range M[i] {

			copy := func(s ...string) {
				for ei := 0; ei < 3; ei++ {
					for ej := 0; ej < 3; ej++ {
						E[i*3+ei][j*3+ej] = s[ei][ej]
					}
				}
			}

			if _, ok := dist[point{i, j}]; ok {
				switch M[i][j] {
				case '|':
					copy(
						".X.",
						".X.",
						".X.")
				case '-':
					copy(
						"...",
						"XXX",
						"...")

				case 'F':
					copy(
						"...",
						".XX",
						".X.")
				case '7':
					copy(
						"...",
						"XX.",
						".X.")
				case 'L':
					copy(
						".X.",
						".XX",
						"...")

				case 'J':
					copy(
						".X.",
						"XX.",
						"...")

				default:
					panic("blah")
				}
			} else {
				copy(
					"...",
					"...",
					"...")

			}
		}
	}

	for i := range E {
		for j := range E[i] {
			mark(E, point{i, j})
		}
	}

	/*
		for i := range E {
			fmt.Println(string(E[i]))
		}
	*/

	inside := 0

	for i := range M {
		for j := range M[i] {
			if _, ok := dist[point{i, j}]; ok {
				continue
			}

			switch check(E, point{i * 3, j * 3}) {
			case 'O':
				M[i][j] = 'O'
			case 'I':
				M[i][j] = 'I'
				inside++
			default:
				panic("blah")
			}
		}
	}

	/*
		pln()
		for i := range M {
			pln(string(M[i]))
		}
	*/

	Sol(inside)
}

func check(E [][]byte, start point) byte {
	ch := byte(0)
	for ei := 0; ei < 3; ei++ {
		for ej := 0; ej < 3; ej++ {
			if ch == 0 {
				ch = E[start.i+ei][start.j+ej]
			} else {
				if ch != E[start.i+ei][start.j+ej] {
					panic("blah")
				}
			}
		}
	}
	return ch
}

func mark(E [][]byte, start point) {
	if E[start.i][start.j] != '.' {
		return
	}

	fringe := make(Set[point])
	fringe[start] = true
	reachable := make(Set[point])
	outside := false
	for len(fringe) > 0 {
		cur := OneKey(fringe)
		delete(fringe, cur)
		reachable[cur] = true

		add := func(p point) {
			if reachable[p] {
				return
			}
			if p.i < 0 || p.i >= len(E) || p.j < 0 || p.j >= len(E[p.i]) {
				outside = true
				return
			}
			if E[p.i][p.j] != '.' {
				return
			}
			fringe[p] = true
		}

		add(point{cur.i, cur.j - 1})
		add(point{cur.i, cur.j + 1})
		add(point{cur.i - 1, cur.j})
		add(point{cur.i + 1, cur.j})
	}

	if outside {
		for p := range reachable {
			E[p.i][p.j] = 'O'
		}
	} else {
		for p := range reachable {
			E[p.i][p.j] = 'I'
		}
	}
}
