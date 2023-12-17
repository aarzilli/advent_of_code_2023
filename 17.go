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

var M [][]int

const (
	west = iota
	south
	east
	north
)

type point struct{i, j int}
type state struct {
	point
	dir int
	count int
}

const ispart1 = false

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]int, len(lines))
	for i := range lines {
		M[i] = Vatoi(Spac(lines[i], "", -1))
	}
	
	djk := NewDijkstra[state](state{point:point{-1,-1}})
	var cur state
	djk.PopTo(&cur)
	djk.Add(cur, state{point{0,1}, west, 1}, M[0][1])
	djk.Add(cur, state{point{1,0}, south, 1}, M[1][0])
	
	min := -1
	
	for djk.PopTo(&cur) {
		if cur.i == len(M)-1 && cur.j == len(M[cur.i])-1 {
			//pln("path found", djk.Dist[cur])
			//debugsol(djk, cur)
			if min == -1 || djk.Dist[cur] < min {
				min = djk.Dist[cur]
			}
		}
		
		advance := func(dir int, count int) {
			next := cur
			next.dir = dir
			next.count = count
			
			var di, dj int
			switch dir {
			case west:
				dj = +1
			case south:
				di = +1
			case east:
				dj = -1
			case north:
				di = -1
			default:
				panic("blah")
			}
			
			next.point.i += di
			next.point.j += dj
			if next.point.i < 0 || next.point.i >= len(M) || next.point.j < 0 || next.point.j >= len(M[next.point.i]) {
				return
			}
			djk.Add(cur, next, M[next.point.i][next.point.j])
		}
		
		if ispart1 {
			if cur.count < 3 {
				advance(cur.dir, cur.count+1)
			}
			dira := normdir(cur.dir-1)
			dirb := normdir(cur.dir+1)
			advance(dira, 1)
			advance(dirb, 1)
		} else {
			if cur.count < 4 {
				advance(cur.dir, cur.count+1)
			} else if cur.count < 10 {
				advance(cur.dir, cur.count+1)
				advance(normdir(cur.dir-1), 1)
				advance(normdir(cur.dir+1), 1)
			} else {
				advance(normdir(cur.dir-1), 1)
				advance(normdir(cur.dir+1), 1)
			}
		}
	}
	
	Sol(min)
}

func normdir(dir int) int {
	switch dir  {
	case -1:
		return north
	case 4:
		return west
	}
	return dir
}

func debugsol(djk *Dijkstra[state], cur state) {
	r := djk.PathTo(cur)
	M2 := make([][]byte, len(M))
	for i := range M {
		M2[i] = make([]byte, len(M[i]))
		for j := range M[i] {
			M2[i][j] = byte(M[i][j]) + '0'
		}
	}
	for _, s := range r {
		//pln(s)
		if s.point.i == -1 && s.point.j == -1 {
			continue
		}
		M2[s.point.i][s.point.j] = dir2str(s.dir)
	}
	
	for i := range M2 {
		pln(string(M2[i]))
	}
}

func dir2str(dir int) byte {
	switch dir {
	case west:
		return '>'
	case south:
		return 'v'
	case north:
		return '^'
	case east:
		return '<'
	default:
		panic("blah")
	}
}