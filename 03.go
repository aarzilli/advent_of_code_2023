package main

import (
	. "aoc/util"
	"os"
)

var M []string

func main() {
	lines := Input(os.Args[1], "\n", true)
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

	part1 := 0
	part2 := 0

	symboladj := make(Set[number])
	for i := range M {
		for j := range M[i] {
			if !issymbol(M[i][j]) {
				continue
			}
			gearadj := make(Set[number])
			for i0 := -1; i0 <= 1; i0++ {
				for j0 := -1; j0 <= 1; j0++ {
					if i0 == 0 && j0 == 0 {
						continue
					}
					ch := get(i+i0, j+j0)
					if isdigit(ch) {
						num := expand([2]int{i + i0, j + j0})
						gearadj.Add(num)
					}
				}
			}
			if M[i][j] == '*' && len(gearadj) == 2 {
				mul := 1
				for num := range gearadj {
					mul *= num.n
				}
				part2 += mul
			}
			symboladj.AddSet(gearadj)
		}
	}

	for num := range symboladj {
		part1 += num.n
	}

	Sol(part1)
	Sol(part2)
}

func isdigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func issymbol(ch byte) bool {
	return !isdigit(ch) && ch != '.' && ch != 0
}

type number struct {
	p [2]int
	n int
}

func expand(p [2]int) number {
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
	return number{[2]int{p[0], j0}, Atoi(M[p[0]][j0:j1])}
}
