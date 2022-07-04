package main

import "fmt"

func binSearch(xs []int, startIdx int) int {

	if len(xs) == 1 { // Only one element in array
		return -1
	}

	if xs[0] < xs[len(xs)-1] { // Array is not rotated
		return -1
	}

	i := len(xs) / 2
	if xs[i] < xs[i-1] {
		return startIdx + i
	}

	if len(xs) > 1 {
		idxL := binSearch(xs[:i], 0)
		idxR := binSearch(xs[i:], i)
		if idxL == -1 && idxR == -1 {
			return -1
		}
		if idxL == -1 {
			return idxR + startIdx
		} else {
			return idxL + startIdx
		}
	} else {
		return -1
	}
}

func targetIdx(xs []int, startIdx, target int) int {

	fmt.Printf("xs: %+v, startIdx: %d, target: %d\n", xs, startIdx, target)

	if len(xs) == 0 {
		return -1
	}

	if len(xs) == 1 {
		if xs[0] == target {
			return startIdx
		} else {
			return -1
		}
	}

	idxL := targetIdx(xs[:len(xs)/2], 0, target)
	idxR := targetIdx(xs[len(xs)/2:], len(xs)/2, target)

	if idxL == -1 && idxR == -1 {
		return -1
	}
	if idxL == -1 {
		return idxR + startIdx
	} else {
		return idxL + startIdx
	}
}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}

	if nums[0] < nums[len(nums)-1] { // Array is not rotated
		return targetIdx(nums, 0, target)
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	} else if len(nums) == 1 {
		return -1
	}

	minIdx := binSearch(nums, 0)

	fmt.Printf("minIdx: %d\n", minIdx)

	if target >= nums[0] {
		return targetIdx(nums[0:minIdx], 0, target)
	} else {
		return targetIdx(nums[minIdx:], minIdx, target)
	}
}

func main() {
	xs := []int{3, 5, 1}
	target := 3

	fmt.Printf("targetIdx: %d\n", search(xs, target))
}
