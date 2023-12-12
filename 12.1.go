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

var startpattern string
var startnums []int

func main() {
	lines := Input(os.Args[1], "\n", true)
	/*part1 := 0
	for _, line := range lines {
		v := Spac(line, " ", -1)
		pattern := v[0]
		nums := Vatoi(Spac(v[1], ",", -1))
		pln(pattern, nums)
		startpattern = pattern
		startnums = nums
		part1 += enum(pattern, nums, []byte{})
		pln(part1)
		//pln(enum(pattern, nums))
	}
	Sol(part1)*/

	part2 := 0
	for _, line := range lines {
		v := Spac(line, " ", -1)
		pattern := v[0]
		nums := Vatoi(Spac(v[1], ",", -1))
		pattern, nums = unfold(pattern, nums)
		pln(pattern, nums)
		startpattern = pattern
		startnums = nums
		part2 += enum(pattern, nums, []byte{})
		pln(part2)
		//pln(enum(pattern, nums))
	}
	Sol(part2)

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

func enum(pattern string, nums []int, out []byte) int {
	//pf("recursive call %s %v %s\n", pattern, nums, out)
	if len(pattern) == 0 {
		if len(nums) == 0 {
			consistencycheck(string(out))
			return 1
		} else {
			return 0
		}
	}
	if pattern[0] == '?' {
		return enumsub('.', pattern, nums, out) + enumsub('#', pattern, nums, out)
	} else {
		return enumsub(pattern[0], pattern, nums, out)
	}
}

func enumsub(ch byte, pattern string, nums []int, out []byte) int {
	switch ch {
	case '.':
		return enum(pattern[1:], nums, append(out, '.'))
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
			for j := 0; j < nums[0]; j++ {
				out = append(out, '#')
			}
			return enum(pattern[i:], nums[1:], out)
		}
		switch pattern[i] {
		case '#':
			return 0
		case '?', '.':
			for j := 0; j < nums[0]; j++ {
				out = append(out, '#')
			}
			return enum(pattern[i+1:], nums[1:], append(out, '.'))
		default:
			panic("blah")
		}
	default:
		panic("blah")
	}
}

func consistencycheck(out string) {
}

// 12272 wrong
