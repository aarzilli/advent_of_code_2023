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

type point struct{ i, j int }
type instr struct {
	dir   byte
	count int
	color string
}

var M = map[point]bool{}

type intvl struct {
	i          int
	start, end int
	filled     bool
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	instrs := []instr{}
	for _, line := range lines {
		v := Spac(line, " ", -1)
		instrs = append(instrs, instr{
			dir:   v[0][0],
			count: Atoi(v[1]),
			color: v[2][2 : len(v[2])-1],
		})
	}

	cur := point{0, 0}
	M[cur] = true
	for _, instr := range instrs {
		var di, dj int
		switch instr.dir {
		case 'R':
			dj = 1
		case 'L':
			dj = -1
		case 'U':
			di = -1
		case 'D':
			di = +1
		}

		for k := 0; k < instr.count; k++ {
			cur.i += di
			cur.j += dj
			M[cur] = true
		}
	}

	minj := map[int]int{}
	maxj := map[int]int{}
	mini, maxi := 0, 0

	for p := range M {
		if !M[p] {
			continue
		}
		if p.i < mini {
			mini = p.i
		}
		if p.i > maxi {
			maxi = p.i
		}
		if min, ok := minj[p.i]; !ok || p.j < min {
			minj[p.i] = p.j
		}
		if max, ok := maxj[p.i]; !ok || p.j > max {
			maxj[p.i] = p.j
		}
	}

	Minj, Maxj := 0, 0
	for i, _ := range minj {
		if minj[i] < Minj {
			Minj = minj[i]
		}
		if maxj[i] > Maxj {
			Maxj = maxj[i]
		}
	}

	intvls := map[int][]intvl{}

	for i := mini; i <= maxi; i++ {
		js := []int{}
		for j := Minj; j <= Maxj; j++ {
			if M[point{i, j}] {
				js = append(js, j)
			}
		}
		for len(js) > 0 {
			start := findconsecutive(js)
			js = js[len(start):]
			intvls[i] = append(intvls[i], intvl{i, start[0], start[len(start)-1] + 1, true})
			if len(js) > 0 {
				intvls[i] = append(intvls[i], intvl{i, start[len(start)-1] + 1, js[0], false})
			}
		}
		pln(i, ":", intvls[i])
	}

	interior0 := intvls[mini][0]
	interior0.start++
	interior0.end--
	pln(interior0)
	tofill := findintersectingempty(intvls, mini+1, interior0)
	part1 := len(M)
	for len(tofill) > 0 {
		cur := tofill[len(tofill)-1]
		tofill = tofill[:len(tofill)-1]
		if cur.filled {
			continue
		}
		pln("filling", cur)
		cur.filled = true
		for j := cur.start; j < cur.end; j++ {
			M[point{cur.i, j}] = true
		}
		part1 += cur.end - cur.start
		tofill = append(tofill, findintersectingempty(intvls, cur.i+1, *cur)...)
		tofill = append(tofill, findintersectingempty(intvls, cur.i-1, *cur)...)
	}

	Sol(part1)

	/*
		for i := mini; i <= maxi; i++ {
			for j := Minj; j <= Maxj; j++ {
				if M[point{i, j}] {
					pf("#")
				} else {
					pf(".")
				}
			}
			pln()
		}*/
}

func findconsecutive(v []int) []int {
	for i := 1; i < len(v); i++ {
		if v[i] != v[i-1]+1 {
			return v[:i]
		}
	}
	return v
}

func findintersectingempty(intvls map[int][]intvl, i int, prev intvl) []*intvl {
	r := []*intvl{}
	for k := range intvls[i] {
		intvl := &intvls[i][k]
		if !intvl.filled && (intvl.contains(prev.start) || intvl.contains(prev.end-1) || prev.contains(intvl.start) || prev.contains(intvl.end-1)) {
			r = append(r, intvl)
		}
	}
	return r
}

func (a *intvl) contains(p int) bool {
	return a.start <= p && p < a.end
}

/*func floodfill(start point) {
	fringe := make(Set[point])
	fringe[start] = true
	for len(fringe) > 0 {
		cur := OneKey(fringe)
		delete(fringe, cur)
		if M[cur] {
			continue
		}
		M[cur] = true
		fringe[point{cur.i-1, cur.j}] = true
		fringe[point{cur.i+1, cur.j}] = true
		fringe[point{cur.i, cur.j-1}] = true
		fringe[point{cur.i, cur.j+1}] = true
	}
}*/
