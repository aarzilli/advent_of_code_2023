package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"slices"
	_ "sort"
	"strconv"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type point struct{ i, j int }
type instr struct {
	dir   byte
	count int
	color string
}

type intvl struct {
	i          int
	start, end int
	filled     bool
}

type line struct {
	width  int
	intvls []intvl
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	instrs := []instr{}
	for _, line := range lines {
		v := Spac(line, " ", -1)
		instrs = append(instrs, instr{
			dir:   v[0][0],
			count: Atoi(v[1]),
			color: v[2][2 : len(v[2])-1],
		})
	}

	Sol(part1(instrs))

	digitdir := map[byte]byte{
		'0': 'R',
		'1': 'D',
		'2': 'L',
		'3': 'U',
	}

	instrs2 := []instr{}

	for i := range instrs {
		n, _ := strconv.ParseInt(instrs[i].color[:5], 16, 0)
		instrs2 = append(instrs2, instr{dir: digitdir[instrs[i].color[5]], count: int(n)})
	}

	Sol(part1(instrs2))
}

func part1(instrs []instr) int {
	verts := []intvl{}
	var maplines = map[int]*line{}

	mini, maxi, minj, maxj := 0, 0, 0, 0
	border := 0

	cur := point{0, 0}
	for _, instr := range instrs {
		var di, dj int
		switch instr.dir {
		case 'R':
			dj = 1
		case 'L':
			dj = -1
		case 'U':
			di = -1
		case 'D':
			di = +1
		}

		if dj != 0 {
			if maplines[cur.i] == nil {
				maplines[cur.i] = &line{width: 1}
			}
		}

		if dj > 0 {
			maplines[cur.i].intvls = append(maplines[cur.i].intvls, intvl{cur.i, cur.j, cur.j + (dj * instr.count) + 1, true})
		} else if dj < 0 {
			maplines[cur.i].intvls = append(maplines[cur.i].intvls, intvl{cur.i, cur.j + (dj * instr.count), cur.j + 1, true})
		} else if di > 0 {
			verts = append(verts, intvl{cur.j, cur.i, cur.i + (di * instr.count) + 1, true})
		} else if di < 0 {
			verts = append(verts, intvl{cur.j, cur.i + (di * instr.count), cur.i + 1, true})
		}

		cur.i += di * instr.count
		cur.j += dj * instr.count

		border += instr.count

		if cur.i < mini {
			mini = cur.i
		}
		if cur.i > maxi {
			maxi = cur.i
		}
		if cur.j < minj {
			minj = cur.j
		}
		if cur.j > maxj {
			maxj = cur.j
		}
	}

	for i := range maplines {
		line := maplines[i]
		if line.width != 1 {
			continue
		}
		drawverts(maplines, i, verts)
		drawverts(maplines, i+1, verts)
	}

	for i := range maplines {
		if maplines[i].width == 0 {
			maplines[i].width = nextline(maplines, i) - i
		}
	}

	for _, line := range maplines {
		slices.SortFunc(line.intvls, func(a, b intvl) int { return a.start - b.start })
		newintvls := []intvl{}
		for k := 0; k < len(line.intvls); k++ {
			if k+1 < len(line.intvls) {
				cur := &line.intvls[k]
				next := &line.intvls[k+1]
				newintvls = append(newintvls, intvl{cur.i, cur.end, next.start, false})
			}
		}
		line.intvls = append(line.intvls, newintvls...)
	}

	/*	idxs := Keys(maplines)
		sort.Ints(idxs)
		for _, i := range idxs {
			pln(i, ":", maplines[i])
		}
		for i := mini; i <= maxi; i++ {
			if maplines[i] == nil {
				continue
			}
			pf("%-20s", fmt.Sprintf("%d:%02d", i, maplines[i].width))
			for j := minj; j <= maxj; j++ {
				pf("%c", marked(i, j))
			}
			pln()
		}*/

	interior0 := maplines[mini].intvls[0]
	interior0.start++
	interior0.end--
	//pln(interior0)
	tofill := findintersectingempty(maplines, mini+1, interior0)
	part1 := 0
	for len(tofill) > 0 {
		cur := tofill[len(tofill)-1]
		tofill = tofill[:len(tofill)-1]
		if cur.filled {
			continue
		}
		//pln("filling", cur)
		cur.filled = true
		part1 += (cur.end - cur.start) * maplines[cur.i].width
		tofill = append(tofill, findintersectingempty(maplines, cur.i+maplines[cur.i].width, *cur)...)
		tofill = append(tofill, findintersectingempty(maplines, cur.i-1, *cur)...)
	}

	return part1 + border
}

/*func marked(i, j int) byte {
	if maplines[i] == nil {
		return ' '
	}
	for _, intvl := range maplines[i].intvls {
		if intvl.contains(j)  {
			if intvl.filled {
				return '#'
			} else {
				return '.'
			}
		}
	}
	return ' '
}*/

func drawvert(maplines map[int]*line, i, j int) {
	if maplines[i] == nil {
		maplines[i] = &line{}
	}
	for _, intvl := range maplines[i].intvls {
		if intvl.contains(j) {
			return
		}
	}
	maplines[i].intvls = append(maplines[i].intvls, intvl{i, j, j + 1, true})
}

func nextline(maplines map[int]*line, i int) int {
	mini2 := i
	first := true
	for i2, _ := range maplines {
		if i2 > i && (first || i2 < mini2) {
			first = false
			mini2 = i2
		}
	}
	return mini2
}

func drawverts(maplines map[int]*line, i int, verts []intvl) {
	for _, vert := range verts {
		if vert.contains(i) {
			drawvert(maplines, i, vert.i)
		}
	}
}

func findconsecutive(v []int) []int {
	for i := 1; i < len(v); i++ {
		if v[i] != v[i-1]+1 {
			return v[:i]
		}
	}
	return v
}

func findintersectingempty(maplines map[int]*line, tgti int, prev intvl) []*intvl {
	r := []*intvl{}
	for i := range maplines {
		if i <= tgti && tgti < i+maplines[i].width {
			for k := range maplines[i].intvls {
				intvl := &maplines[i].intvls[k]
				if !intvl.filled && (intvl.contains(prev.start) || intvl.contains(prev.end-1) || prev.contains(intvl.start) || prev.contains(intvl.end-1)) {
					r = append(r, intvl)
				}
			}
			break
		}
	}
	return r
}

func (a *intvl) contains(p int) bool {
	return a.start <= p && p < a.end
}

/*func floodfill(start point) {
	fringe := make(Set[point])
	fringe[start] = true
	for len(fringe) > 0 {
		cur := OneKey(fringe)
		delete(fringe, cur)
		if M[cur] {
			continue
		}
		M[cur] = true
		fringe[point{cur.i-1, cur.j}] = true
		fringe[point{cur.i+1, cur.j}] = true
		fringe[point{cur.i, cur.j-1}] = true
		fringe[point{cur.i, cur.j+1}] = true
	}
}*/
