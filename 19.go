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

type rule struct {
	cond byte
	op1 string
	op2 int
	dst string
}

type step struct {
	name string
	rules []rule
}

type partset struct {
	min, max map[string]int
}

var program = map[string]*step{}

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	lines := Noempty(Spac(groups[0], "\n", -1))
	for _, line := range lines {
		step := parse(line)
		program[step.name] = step
	}
	
	part1 := 0
	lines = Noempty(Spac(groups[1], "\n", -1))
	for _, line := range lines {
		part := parsePart(line)
		if process(part, "in") {
			part1 += sum(part)
		}
	}
	Sol(part1)
	
	ps := partset{
		min: map[string]int{ "x": 1, "m": 1, "a": 1, "s": 1 },
		max: map[string]int{ "x": 4000, "m": 4000, "a": 4000, "s": 4000 },
	}
	
	Sol(enum("in", ps))
}

func enum(cur string, ps partset) int {
	switch cur {
	case "A":
		//pln(ps, ps.size())
		return ps.size()
	case "R":
		return 0
	}
	
	tot := 0
	for _, rl := range program[cur].rules {
		passps, failps := ps.split(rl)
		if passps.size() + failps.size() != ps.size() {
			pln(ps)
			pln(passps)
			pln(failps)
			panic("blah")
		}
		tot += enum(rl.dst, passps)
		ps = failps
	}
	return tot
}

func (ps partset) size() int {
	r := 1
	for _, p := range []string{ "x", "m", "a", "s" } {
		r *= ps.max[p] - ps.min[p] + 1
	}
	return r
}

func (ps partset) split(cur rule) (passps, failps partset) {
	passps = ps.copy()
	failps = ps.copy()
	
	switch cur.cond {
	case '!':
		return passps, partset{ min: map[string]int{ "x": 0 }, max: map[string]int{ "x": -1 } }
	case '<':
		passps.setmax(cur.op1, cur.op2 - 1)
		failps.setmin(cur.op1, cur.op2)
		return passps, failps
	case '>':
		passps.setmin(cur.op1, cur.op2 + 1)
		failps.setmax(cur.op1, cur.op2)
		return passps, failps
	}
	panic("blah")
}

func (ps partset) copy() partset {
	min := make(map[string]int)
	for k, v := range ps.min {
		min[k] = v
	}
	max := make(map[string]int)
	for k, v := range ps.max {
		max[k] = v
	}
	return partset{min, max}
}

func (ps partset) setmax(p string, n int) {
	if n < ps.max[p] {
		ps.max[p] = n
	}
}

func (ps partset) setmin(p string, n int) {
	if n > ps.min[p] {
		ps.min[p] = n
	}
}

func sum(part map[string]int) int {
	r := 0
	for _, v := range part {
		r += v
	}
	return r
}

func process(part map[string]int, pos string) bool {
	for {
		nextpos := ""
		for _, rl := range program[pos].rules {
			pass := false
			switch rl.cond {
			case '!':
				pass = true
			case '<':
				if part[rl.op1] < rl.op2 {
					pass = true
				}
			case '>':
				if part[rl.op1] > rl.op2 {
					pass = true
				}
			default:
				panic("blah")
			}
			if pass {
				nextpos = rl.dst
				break
			}
		}
		switch nextpos {
		case "A":
			return true
		case "R":
			return false
		case "":
			panic("blah")
		default:
			pos = nextpos
		}
	}
}

func parsePart(in string) map[string]int {
	if in[0] != '{' || in[len(in)-1] != '}' {
		panic("blah")
	}
	in = in[1:len(in)-1]
	v := Spac(in, ",", -1)
	r := map[string]int{}
	for _, s := range v {
		v := Spac(s, "=", -1)
		r[v[0]] = Atoi(v[1])
	}
	return r
}

func parse(line string) *step {
	v := Spac(line, "{", 2)
	r := &step{ name: v[0] }
	if v[1][len(v[1])-1] != '}' {
		panic("blah")
	}
	r.rules = parseRules(v[1][:len(v[1])-1])
	return r
}

func parseRules(in string) []rule {
	v := Spac(in, ",", -1)
	r := []rule{}
	for i := range v {
		r = append(r, *parseRule(v[i]))
	}
	return r
}

func parseRule(in string) *rule {
	v := Spac(in, ":", -1)
	if len(v) == 1 {
		return &rule{ cond: '!', dst: in }
	}
	dest := v[1]
	in = v[0]
	if v := Spac(in, "<", -1); len(v) == 2 {
		return &rule{ op1: v[0], op2: Atoi(v[1]), cond: '<', dst: dest }
	} else if v = Spac(in, ">", -1); len(v) == 2 {
		return &rule{ op1: v[0], op2: Atoi(v[1]), cond: '>', dst: dest }
	} else {
		panic("blah")
	}
}


func acceptreject(dest string) {
	switch dest {
	case "A", "R":
		//ok
	default:
		panic("blah")
	}
}