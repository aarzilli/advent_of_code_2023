package main

import (
	. "aoc/util"
	"os"
)

var M [][]int

const (
	west = iota
	south
	east
	north
)

type point struct{ i, j int }
type state struct {
	point
	dir   int
	count int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]int, len(lines))
	for i := range lines {
		M[i] = Vatoi(Spac(lines[i], "", -1))
	}
	Sol(solve(0, 3))
	Sol(solve(4, 10))
}

func solve(canturn, mustturn int) int {
	djk := NewDijkstra[state](state{point: point{-1, -1}})
	var cur state
	djk.PopTo(&cur)
	djk.Add(cur, state{point{0, 1}, west, 1}, M[0][1])
	djk.Add(cur, state{point{1, 0}, south, 1}, M[1][0])

	min := -1

	for djk.PopTo(&cur) {
		if cur.i == len(M)-1 && cur.j == len(M[cur.i])-1 {
			if cur.count >= canturn {
				if min == -1 || djk.Dist[cur] < min {
					min = djk.Dist[cur]
				}
			}
		}

		if min >= 0 && djk.Dist[cur] > min {
			continue
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

		if cur.count < mustturn {
			advance(cur.dir, cur.count+1)
		}
		if cur.count >= canturn {
			advance(normdir(cur.dir-1), 1)
			advance(normdir(cur.dir+1), 1)
		}
	}

	return min
}

func normdir(dir int) int {
	switch dir {
	case -1:
		return north
	case 4:
		return west
	}
	return dir
}
