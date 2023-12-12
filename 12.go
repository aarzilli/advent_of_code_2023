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
	part1 := 0
	for _, line := range lines {
		v := Spac(line, " ", -1)
		pattern := v[0]
		nums := Vatoi(Spac(v[1], ",", -1))
		pattern, nums = unfold(pattern, nums)
		pln(pattern, nums)
		part1 += enum0(pattern, nums, []byte{})
		pln(part1)
		//pln(enum(pattern, nums))
	}
	Sol(part1)

	/*part2 := 0
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
	Sol(part2)*/

}

func enum0(pattern string, nums []int, out []byte) int {
	v := Noempty(Spac(pattern, ".", -1))
	return enum0_2("", v, nums, out)
}

type enum0_2key string

var memoize0_2 = map[enum0_2key]int{}

func tokey0_2(v []string, nums []int) enum0_2key {
	strnums := make([]string, len(nums))
	for i := range nums {
		strnums[i] = fmt.Sprintf("%d", nums[i])
	}
	return enum0_2key(strings.Join(v, ".") + " " + strings.Join(strnums, ","))
}

func enum0_2(depth string, v []string, nums []int, out []byte) int {
	k := tokey0_2(v, nums)
	if r, ok := memoize0_2[k]; ok {
		return r
	}
	//pf("%senum0_2 %q %v\n", depth, v, nums)
	if len(v) == 0 {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if len(nums) == 0 {
		a := enum(v[0], []int{}, out)
		//pf("%senum of %q with %v is %d\n", depth, v[0], nums, a)
		if a == 0 {
			return 0
		}
		r := a * enum0_2(depth+"   ", v[1:], []int{}, out)
		memoize0_2[k] = r
		return r
	}
	z := 0
	for i := 0; i < len(nums)+1; i++ {
		a := enum(v[0], nums[:i], out)
		//pf("%senum of %q with %v is %d\n", depth, v[0], nums[:i], a)
		if a == 0 {
			continue
		}
		z += a * enum0_2(depth+"   ", v[1:], nums[i:], out)
	}
	//pf("%s-> %d\n", depth, z)
	memoize0_2[k] = z
	return z
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

func enum(pattern string, nums []int, out []byte) int {
	//pf("recursive call %s %v %s\n", pattern, nums, out)
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
		r := enumsub('.', pattern, nums, out) + enumsub('#', pattern, nums, out)
		memoize[k] = r
		return r
	} else {
		r := enumsub(pattern[0], pattern, nums, out)
		memoize[k] = r
		return r
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

// 12272 wrong
