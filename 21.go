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

var M [][]byte

type point struct{ i, j int }

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]byte, len(lines))
	var start point
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M[i] {
			if M[i][j] == 'S' {
				start = point{i, j}
			}
		}
	}

	/*fringe := make(Set[point])
	fringe[start] = true

	var N = 6
	if len(M) > 20 {
		N = 64
	}

	for cnt := 0; cnt < N; cnt++ {
		newfringe := make(Set[point])

		add := func(i, j int) {
			if i < 0 || i >= len(M) || j < 0 || j >= len(M[i]) {
				return
			}
			if M[i][j] == '#' {
				return
			}
			newfringe[point{i,j}] = true
		}

		for cur := range fringe {
			add(cur.i-1, cur.j)
			add(cur.i+1, cur.j)
			add(cur.i, cur.j+1)
			add(cur.i, cur.j-1)
		}

		fringe = newfringe
	}

	Sol(len(fringe))*/

	/*for i := -20; i <= 20; i++ {
		pln("i", i, "mod(i)", mod(i, len(M)), "I", tobig(i, len(M)))
	}*/

	//part2dumb(start, 500)
	//findshit(start, 500) -> 13
	//findshit(start, 600) -> 21
	//findshit(start, 700) -> 25
	//findshit(start, 800) -> 25 :)
	//findshit(start, 900) -> 25

	// with simplified entryconf (realinput)
	//findshit(start, 900) // -> 9 (solution: 722900)
	//findshit(start, 750) // -> 9 (solution: 502657)
	//findshit(start, 500) // -> 9 (solution: 223681)
	//part2(500)

	part2dumb(start, 900)
}

func tostrwith(I, J int, fringe Set[point]) string {
	r := []byte{}
	for i := range M {
		for j := range M[i] {
			if fringe[point{i + I*len(M), j + J*len(M[0])}] {
				r = append(r, 'O')
			} else {
				r = append(r, M[i][j])
			}
		}
		r = append(r, '\n')
	}
	return string(r)
}

func ecfor(P point, fringe Set[point]) string {
	return tostrwith(P.i, P.j, fringe)
	/*r := []byte{}
	for dI := -1; dI <= 1; dI++ {
		for i := range M {
			for dJ := -1; dJ <= 1; dJ++ {
				for j := range M[i] {
					if fringe[point{i + (P.i+dI)*len(M), j + (P.j+dJ)*len(M[0])}] {
						r = append(r, 'O')
					} else {
						r = append(r, M[i][j])
					}
				}
			}
			r = append(r, '\n')
		}
	}
	return string(r)*/
}

func get(i, j int) byte {
	i2 := mod(i, len(M))
	j2 := mod(j, len(M[i2]))
	if i2 < 0 || j2 < 0 {
		pln(j, j2, len(M[i2]))
	}
	return M[i2][j2]
}

func mod(n, m int) int {
	return ((n % m) + m) % m
}

func tobig(n, m int) int {
	return (n - mod(n, m)) / m
}

func part2dumb(start point, N int) {
	fringe := make(Set[point])
	fringe[start] = true

	for cnt := 0; cnt < N; cnt++ {
		if cnt%131 == 65 {
			pln(cnt, len(fringe))
		}
		newfringe := make(Set[point])

		add := func(i, j int) {
			if get(i, j) == '#' {
				return
			}
			newfringe[point{i, j}] = true
		}

		for cur := range fringe {
			add(cur.i-1, cur.j)
			add(cur.i+1, cur.j)
			add(cur.i, cur.j+1)
			add(cur.i, cur.j-1)
		}

		fringe = newfringe
	}

	pln(len(fringe))
}

var macrocells = map[point]*macrocell{}
var mckinds = map[string]*macrocell{} // macrocell kind for entry configuration string

type macrocell struct {
	discovered     int
	seen           map[string]int
	loopat, loopto int
	entryconf      string

	exitat map[point]int    // exits to point at count cnt
	exitto map[point]string // exits to point in entry configuration string
}

func findshit(start point, N int) {
	fringe := make(Set[point])
	fringe[start] = true

	for cnt := 0; cnt < N; cnt++ {
		macrocheck(fringe, cnt)
		newfringe := make(Set[point])

		add := func(P point, i, j int) {
			if get(i, j) == '#' {
				return
			}
			var Q point
			Q.i = tobig(i, len(M))
			Q.j = tobig(j, len(M))
			if macrocells[Q] == nil {
				if _, ok := macrocells[P].exitat[Q]; !ok {
					macrocells[P].exitat[Q] = cnt
				}
			}
			newfringe[point{i, j}] = true
		}

		for cur := range fringe {
			var P point
			P.i = tobig(cur.i, len(M))
			P.j = tobig(cur.j, len(M[0]))
			add(P, cur.i-1, cur.j)
			add(P, cur.i+1, cur.j)
			add(P, cur.i, cur.j+1)
			add(P, cur.i, cur.j-1)
		}

		fringe = newfringe
	}

	findmckinds()

	/*
		pln("Macrocells:")
		for P, mc := range macrocells {
			if mc.loopat != 0 {
				pln(P, "discovered:", mc.discovered, "loopat", mc.loopat, mc.loopto)
				pln("entry:")
				pln(mc.getseen(mc.discovered))
				pln()
			}
		}*/

	pln(len(fringe))
}

