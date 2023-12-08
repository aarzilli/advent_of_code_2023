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
	name string
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
		M[src] = &node{ src, v }
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
}
