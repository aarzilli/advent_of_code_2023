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

var Adj = map[point][]edge{}

type edge struct {
	steps int
	next  point
}

const part2 = false

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}

	if part2 {
		for i := range M {
			for j := range M[i] {
				if M[i][j] != '#' {
					M[i][j] = '.'
				}
			}
		}
	}

	start := point{0, 1}
	seen := make(Set[point])
	seen[start] = true
	edg, nb, _ := walk(start, Set[point]{start: true})
	Adj[start] = append(Adj[start], edg)
	for _, next := range nb {
		makegraph(edg.next, next, seen)
	}

	for cur := range Adj {
		for _, e := range Adj[cur] {
			pln(cur, e)
		}
		pln()
	}

	enum(start, 0, make(Set[point]), []point{start})
	Sol(maxdist)
}

var maxdist int

func enum(cur point, dist int, seen Set[point], path []point) {
	seen[cur] = true
	defer delete(seen, cur)

	if isend(cur) {
		pln(dist, path)
		if dist > maxdist {
			maxdist = dist
		}
	}

	for _, e := range Adj[cur] {
		if !seen[e.next] {
			enum(e.next, dist+e.steps, seen, append(path, e.next))
		}
	}
}

func isend(p point) bool {
	return p.i == len(M)-1 && p.j == len(M[p.i])-2
}

func walk(start point, seen Set[point]) (r edge, nb []point, ok bool) {
	p := start
	for {
		seen[p] = true
		if isend(p) {
			r.next = p
			return r, []point{}, true
		}
		v := neighbors(p)
		v = Filter(func(p2 point) bool { return !seen[p2] }, v)
		if len(v) > 1 {
			r.next = p
			return r, v, true
		}
		if len(v) == 0 {
			return edge{steps: -1, next: point{-1, -1}}, []point{}, false
		}
		p = v[0]
		r.steps++
	}
}

func makegraph(cur, next point, seen Set[point]) {
	seen[cur] = true
	edg, nb, ok := walk(next, Set[point]{cur: true, next: true})
	if !ok {
		return
	}
	edg.steps++
	Adj[cur] = append(Adj[cur], edg)
	if !seen[edg.next] {
		for _, next := range nb {
			makegraph(edg.next, next, seen)
		}
	}
}

func neighbors(p point) []point {
	r := []point{}

	switch M[p.i][p.j] {
	case '>':
		return []point{{p.i, p.j + 1}}
	case 'v':
		return []point{{p.i + 1, p.j}}
	case '^':
		return []point{{p.i - 1, p.j}}
	case '<':
		return []point{{p.i, p.j - 1}}
	}

	for _, dp := range []point{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}} {
		p2 := p
		p2.i += dp.i
		p2.j += dp.j

		if valid(p2) /*&& canmove(p, p2)*/ {
			r = append(r, p2)
		}
	}
	return r
}

func valid(p point) bool {
	if p.i < 0 || p.i >= len(M) || p.j < 0 || p.j >= len(M[p.i]) {
		return false
	}
	if M[p.i][p.j] == '#' {
		return false
	}
	return true
}

func canmove(p1, p2 point) bool {
	switch M[p2.i][p2.j] {
	case '>':
		if p1.i == p2.i && p1.j == p2.j+1 {
			return false
		}
	case '<':
		if p1.i == p2.i && p1.j == p2.j-1 {
			return false
		}
	case '^':
		if p1.i == p2.i-1 && p1.j == p2.j {
			return false
		}
	case 'v':
		if p1.i == p2.i+1 && p1.j == p2.j {
			return false
		}
	}
	return true
}

// 5858
