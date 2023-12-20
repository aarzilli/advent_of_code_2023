package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

var kind = map[string]byte{}
var Adj = map[string][]string{}
var queue = []pulse{}
var flipflop = map[string]bool{}
var conj = map[string]map[string]bool{}
var highs, lows int
var seen = map[string]*result{}
var cycle int

type result struct {
	highs, lows int
	nextsig     string
}

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

	/*pf("digraph G {\n")
	for cur := range Adj {
		lbl := cur
		switch kind[cur]  {
		case '%', '&':
			lbl = string(kind[cur])+cur
		}
		pf("%s [label=%q];\n", cur, lbl)
		for _, next := range Adj[cur] {
			pf("%s -> %s;\n", cur, next)
		}
	}
	pf("}\n")*/

	for cnt := 0; cnt < 100000; cnt++ {
		//printthing()
		fmt.Println("BUTTON PUSH", cycle+1)
		send(false, "button", broad)
		cycle++
		for len(queue) > 0 {
			process()
		}
	}

	/*
		const N = 1000000
		sig := signature()
		var cnt int
		var preloop result
		loopfound := false
		for cnt = 0; cnt < N; cnt++ {
			pln("signature:", sig)
			if seen[sig] != nil {
				loopfound = true
				break
			}
			highs, lows = 0, 0
			send(false, "button", broad)
			for len(queue) > 0 {
				process()
			}
			nextsig := signature()
			preloop.highs += highs
			preloop.lows += lows
			seen[sig] = &result{ highs, lows, nextsig }
			sig = nextsig
		}

		if !loopfound {
			Sol(preloop.highs*preloop.lows)
		} else {
			var looplen int
			var loopresult result

			startsig := sig
			for {
				looplen++
				loopresult.highs += seen[sig].highs
				loopresult.lows += seen[sig].lows
				sig = seen[sig].nextsig
				if sig == startsig {
					break
				}
			}

			pln("remaining", N-cnt, "looplen", looplen, "each loop", loopresult.highs, loopresult.lows)

			if (N-cnt) % looplen != 0 {
				panic("uneven")
			}

			mul := (N-cnt) / looplen
			tothighs := preloop.highs + loopresult.highs * mul
			totlows := preloop.lows + loopresult.lows * mul
			pln(tothighs, totlows)

			Sol(tothighs*totlows)
		}*/
}

func printthing() {
	r := []byte{}
	for _, k := range []string{"dc", "nx", "qr", "pz", "rv", "mp", "cj", "pg", "df", "rs", "gq", "vx"} {
		if flipflop[k] {
			r = append(r, '1')
		} else {
			r = append(r, '0')
		}
	}
	pln(string(r))
}

func conjand(name string) bool {
	r := true
	for _, v := range conj[name] {
		if !v {
			r = false
			break
		}
	}
	return r
}

const debugprocess = false

func process() {
	p := queue[0]
	queue = queue[1:]
	switch {
	case p.dst == broad:
		for _, x := range Adj[p.dst] {
			if debugprocess {
				pln(p.dst, p.lvl, "->", x)
			}
			send(p.lvl, p.dst, x)
		}
	case kind[p.dst] == '%':
		if !p.lvl {
			flipflop[p.dst] = !flipflop[p.dst]
			for _, x := range Adj[p.dst] {
				if debugprocess {
					pln(p.dst, flipflop[p.dst], "->", x)
				}
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
		if p.dst == os.Args[2] {
			pln(p.dst, "send", !r, cycle)
			if r {
				cycle = 0
			}
		}
		for _, x := range Adj[p.dst] {
			if debugprocess {
				pln(p.dst, !r, "->", x)
			}
			send(!r, p.dst, x)
		}

	default:
		//do nothing
	}
}

func send(lvl bool, src, dst string) {
	if lvl {
		highs++
	} else {
		lows++
	}
	if len(queue) == 0 {
		queue = []pulse{{lvl, src, dst}}
		return
	}
	queue = append(queue, pulse{lvl, src, dst})
}

func signature() string {
	r := []byte{}
	ffks := Keys(flipflop)
	sort.Strings(ffks)
	for _, ffk := range ffks {
		if flipflop[ffk] {
			r = append(r, '1')
		} else {
			r = append(r, '0')
		}
	}

	cks := Keys(conj)
	sort.Strings(cks)
	for _, ck := range cks {
		r = append(r, []byte(ck)...)
		r = append(r, '{')
		pks := Keys(conj[ck])
		sort.Strings(pks)
		for _, pk := range pks {
			if conj[ck][pk] {
				r = append(r, '1')
			} else {
				r = append(r, '0')
			}
		}
		r = append(r, '}')
	}
	return string(r)
}
