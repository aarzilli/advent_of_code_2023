package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
	"strings"
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
	
	
	/*
	cur := M["AAA"]
	step := 0
	i := 0
	for {
		pln(cur.name, step)
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
	
	Sol(step)*/
	
	curset := make(Set[string])
	for i := range M {
		if M[i].name[len(M[i].name)-1] == 'A' {
			curset[M[i].name] = true
		}
	}
	
	/*seen := map[string]bool{}
	
	step := 0
	i := 0
	for {
		signature := printset(curset)
		pln(signature, len(seen))
		if seen[signature] {
			pln("repeated")
			os.Exit(1)
		}
		seen[signature] = true
		m := instr[i%len(instr)]
		ok := true
		for name := range curset {
			if name[len(name)-1] != 'Z' {
				ok = false
				break
			}
		}
		if ok {
			break
		}
		
		newset := make(Set[string])
		for name := range curset {
			var next string
			if m == 'L' {
				next = M[name].child[0]
			} else {
				next = M[name].child[1]
			}
			newset[next] = true
		}
		curset = newset
		
		step++
		i++
	}
	
	Sol(step)*/
	
	/*
	for name := range curset {
		cur := M[name]
		step := 0
		i := 0
		for {
			pln(step, cur.name)
			m := instr[i%len(instr)]
			if name[len(name)-1] == 'Z' {
				pln("at", name, "after", step)
			}
			if m == 'L' {
				cur = M[cur.child[0]]
			} else {
				cur = M[cur.child[1]]
			}
			step++
			i++
		}
	}*/
	
	first := map[string]int{}
	period := map[string]int{}
	
	for name := range curset {
		cur := M[name]
		step := 0
		i := 0
		for {
			m := instr[i%len(instr)]
			if cur.name[len(cur.name)-1] == 'Z' {
				if first[name] == 0 {
					first[name] = step
				} else if period[name] == 0 {
					period[name] = step - first[name]
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
	
	pln(first)
	pln(period)
	
}

func printset(curset Set[string]) string {
	x := []string{}
	for name := range curset {
		x = append(x, name)
	}
	sort.Strings(x)
	return strings.Join(x, " ")
}