func findmckinds() {
	for P, mc := range macrocells {
		if mc.loopat != 0 {
			mckind := mc2mckind(P, mc)
			if mckinds[mckind.entryconf] == nil {
				mckinds[mckind.entryconf] = mckind
			}
			assertmckindeq(mckinds[mckind.entryconf], mckind)
		}
	}
	pln("num kinds", len(mckinds))
}

func mc2mckind(P point, mc *macrocell) *macrocell {
	d := mc.discovered
	mckind := &macrocell{
		entryconf:  mc.entryconf,
		discovered: 0,
		loopat:     mc.loopat - d, loopto: mc.loopto - d,
		seen:   make(map[string]int),
		exitat: make(map[point]int),
		exitto: make(map[point]string),
	}

	for s, cnt := range mc.seen {
		mckind.seen[s] = cnt - d
	}

	for Q, cnt := range mc.exitat {
		mckind.exitat[point{Q.i - P.i, Q.j - P.j}] = cnt - d
	}

	for Q, s := range mc.exitto {
		mckind.exitto[point{Q.i - P.i, Q.j - P.j}] = s
	}

	return mckind
}

func assertmckindeq(a, b *macrocell) {
	if a.discovered != b.discovered || a.loopat != b.loopat || a.loopto != b.loopto || len(a.seen) != len(b.seen) || len(a.exitat) != len(b.exitat) || len(a.exitto) != len(b.exitto) {
		panic("imfucked")
	}
	for s := range a.seen {
		if a.seen[s] != b.seen[s] {
			panic("imfucked")
		}
	}
	for Q := range a.exitat {
		if a.exitat[Q] != b.exitat[Q] {
			panic("imfucked")
		}
	}
	for Q := range a.exitto {
		if a.exitto[Q] != b.exitto[Q] {
			panic("imfucked")
		}
	}
}

func macrocheck(fringe Set[point], cnt int) {
	macroset := make(Set[point])
	for p := range fringe {
		var P point
		P.i = tobig(p.i, len(M))
		P.j = tobig(p.j, len(M[0]))
		macroset[P] = true
	}

	for P := range macroset {
		mc := macrocells[P]
		if mc == nil {
			mc = &macrocell{
				entryconf:  ecfor(P, fringe),
				discovered: cnt,
				seen:       make(map[string]int),
				exitat:     make(map[point]int),
				exitto:     make(map[point]string),
			}
			macrocells[P] = mc
		}
		if mc.loopat > 0 {
			continue
		}

		for Q := range mc.exitat {
			if mc.exitto[Q] == "" {
				mc.exitto[Q] = ecfor(Q, fringe)
			}
		}

		sig := tostrwith(P.i, P.j, fringe)
		if _, ok := mc.seen[sig]; ok {
			mc.loopat = cnt
			mc.loopto = mc.seen[sig]
		} else {
			mc.seen[sig] = cnt
		}
	}
}

type qentry struct {
	at        int
	entryconf string
}

func part2(N int) {
	queue := []qentry{{0, macrocells[point{0, 0}].entryconf}}
	cnt := 0
	solmcks := []qentry{}
	for queue[0].at < N {
		cnt = queue[0].at
		mck := mckinds[queue[0].entryconf]
		solmcks = append(solmcks, queue[0])
		queue = queue[1:]
		for Q := range mck.exitat {
			queue = append(queue, qentry{cnt + mck.exitat[Q], mck.exitto[Q]})
		}
	}

	cnt = N
	r := 0
	// bugged somehow...
	for _, solmck := range solmcks {
		mck := mckinds[solmck.entryconf]
		if mck.loopat+solmck.at > cnt {
			r += countO(mck.getseen(cnt - solmck.at))
		} else {
			d := (cnt - (mck.loopat + solmck.at)) % (mck.loopat - mck.loopto)
			r += countO(mck.getseen(mck.loopto + d))
		}
	}
	pln("calculated", r)
}

func countO(s string) int {
	r := 0
	for i := range s {
		if s[i] == 'O' {
			r++
		}
	}
	return r
}

func (mc *macrocell) getseen(cnt int) string {
	for s, k := range mc.seen {
		if k == cnt {
			return s
		}
	}
	panic("blah")
}
