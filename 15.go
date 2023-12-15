package main

import (
	. "aoc/util"
	"os"
	"slices"
)

type lens struct {
	kind string
	pwr  int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1 := 0
	for _, x := range Spac(lines[0], ",", -1) {
		part1 += hash(x)
	}
	Sol(part1)

	fl := map[string]int{}
	order := map[string]int{}

	for i, x := range Spac(lines[0], ",", -1) {
		v := Spac(x, "=", 2)
		if len(v) == 2 {
			if _, ok := fl[v[0]]; !ok {
				order[v[0]] = i
			}
			fl[v[0]] = Atoi(v[1])
		} else {
			delete(fl, x[:len(x)-1])
			delete(order, x[:len(x)-1])
		}
	}

	orderv := []string{}
	for k := range order {
		orderv = append(orderv, k)
	}
	slices.SortFunc(orderv, func(a, b string) int { return order[a] - order[b] })
	part2 := 0
	boxes := [256]int{}
	for _, k := range orderv {
		h := hash(k)
		pos := boxes[h]
		boxes[h]++
		part2 += (h + 1) * (pos + 1) * fl[k]
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
