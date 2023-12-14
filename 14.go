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
var G = map[string]*node{}

type node struct {
	id    int
	name  string
	next  string
	score int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}

	//printmatrix()

	Sol(score())
	cur := lookup()

	cyclefound := false
	offset := 0
	length := 0

	cycle := 1
	for cycle = 1; cycle <= 10000; cycle++ {
		rollnorth()
		rollwest()
		rollsouth()
		rolleast()
		pln("cycle", cycle)
		next := lookup()
		cur.next = next.name
		if !cyclefound {
			if next.next != "" {
				pln("cycle at ", next.id, " length", cur.id-next.id)
				offset = next.id
				length = cur.id - next.id
				cyclefound = true
				cur = next
				break
			} else {
				cur = next
			}
		} else {
			cur = next
			predicted := byid(((cycle+1)-offset)%(length+1) + offset)
			if predicted != cur {
				pln("mismatch predicted=", predicted.id, "got=", cur.id)
				panic("blah")
			}
			pln("prediction ok", predicted.id)
		}
		pln(cur.id, " ", score())
		pln()
	}

	cycle = 100000000
	predicted := byid(((cycle+1)-offset)%(length+1) + offset)
	pln("predicted after million: ", predicted.id)
	Sol(predicted.score)
}

func byid(id int) *node {
	for _, n := range G {
		if n.id == id {
			return n
		}
	}
	pln("no id", id)
	panic("blah")
}

var curid = 0

func lookup() *node {
	s := printmatrix()
	if G[s] != nil {
		return G[s]
	}

	n := &node{
		id:    curid + 1,
		name:  s,
		score: score()}
	curid++
	G[s] = n
	return n
}

func printmatrix() string {
	s := ""
	for i := range M {
		s += string(M[i]) + "\n"
	}
	return s[:len(s)-1]
}

func rollnorth() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'O' {
				movenorth(i, j)
			}
		}
	}
}

func rollsouth() {
	for i := len(M) - 1; i >= 0; i-- {
		for j := range M[i] {
			if M[i][j] == 'O' {
				movesouth(i, j)
			}
		}
	}
}

func rollwest() {
	for j := range M[0] {
		for i := range M {
			if M[i][j] == 'O' {
				moveeast(i, j)
			}
		}
	}
}

func rolleast() {
	for j := len(M[0]) - 1; j >= 0; j-- {
		for i := range M {
			if M[i][j] == 'O' {
				movewest(i, j)
			}
		}
	}
}

func movenorth(i, j int) {
	for i-1 >= 0 {
		if M[i-1][j] == '.' {
			M[i][j] = '.'
			M[i-1][j] = 'O'
			i--
		} else {
			break
		}
	}
}

func movesouth(i, j int) {
	for i+1 < len(M) {
		if M[i+1][j] == '.' {
			M[i][j] = '.'
			M[i+1][j] = 'O'
			i++
		} else {
			break
		}
	}
}

func moveeast(i, j int) {
	for j-1 >= 0 {
		if M[i][j-1] == '.' {
			M[i][j] = '.'
			M[i][j-1] = 'O'
			j--
		} else {
			break
		}
	}
}

func movewest(i, j int) {
	for j+1 < len(M[i]) {
		if M[i][j+1] == '.' {
			M[i][j] = '.'
			M[i][j+1] = 'O'
			j++
		} else {
			break
		}
	}
}

func score() int {
	r := 0
	for i := range M {
		s := len(M) - i
		for j := range M[i] {
			if M[i][j] == 'O' {
				r += s
			}
		}
	}
	return r
}
