package main

import (
	. "aoc/util"
	"os"
	"slices"
	"sort"
)

type hand struct {
	h   string
	bid int
}

var strength1 = map[string]int{"A": 13, "K": 12, "Q": 11, "J": 10, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1}
var strength2 = map[string]int{"A": 13, "K": 12, "Q": 11, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1, "J": 0}
var typestrength = map[string]int{"fiveofakind": 7, "fourofakind": 6, "fullhouse": 5, "threeofakind": 4, "twopair": 3, "onepair": 2, "highcard": 1}

func main() {
	lines := Input(os.Args[1], "\n", true)
	v := make([]hand, len(lines))
	for i, line := range lines {
		v[i] = parse(line)
	}
	sort.Slice(v, func(i, j int) bool { return !better(v[i].h, v[j].h, false) })
	Sol(score(v))
	sort.Slice(v, func(i, j int) bool { return !better(v[i].h, v[j].h, true) })
	Sol(score(v))
}

func parse(line string) hand {
	v := Spac(line, " ", -1)
	return hand{v[0], Atoi(v[1])}
}

func typ(line string, part2 bool) string {
	m := map[rune]int{}
	for _, ch := range line {
		m[ch]++
	}

	if part2 {
		if m['J'] > 0 {
			j := m['J']
			delete(m, 'J')

			highest := ' '
			for ch := range m {
				if m[ch] > m[highest] {
					highest = ch
				}
			}
			if highest == ' ' {
				highest = 'J'
			}
			m[highest] += j
		}
	}

	cnts := []int{}
	for ch := range m {
		cnts = append(cnts, m[ch])
	}
	sort.Ints(cnts)
	slices.Reverse(cnts)
	switch {
	case equals(cnts, 5):
		return "fiveofakind"
	case equals(cnts, 4, 1):
		return "fourofakind"
	case equals(cnts, 3, 2):
		return "fullhouse"
	case equals(cnts, 3, 1, 1):
		return "threeofakind"
	case equals(cnts, 2, 2, 1):
		return "twopair"
	case equals(cnts, 2, 1, 1, 1):
		return "onepair"
	default:
		return "highcard"
	}
}

func equals(a []int, b ...int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func better(a, b string, part2 bool) bool {
	ta := typestrength[typ(a, part2)]
	tb := typestrength[typ(b, part2)]
	if ta != tb {
		return ta > tb
	}
	strength := strength1
	if part2 {
		strength = strength2
	}
	for i := range a {
		sa := strength[a[i:i+1]]
		sb := strength[b[i:i+1]]
		if sa != sb {
			return sa > sb
		}
	}
	panic("blah")
}

func score(v []hand) (r int) {
	for i := range v {
		w := v[i].bid * (i + 1)
		r += w
	}
	return r
}
