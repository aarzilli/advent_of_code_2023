package main

import (
	. "aoc/util"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Adj := map[string]map[string]bool{}
	for _, line := range lines {
		v := Splitsimilar(line, SSRemoveSymbols)
		for i := 1; i < len(v); i++ {
			add(Adj, v[0], v[i])
			add(Adj, v[i], v[0])
		}
	}
	/*pf("graph G {\n")
	for cur := range Adj {
		for next := range Adj[cur] {
			pf("%s -- %s\n", cur, next)
		}
	}
	pf("}\n")*/
	/*for cur := range Adj {
		if len(Adj) < 3 {
			pln(cur, len(Adj))
		}
	}*/

	// different sides
	/*start := "cmg"
	end := "bvb"*/

	// same side
	/*start := "qnr"
	end := "frs"*/

	/*
		start := "hfx"
		end := "rhn"
		path := search(Adj, start, end, make(Set[string]))
		pln(start, end, ":", path)
		adj2 := copygraph(Adj)
		removedgs(adj2, path)
		path2 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path2)
		removedgs(adj2, path2)
		path3 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path3)
		removedgs(adj2, path3)
		path4 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path4)
	*/

	//pln(findcandidates(Adj))

	if len(os.Args) < 4 {
		delete(Adj["mhb"], "zqg")
		delete(Adj["zqg"], "mhb")

		delete(Adj["jlt"], "sjr")
		delete(Adj["sjr"], "jlt")

		edgefreq := map[[2]string]int{}

		for i := 0; i < 10000; i++ {
			edges := make(Set[[2]string])
			mst(Adj, randomnode(Adj), make(Set[string]), edges)
			for edge := range edges {
				edgefreq[edge]++
			}
		}
		printbest(edgefreq, 200)
	} else {
		pln("testing", os.Args[2:])
		{
			seen := make(Set[string])
			mst(Adj, randomnode(Adj), seen, make(Set[[2]string]))
			pln("before", len(seen), len(Adj))
		}
		for _, edgstr := range os.Args[2:] {
			edge := Spac(edgstr, ",", 2)
			if !Adj[edge[0]][edge[1]] {
				panic("nonexistent")
			}
			delete(Adj[edge[0]], edge[1])
			delete(Adj[edge[1]], edge[0])
		}
		seen := make(Set[string])
		mst(Adj, randomnode(Adj), seen, make(Set[[2]string]))
		pln("after", len(seen), len(Adj))
	}
}

func printbest(edgefreq map[[2]string]int, lim int) {
	if len(edgefreq) == 0 {
		return
	}
	edges := Keys(edgefreq)
	slices.SortFunc(edges, func(a, b [2]string) int { return -edgefreq[a] + edgefreq[b] })

	for _, edge := range edges[:min(len(edges), lim)] {
		pln(edge, edgefreq[edge])
	}
}

func mst(adj map[string]map[string]bool, cur string, seen Set[string], edges Set[[2]string]) {
	seen[cur] = true
	nbs := Keys(adj[cur])
	rand.Shuffle(len(nbs), func(i, j int) { nbs[i], nbs[j] = nbs[j], nbs[i] })
	for _, next := range nbs {
		if !seen[next] {
			addedge(edges, cur, next)
			mst(adj, next, seen, edges)
		}
	}
}

func addedge(edges Set[[2]string], a, b string) {
	r := [2]string{a, b}
	sort.Strings(r[:])
	edges[r] = true
}

func findcandidates(Adj map[string]map[string]bool) Set[string] {
	for {
		start := randomnode(Adj)
		end := randomnode(Adj)
		path := search(Adj, start, end, make(Set[string]))
		pln(start, end, ":", path)
		adj2 := copygraph(Adj)
		removedgs(adj2, path)
		path2 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path2)
		removedgs(adj2, path2)
		path3 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path3)
		removedgs(adj2, path3)
		path4 := search(adj2, start, end, make(Set[string]))
		pln(start, end, ":", path4)
		if path4 == nil {
			return pathunion(path, path2, path3)
		}
	}
}

func search(Adj map[string]map[string]bool, cur, end string, seen Set[string]) []string {
	seen[cur] = true

	if cur == end {
		return []string{cur}
	}

	for next := range Adj[cur] {
		if !seen[next] {
			r := search(Adj, next, end, seen)
			if r != nil {
				return append(r, cur)
			}
		}
	}
	return nil
}

func add(Adj map[string]map[string]bool, a, b string) {
	if Adj[a] == nil {
		Adj[a] = make(map[string]bool)
	}
	Adj[a][b] = true
}

func randomnode(Adj map[string]map[string]bool) string {
	v := Keys(Adj)
	return v[rand.Intn(len(v))]
}

func copygraph(in map[string]map[string]bool) map[string]map[string]bool {
	out := make(map[string]map[string]bool)
	for cur := range in {
		out[cur] = make(map[string]bool)
		for next := range in[cur] {
			out[cur][next] = true
		}
	}
	return out
}

func removedgs(adj map[string]map[string]bool, path []string) {
	for i := len(path) - 1; i > 0; i-- {
		delete(adj[path[i]], path[i-1])
		delete(adj[path[i-1]], path[i])
	}
}

func pathunion(path, path2, path3 []string) Set[string] {
	return Union(Union(path2edgeset(path), path2edgeset(path2)), path2edgeset(path3))
}

func path2edgeset(path []string) Set[string] {
	r := make(Set[string])
	for i := len(path) - 1; i > 0; i-- {
		v := []string{path[i], path[i-1]}
		sort.Strings(v)
		r[strings.Join(v, ",")] = true
	}
	return r
}
