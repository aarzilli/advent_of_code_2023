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

var conf = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	r := 0
	r2 := 0
	for _, line := range lines {
		v := Spac(line, ":", 2)
		//pf("%q\n", v)
		gameid := Getints(v[0], false)[0]
		//pf("gameid: %d\n", gameid)
		v = Spac(v[1], ";", -1)
		valid := true
		smallest := map[string]int{}
		for _, set := range v {
			//pf("%q\n", set)
			if !validset(set, smallest) {
				valid = false
			}
		}
		if valid {
			//pf("valid\n")
			r += gameid
		}
		pwr := 1
		for _, k := range smallest {
			pwr *= k
		}
		//pf("%d %v %d\n", gameid, smallest, pwr)
		r2 += pwr
		
	}
	Sol(r)
	Sol(r2)
}

func validset(set string, smallest map[string]int) bool {
	v := Spac(set, ",", -1)
	valid := true
	for _, f := range v {
		if !validpair(f, smallest) {
			//pf("invalid %q\n", f)
			valid = false
		}
	}
	return valid
}

func validpair(f string, smallest map[string]int) bool {
	v := Spac(f, " ", -1)
	n := Atoi(v[0])
	if s, ok := smallest[v[1]]; !ok || s < n {
		smallest[v[1]] = n
	}
	//pf("%d %q %v\n", n, v[1], n <= conf[v[1]])
	return n <= conf[v[1]]
}