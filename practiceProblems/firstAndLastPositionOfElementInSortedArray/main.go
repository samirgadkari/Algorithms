package main

import (
	"fmt"
)

const (
	printDebugMsgs bool = true
)

type concurrency struct {
	doneChs []chan struct{}
}

func NewConcurrency() *concurrency {
	d := make([]chan struct{}, 0)
	return &concurrency{d}
}

type upperFunc func([]int, int, int, int, *concurrency)
type lowerFunc func([]int, int, int, *concurrency)

type concurrencyFunc func(params ...interface{})

func (c *concurrency) AddUpperFunc(f upperFunc, nums []int,
	offset, target, rightmostIdx int, c1 *concurrency) {
	done := make(chan struct{})
	c.doneChs = append(c.doneChs, done)
	go func() {
		defer close(done)
		f(nums, offset, target, rightmostIdx, c1)
	}()
}

func (c *concurrency) AddLowerFunc(f lowerFunc, nums []int,
	offset, target int, c1 *concurrency) {
	done := make(chan struct{})
	c.doneChs = append(c.doneChs, done)
	go func() {
		defer close(done)
		f(nums, offset, target, c1)
	}()
}

// Only when all the done channels are signalled, we will come out of this function
func (c *concurrency) And() {
	for i, doneCh := range c.doneChs {

		fmt.Printf("Waiting on done channel %d\n", i)
		select {
		case <-doneCh:
		}
	}
}

var start int = -1
var end int = -1

func lower(nums []int, offset, target int, c *concurrency) {

	if printDebugMsgs {
		fmt.Printf("offset: %d, nums: %+v\n", offset, nums)
	}

	if len(nums) < 2 {
		if nums[0] == target && offset == 0 {
			start = offset
		}
		return
	}

	i := len(nums)/2 - 1
	if printDebugMsgs {
		fmt.Printf("i: %d\n", i)
	}

	switch {
	case nums[i+1] == target && nums[i] != target:
		start = offset + i + 1

	case nums[i] == target:
		c.AddLowerFunc(lower, nums[:i+1], offset, target, c)

	case nums[i] < target:
		c.AddLowerFunc(lower, nums[i+1:], offset+i+1, target, c)

	default: // This is the value > target case.
		if i > 0 {
			c.AddLowerFunc(lower, nums[:i+1], offset, target, c)
		}
	}
}

func runLower(nums []int, offset, target int) {

	c := NewConcurrency()
	c.AddLowerFunc(lower, nums, offset, target, c)
	c.And()
}

func upper(nums []int, offset, target int, rightmostIdx int, c *concurrency) {

	if printDebugMsgs {
		fmt.Printf("offset: %d, nums: %+v\n", offset, nums)
	}

	if len(nums) < 2 {
		if nums[0] == target && offset == rightmostIdx {
			end = offset
		}
		end = -1
		return
	}

	i := len(nums)/2 - 1
	if printDebugMsgs {
		fmt.Printf("i: %d\n", i)
	}

	switch {
	case nums[i] == target && nums[i+1] != target:
		end = offset + i

	case nums[i] == target:
		c.AddUpperFunc(upper, nums[i+1:], offset+i+1, target, rightmostIdx, c)

	case nums[i] < target:
		c.AddUpperFunc(upper, nums[i+1:], offset+i+1, target, rightmostIdx, c)

	default: // This is the value > target case.

		c.AddUpperFunc(upper, nums[:i+1], offset, target, rightmostIdx, c)
	}
}

func runUpper(nums []int, offset, target int, rightmostIdx int) {

	c := NewConcurrency()
	c.AddUpperFunc(upper, nums, offset, target, rightmostIdx, c)
	c.And()
}

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	runLower(nums, 0, target)

	if printDebugMsgs {
		fmt.Printf("start: %d\n", start)
	}

	runUpper(nums, 0, target, len(nums)-1)
	if printDebugMsgs {
		fmt.Printf("end: %d\n", end)
	}

	if printDebugMsgs {
		fmt.Printf("start: %d, end: %d\n", start, end)
	}
	return []int{start, end}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Printf("nums: %v, target: %d, result: %v\n",
		nums, target, searchRange(nums, target))

	/*
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 6
		fmt.Printf("nums: %v, target: %d, result: %v\n",
			nums, target, searchRange(nums, target))
	*/
}
