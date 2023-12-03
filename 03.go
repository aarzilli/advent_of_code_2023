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

var M []string

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = lines
	get := func(i, j int) byte {
		if i < 0 || i >= len(M) {
			return 0
		}
		if j < 0 || j >= len(M[i]) {
			return 0
		}
		return M[i][j]
	}
	symboladj := [][2]int{}
	for i := range M {
		for j := range M[i] {
			if !isdigit(M[i][j]) {
				continue
			}
			
			for i0 := -1; i0 <= 1; i0++ {
				for j0 := -1; j0 <= 1; j0++ {
					if i0 == 0 && j0 == 0 {
						continue
					}
					ch := get(i+i0, j+j0)
					if issymbol(ch) {
						if len(symboladj) == 0 || symboladj[len(symboladj)-1] != [2]int{i,j} {
							symboladj = append(symboladj, [2]int{i,j})
						}
					}
				}
			}
		}
	}
	//pf("%v\n", symboladj)
	
	seen := make(map[[2]int]bool)
	part1 := 0
	for _, p := range symboladj {
		p0, n := expand(p)
		if seen[p0] {
			continue
		}
		seen[p0] = true
		//pf("%v %v %d\n", p, p0, n)
		part1 += n
	}
	Sol(part1)
	
	part2 := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] != '*' {
				continue
			}
			gearadj := map[[2]int]struct{}{}
			for i0 := -1; i0 <= 1; i0++ {
				for j0 := -1; j0 <= 1; j0++ {
					if i0 == 0 && j0 == 0 {
						continue
					}
					ch := get(i+i0, j+j0)
					if isdigit(ch) {
						p, _ := expand([2]int{i+i0, j+j0})
						gearadj[p] = struct{}{}
					}
				}
			}
			if len(gearadj) == 2 {
				//pf("%d,%d %v ", i, j, gearadj)
				mul := 1
				for p := range gearadj {
					_, n := expand(p)
					mul *= n
				}
				//pf("%d\n", mul)
				part2 += mul
			}
		}
	}
	Sol(part2)
}

func isdigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func issymbol(ch byte) bool {
	return !isdigit(ch) && ch != '.' && ch != 0
}

func expand(p [2]int) ([2]int, int) {
	var j0, j1 int
	for j0 = p[1]; j0 >= 0; j0-- {
		if !isdigit(M[p[0]][j0]) {
			j0++
			break
		}
	}
	if j0 < 0 {
		j0 = 0
	}
	for j1 = p[1]; j1 < len(M[p[0]]); j1++ {
		if !isdigit(M[p[0]][j1]) {
			break
		}
	}
	return [2]int{p[0], j0 }, Atoi(M[p[0]][j0:j1])
}