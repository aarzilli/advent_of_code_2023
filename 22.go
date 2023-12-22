package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"slices"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type point struct {
	x, y, z int
}

type brick struct {
	a, b point
}

var bricks []brick

func main() {
	lines := Input(os.Args[1], "\n", true)
	for _, line := range lines {
		v := Getints(line, false)
		bricks = append(bricks, brick{point{v[0], v[1], v[2]}, point{v[3], v[4], v[5]}})
	}
	slices.SortFunc(bricks, func(a, b brick) int { return a.a.z - b.a.z })
	for i := range bricks {
		drop(&bricks[i], bricks[:i], -1)
	}
	slices.SortFunc(bricks, func(a, b brick) int { return a.a.z - b.a.z })

	for i := range bricks {
		b := bricks[i]
		pf("%d,%d,%d~%d,%d,%d\n", b.a.x, b.a.y, b.a.z, b.b.x, b.b.y, b.b.z)
	}

	forpart2 := []int{}
	part1 := 0
	for i := range bricks {
		if destroycheck(i) {
			part1++
		} else {
			forpart2 = append(forpart2, i)
		}
	}
	Sol(part1)

	part2 := 0
	for _, i := range forpart2 {
		part2 += fallen(i)
	}
	Sol(part2)
}

func fallen(i int) int {
	bricks2 := make([]brick, len(bricks))
	copy(bricks2, bricks)
	for j := range bricks2 {
		drop(&bricks2[j], bricks2[:j], i)
	}
	r := 0
	for j := range bricks2 {
		if bricks2[j] != bricks[j] {
			r++
		}
	}
	return r
}

func destroycheck(i int) bool {
	s, e := restingon(bricks[i], i+1)
	candestroy := true
	for j := s; j < e; j++ {
		br2 := bricks[j]
		br2c := br2
		drop(&br2, bricks[:s], i)
		if br2 != br2c {
			candestroy = false
			break
		}
	}

	//pln(candestroy)
	return candestroy
}

func restingon(br brick, k int) (start, end int) {
	start = -1
	for i := k; i < len(bricks); i++ {
		if bricks[i].a.z == br.b.z+1 && start == -1 {
			start = i
		}
		if bricks[i].a.z > br.b.z+1 {
			return start, i
		}
	}
	if start < 0 {
		return 0, 0
	}
	return start, len(bricks)
}

func drop(br *brick, tocheck []brick, skip int) {
	for br.a.z > 1 {
		br.a.z--
		br.b.z--
		stop := false
		for i := len(tocheck) - 1; i >= 0; i-- {
			if i == skip {
				continue
			}
			br2 := &tocheck[i]
			if intersect(*br, *br2) {
				stop = true
				break
			}
		}
		if stop {
			br.a.z++
			br.b.z++
			break
		}
	}
}

func intersect(a, b brick) bool {
	if a.a.x > b.b.x || a.b.x < b.a.x || a.a.y > b.b.y || a.b.y < b.a.y ||
		a.a.z > b.b.z || a.b.z < b.a.z {
		return false
	}
	return true
}
