package main

import (
	. "aoc/util"
	"os"
)

var maps = make(map[string][]mapping)

type mapping []int

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	seeds := Vatoi(Spac(Spac(groups[0], ":", -1)[1], " ", -1))
	for i := 1; i < len(groups); i++ {
		group := groups[i]
		lines := Spac(group, "\n", -1)
		name := Spac(lines[0], " ", -1)[0]

		for j := 1; j < len(lines); j++ {
			maps[name] = append(maps[name], Vatoi(Spac(lines[j], " ", -1)))
		}
	}

	minloc := -1
	for _, seed := range seeds {
		p := seedtoloc(seed, nil)
		if minloc == -1 || p < minloc {
			minloc = p
		}
	}

	Sol(minloc, 331445006)

	minloc = -1
	for i := 0; i < len(seeds); i += 2 {
		seedstart := seeds[i]
		seedend := seedstart + seeds[i+1]

		seed := seedstart
		for seed < seedend {
			rng := seedend - seed
			p := seedtoloc(seed, &rng)
			if minloc == -1 || p < minloc {
				minloc = p
			}
			seed += rng
		}
	}

	Sol(minloc, 6472060)
}

func applymap(amap []mapping, in int, prng *int) int {
	for i := range amap {
		if amap[i][1] <= in && in < amap[i][1]+amap[i][2] {
			if prng != nil {
				rng := amap[i][1] + amap[i][2] - in
				if rng < *prng {
					*prng = rng
				}
			}
			return amap[i][0] + in - amap[i][1]
		}
	}
	if prng != nil {
		// find next mapping to update prng
		for i := range amap {
			if in < amap[i][1] {
				rng := amap[i][1] - in
				if rng < *prng {
					*prng = rng
				}
			}
		}
	}
	return in
}

func seedtoloc(seed int, prng *int) int {
	p := seed
	for _, mapname := range []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"} {
		newp := applymap(maps[mapname], p, prng)
		p = newp
	}
	return p
}
