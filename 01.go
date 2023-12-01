package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

const part2 = true

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	r := 0
	for _, line := range lines {
		x := []int{}
		/*for _, ch := range line {
			if ch >= '0' && ch <= '9' {
				x = append(x, ch)
			}
		}*/
		for i := 0; i < len(line); i++ {
			ch := line[i]
			if ch >= '0' && ch <= '9' {
				x = append(x, int(ch-'0'))
			}
			if part2 {
				for j, number := range []string{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" } {
					if strings.HasPrefix(line[i:], number) {
						x = append(x, j+1)
						break
					}
				}
			}
		}
		pln(x, x[0], x[len(x)-1])
		r += x[0]*10 + x[len(x)-1]
	}
	Sol(r)
}
