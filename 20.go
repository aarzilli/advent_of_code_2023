package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

var kind = map[string]byte{}
var Adj = map[string][]string{}
var queue = []pulse{}
var flipflop = map[string]bool{}
var conj = map[string]map[string]bool{}
var pulsecnt = map[bool]int{}
var cycle int
var cyclesearch string
var cyclefound bool

type pulse struct {
	lvl      bool
	src, dst string
}

const broad = "broadcaster"

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		v := Spac(line, "->", -1)
		name := v[0]
		switch ch := name[0]; ch {
		case '%', '&':
			name = name[1:]
			kind[name] = ch
		}
		for _, x := range Spac(v[1], ",", -1) {
			Adj[name] = append(Adj[name], x)
		}
	}

	if len(os.Args) > 2 && os.Args[2] == "graph" {
		fmt.Printf("digraph G {\n")
		for cur := range Adj {
			lbl := cur
			switch kind[cur] {
			case '%', '&':
				lbl = string(kind[cur]) + cur
			}
			fmt.Printf("%s [label=%q];\n", cur, lbl)
			for _, next := range Adj[cur] {
				fmt.Printf("%s -> %s;\n", cur, next)
			}
		}
		fmt.Printf("}\n")
		return
	}

	const N = 1000

	reset()
	var cnt int
	for cnt = 0; cnt < N; cnt++ {
		send(false, "button", broad)
		for len(queue) > 0 {
			process()
		}
	}

	Sol(pulsecnt[true] * pulsecnt[false])

	preds := predecessors(predecessors("rx")[0])
	cyclelen := []int{}

	for _, pred := range preds {
		cyclesearch = pred
		cyclefound = false
		cycle = 0
		reset()
	cyclesearchloop:
		for cnt := 0; cnt < 100000; cnt++ {
			send(false, "button", broad)
			cycle++
			for len(queue) > 0 {
				process()
				if cyclefound {
					cyclelen = append(cyclelen, cycle)
					break cyclesearchloop
				}
			}
		}
	}

	Sol(LCM(cyclelen[0], cyclelen[1], cyclelen[2:]...))
}

func reset() {
	for cur := range Adj {
		if kind[cur] == '%' {
			flipflop[cur] = false
		}
		for _, next := range Adj[cur] {
			if kind[next] == '&' {
				if conj[next] == nil {
					conj[next] = map[string]bool{}
				}
				conj[next][cur] = false
			}
		}
	}
}

func process() {
	p := queue[0]
	queue = queue[1:]
	switch {
	case p.dst == broad:
		for _, x := range Adj[p.dst] {
			send(p.lvl, p.dst, x)
		}
	case kind[p.dst] == '%':
		if !p.lvl {
			flipflop[p.dst] = !flipflop[p.dst]
			for _, x := range Adj[p.dst] {
				send(flipflop[p.dst], p.dst, x)
			}
		}
	case kind[p.dst] == '&':
		conj[p.dst][p.src] = p.lvl
		r := true
		for _, v := range conj[p.dst] {
			if !v {
				r = false
				break
			}
		}
		if cyclesearch == p.dst {
			if !r {
				cyclefound = true
			}
		}
		for _, x := range Adj[p.dst] {
			send(!r, p.dst, x)
		}

	default:
		//do nothing
	}
}

func send(lvl bool, src, dst string) {
	pulsecnt[lvl]++
	if len(queue) == 0 {
		queue = []pulse{{lvl, src, dst}}
		return
	}
	queue = append(queue, pulse{lvl, src, dst})
}

func predecessors(name string) []string {
	r := []string{}
	for cur := range Adj {
		for _, next := range Adj[cur] {
			if next == name {
				r = append(r, cur)
			}
		}
	}
	return r
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
