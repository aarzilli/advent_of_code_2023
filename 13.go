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

func main() {
	mirrors := Input(os.Args[1], "\n\n", true)
	part1 := 0
	old := make([]int, len(mirrors))
	for i, mirrorstr := range mirrors {
		r := findr(Spac(mirrorstr, "\n", -1), -1)
		old[i] = r
		part1 += r
	}
	Sol(part1)

	part2 := 0
	for i, mirrorstr := range mirrors {
		r := nosmudge(Spac(mirrorstr, "\n", -1), old[i])
		pln(r)
		part2 += r
	}
	Sol(part2)
}

func nosmudge(mirrorstrs []string, old int) int {
	mirror := make([][]byte, len(mirrorstrs))
	for i := range mirrorstrs {
		mirror[i] = []byte(mirrorstrs[i])
	}
	for i := range mirror {
		for j := range mirror[i] {
			mirror[i][j] = invert(mirror[i][j])
			if r := findr(tostrs(mirror), old); r >= 0 {
				return r
			}
			mirror[i][j] = invert(mirror[i][j])
		}
	}
	return -1
}

func printmirror(mirror []string) {
	for i := range mirror {
		pln(mirror[i])
	}
	pln()
}

func invert(ch byte) byte {
	switch ch {
	case '.':
		return '#'
	case '#':
		return '.'
	default:
		panic("blah")
	}
}

func tostrs(mirror [][]byte) []string {
	r := make([]string, len(mirror))
	for i := range mirror {
		r[i] = string(mirror[i])
	}
	return r
}

func findr(mirror []string, old int) int {
	for i := range mirror {
		if isrh(mirror[:i], mirror[i:]) {
			if 100*i != old {
				return 100 * i
			}
		}
	}
	m2 := make([]string, len(mirror[0]))
	for j := range mirror[0] {
		for i := range mirror {
			m2[j] += string(mirror[i][j])
		}
	}
	for i := range m2 {
		if isrh(m2[:i], m2[i:]) {
			if i != old {
				return i
			}
		}
	}
	return -1
}

func isrh(a, b []string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	for i := 0; i < Min([]int{len(a), len(b)}); i++ {
		if b[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}

// 27372
// 40089
