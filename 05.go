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

var maps = make(map[string][]mapping)

type mapping []int

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	pf("len %d\n", len(groups))
	seeds := Vatoi(Spac(Spac(groups[0], ":", -1)[1], " ", -1))
	pln(seeds)
	for i := 1; i < len(groups); i++ {
		group := groups[i]
		lines := Spac(group, "\n", -1)
		name := Spac(lines[0], " ", -1)[0]
		//pln(name)
		for j := 1; j < len(lines); j++ {
			line := lines[j]
			maps[name] = append(maps[name], Vatoi(Spac(line, " ", -1)))
		}
	}
	//pln(maps)

	minloc := -1

	for _, seed := range seeds {
		p := seedtoloc(seed)
		if minloc == -1 || p < minloc {
			minloc = p
		}
	}

	Sol(minloc)

	minloc = -1
	for i := 0; i < len(seeds); i += 2 {
		seedstart := seeds[i]
		n := seeds[i+1]

		for j := 0; j < n; j++ {
			seed := seedstart + j
			p := seedtoloc(seed)
			if minloc == -1 || p < minloc {
				minloc = p
			}
		}
	}

	Sol(minloc)
}

func applymap(amap []mapping, seeds int) int {
	for i := range amap {
		if amap[i][1] <= seeds && seeds < amap[i][1]+amap[i][2] {
			return amap[i][0] + seeds - amap[i][1]
		}
	}
	return seeds
}

func seedtoloc(seed int) int {
	p := seed
	for _, mapname := range []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"} {
		np := applymap(maps[mapname], p)
		//pln(p, mapname, np)
		p = np
	}
	//pln("location", p)
	return p
}
