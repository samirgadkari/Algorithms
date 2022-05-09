package main

import (
	"fmt"
)

type permute func(*pHelper)
type pHelper struct {
	nums        []int
	ignore      []int
	numIgnore   int
	keepInitial bool
}

func findTwoNonIgnoredIdx(helper *pHelper) (int, int) {

	idx := make([]int, 0)

	for i, v := range helper.ignore {
		if v == 0 {
			idx = append(idx, i)
		}
	}

	return idx[0], idx[1]
}

func permuteFunc(digits []int) (<-chan []int, permute) {

	result := make(chan []int)

	var r []int
	var p permute
	p = func(helper *pHelper) {

		if helper.numIgnore == len(helper.ignore)-2 {

			if helper.keepInitial {
				r = make([]int, len(helper.nums))
				copy(r, helper.nums)
				result <- r
			}

			idx1, idx2 := findTwoNonIgnoredIdx(helper)
			helper.nums[idx1], helper.nums[idx2] =
				helper.nums[idx2], helper.nums[idx1]

			r = make([]int, len(helper.nums))
			copy(r, helper.nums)
			result <- r

			helper.nums[idx2], helper.nums[idx1] =
				helper.nums[idx1], helper.nums[idx2]
			return
		}

		for j := 1; j <= len(helper.nums); j++ {
			helper.keepInitial = true
			for i := 0; i < len(helper.ignore)-2; i++ {
				if helper.ignore[i] == 0 {
					helper.ignore[i] = 1
					helper.numIgnore++
					p(helper)
					helper.ignore[i] = 0
					helper.numIgnore--
					helper.keepInitial = false
				}
			}

			if j > 1 {
				helper.nums[j-1], helper.nums[0] = helper.nums[0], helper.nums[j-1]
			}

			if j < len(helper.nums) {
				helper.nums[0], helper.nums[j] = helper.nums[j], helper.nums[0]
			}
		}

		close(result)
	}

	return result, p
}

func main() {

	// Permutations of 0, 1, 2:
	// 0 1 2
	// 0 2 1
	// 1 0 2
	// 1 2 0
	// 2 0 1
	// 2 1 0
	var digits = []int{0, 1, 2}

	res, p := permuteFunc(digits)

	helper := &pHelper{
		nums:   digits,
		ignore: make([]int, len(digits)),
	}

	go p(helper)

	for v := range res {

		fmt.Printf("%v\n", v)
	}
}
