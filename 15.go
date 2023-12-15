package main

import (
	. "aoc/util"
	"os"
	"strings"
)

type lens struct {
	kind string
	pwr int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1 := 0
	for _, x := range Spac(lines[0], ",", -1) {
		part1 += hash(x)
	}
	Sol(part1)
	
	box := make([][]lens, 256)
	
	for _, x := range Spac(lines[0], ",", -1) {
		var tgt string
		var op string
		var arg int
		
		eq := strings.Index(x, "=")
		if eq >= 0 {
			v := Spac(x, "=", 2)
			op = "="
			tgt = v[0]
			arg = Atoi(v[1])
		} else {
			if x[len(x)-1] != '-' {
				panic("blah")
			}
			tgt = x[:len(x)-1]
			op = "-"
		}
		
		tgtbox := hash(tgt)
		
		switch op {
		case "=":
			found := false
			for i := range box[tgtbox] {
				if box[tgtbox][i].kind == tgt {
					box[tgtbox][i].pwr = arg
					found = true
				}
			}
			if !found {
				box[tgtbox] = append(box[tgtbox], lens{ tgt, arg })
			}
		case "-":
			idx := -1
			for i := range box[tgtbox] {
				if box[tgtbox][i].kind == tgt {
					idx = i
					break
				}
			}
			if idx >= 0 {
				copy(box[tgtbox][idx:], box[tgtbox][idx+1:])
				box[tgtbox] = box[tgtbox][:len(box[tgtbox])-1]
			}
		}
	}
	
	part2 := 0
	for i := range box {
		for j := range box[i] {
			p := (i+1) * (j+1) * box[i][j].pwr
			part2 += p
		}
	}
	Sol(part2)
}

func hash(seed string) int {
	r := 0
	for i := range seed {
		r += int(seed[i])
		r *= 17
		r %= 256
	}
	return r
}
