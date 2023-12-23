package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

type point struct{ i, j int }

var Adj = map[point][]edge{}

type edge struct {
	steps int
	next  point
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	start := point{0, 1}

	makegraph(start, make(Set[point]))
	enum(start, 0, make(Set[point]), []point{start})
	Sol(maxdist)

	Adj = map[point][]edge{}
	maxdist = 0

	for i := range M {
		for j := range M[i] {
			if M[i][j] != '#' {
				M[i][j] = '.'
			}
		}
	}

	makegraph(start, make(Set[point]))
	enum(start, 0, make(Set[point]), []point{start})
	Sol(maxdist)
}

var maxdist int

func enum(cur point, dist int, seen Set[point], path []point) {
	seen[cur] = true
	defer delete(seen, cur)

	if isend(cur) {
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

func makegraph(cur point, seen Set[point]) {
	if seen[cur] {
		return
	}
	seen[cur] = true
	nb := neighbors(cur)
	for _, next := range nb {
		edg, _, ok := walk(next, Set[point]{cur: true})
		if ok {
			edg.steps++
			Adj[cur] = append(Adj[cur], edg)
			makegraph(edg.next, seen)
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

		if valid(p2) {
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
