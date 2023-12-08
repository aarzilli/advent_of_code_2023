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

type node struct {
	name  string
	child []string
}

var M = map[string]*node{}

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	pf("len %d\n", len(groups))
	instr := groups[0]
	_ = instr
	lines := Spac(groups[1], "\n", -1)
	for _, line := range lines {
		v := Spac(line, "=", -1)
		src := v[0]
		v = Spac(v[1][1:len(v[1])-1], ",", -1)
		M[src] = &node{src, v}
	}

	cur := M["AAA"]
	step := 0
	i := 0
	for {
		m := instr[i%len(instr)]
		if cur.name == "ZZZ" {
			break
		}
		if m == 'L' {
			cur = M[cur.child[0]]
		} else {
			cur = M[cur.child[1]]
		}
		step++
		i++
	}

	Sol(step)

	curset := make(Set[string])
	for i := range M {
		if M[i].name[len(M[i].name)-1] == 'A' {
			curset[M[i].name] = true
		}
	}

	first := map[string]int{}

	for name := range curset {
		cur := M[name]
		step := 0
		i := 0
		for {
			m := instr[i%len(instr)]
			if cur.name[len(cur.name)-1] == 'Z' {
				if first[name] == 0 {
					first[name] = step
					break
				}
			}
			if m == 'L' {
				cur = M[cur.child[0]]
			} else {
				cur = M[cur.child[1]]
			}
			step++
			i++
		}
	}

	v := []int{}
	for _, x := range first {
		v = append(v, x)
	}

	Sol(LCM(v[0], v[1], v[2:]...))
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
