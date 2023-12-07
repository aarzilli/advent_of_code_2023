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

var strength = map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	v := make([]hand, len(lines))
	for i, line := range lines {
		v[i] = parse(line)
	}
	//pln(v)
	/*for i := range v {
		pln(v[i].h, nameof(typ(v[i].h)))
	}*/
	sort.Slice(v, func(i, j int) bool {
		return better(v[i].h, v[j].h)
	})
	//pln(v)
	part1 := 0
	rank := 1
	for i := len(v) - 1; i >= 0; i-- {
		w := v[i].bid * rank
		pln(v[i].h, nameof(typ(v[i].h)), rank)
		part1 += w
		rank++
	}
	Sol(part1) // 251046173 wrong
}

const (
	fiveofakind  = 7
	fourofakind  = 6
	fullhouse    = 5 // ok
	threeofakind = 4 // ok
	twopair      = 3 // ok
	onepair      = 2 // ok
	highcard     = 1 // ok
)

func nameof(typ int) string {
	switch typ {
	case fiveofakind:
		return "fiveofakind"
	case fourofakind:
		return "fourofakind"
	case fullhouse:
		return "fullhouse"
	case threeofakind:
		return "threeofakind"
	case twopair:
		return "twopair"
	case onepair:
		return "onepair"
	case highcard:
		return "highcard"

	}
	panic("blah")
}

func typ(line string) int {
	m := map[rune]int{}
	for _, ch := range line {
		m[ch]++
	}
	cnts := []int{}
	for ch := range m {
		cnts = append(cnts, m[ch])
	}
	sort.Ints(cnts)
	if len(m) == 1 {
		return fiveofakind
	}
	if len(m) == 2 {
		if cnts[1] == 4 {
			return fourofakind
		}
		if cnts[1] == 3 {
			return fullhouse
		}
		panic("blah")
	}
	if len(m) == 3 {
		if cnts[2] == 3 {
			return threeofakind
		}
		if cnts[2] == 2 {
			return twopair
		}
		panic("blah")
	}
	if len(m) == 4 {
		return onepair
	}
	return highcard
}

func better(a, b string) bool {
	if typ(a) > typ(b) {
		return true
	}
	if typ(a) < typ(b) {
		return false
	}
	for i := range a {
		if strength[a[i:i+1]] > strength[b[i:i+1]] {
			return true
		}
		if strength[a[i:i+1]] < strength[b[i:i+1]] {
			return false
		}
	}
	panic("blah")
}

func parse(line string) hand {
	v := Spac(line, " ", -1)
	return hand{v[0], Atoi(v[1])}
}

type hand struct {
	h   string
	bid int
}
