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

func main() {
	lines := Input(os.Args[1], "\n", true)
	part1, part2 := 0, 0
	for _, line := range lines {
		pattern, nums := parse(line)
		part1 += enum(pattern, nums)
		pattern, nums = unfold(pattern, nums)
		part2 += enum(pattern, nums)
	}
	Sol(part1)
	Sol(part2)

}

func parse(line string) (string, []int) {
	v := Spac(line, " ", -1)
	pattern := v[0]
	nums := Vatoi(Spac(v[1], ",", -1))
	return pattern, nums
}

func unfold(s string, nums []int) (string, []int) {
	nums2 := append([]int{}, nums...)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)
	nums2 = append(nums2, nums...)
	const sep = "?"
	return s + sep + s + sep + s + sep + s + sep + s, nums2
}

var memoize = map[string]int{}

func tokey(pattern string, nums []int) string {
	v := make([]string, len(nums))
	for i := range nums {
		v[i] = fmt.Sprintf("%d", nums[i])
	}
	return pattern + " " + strings.Join(v, ",")
}

func enum(pattern string, nums []int) int {
	k := tokey(pattern, nums)
	if r, ok := memoize[k]; ok {
		return r
	}
	if len(pattern) == 0 {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if pattern[0] == '?' {
		r := enumsub('.', pattern, nums) + enumsub('#', pattern, nums)
		memoize[k] = r
		return r
	} else {
		r := enumsub(pattern[0], pattern, nums)
		memoize[k] = r
		return r
	}
}

func enumsub(ch byte, pattern string, nums []int) int {
	switch ch {
	case '.':
		return enum(pattern[1:], nums)
	case '#':
		// must consume number
		if len(nums) == 0 {
			return 0
		}
		ok := true
		var i int
		for i = 0; i < len(pattern) && i < nums[0]; i++ {
			if pattern[i] != '#' && pattern[i] != '?' {
				ok = false
				break
			}
		}
		if !ok {
			return 0
		}
		if i == len(pattern) {
			if len(pattern) < nums[0] {
				return 0
			}
			return enum(pattern[i:], nums[1:])
		}
		switch pattern[i] {
		case '#':
			return 0
		case '?', '.':
			return enum(pattern[i+1:], nums[1:])
		default:
			panic("blah")
		}
	default:
		panic("blah")
	}
}
