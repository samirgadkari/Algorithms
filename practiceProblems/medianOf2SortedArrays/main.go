package main

import "fmt"

func withinRange(l int, v int) int {
	if v < 0 {
		return 0
	} else if v >= l {
		return l - 1
	}

	return v
}

func median(xs []int, ys []int) float64 {

	lenX := len(xs)
	lenY := len(ys)
	lenBoth := lenX + lenY
	midX := withinRange(lenX, lenX/2-1)
	midY := withinRange(lenY, lenY/2-1)

	var midBoth int
	if lenBoth&1 == 0 { // len of merged array is even

		midBoth = lenBoth/2 - 1
	} else {
		midBoth = lenBoth / 2
	}
	midBoth = withinRange(lenBoth, midBoth)

	for midBoth != (midX + midY) {

		fmt.Printf("  midBoth: %2d midX: %2d  midY: %2d\n", midBoth, midX, midY)

		if midBoth < (midX + midY) { // Search in lower indices
			if xs[midX] > ys[midY] {
				midY = withinRange(lenY, midY/2)
			} else {
				midX = withinRange(lenX, midX/2)
			}
		} else { // search in higher indices
			if xs[midX] > ys[midY] {
				midY = withinRange(lenY, midY+(lenY-1-midY)/2)
			} else {
				midX = withinRange(lenX, midX+(lenX-1-midX)/2)
			}
		}
		fmt.Printf("2 midBoth: %2d midX: %2d  midY: %2d\n", midBoth, midX, midY)
	}

	fmt.Printf("lenBoth: %d\n", lenBoth)
	if lenBoth&1 == 0 { // Len of merged array is even

		fmt.Printf("even length\n")
		return float64((xs[midX] + ys[midY])) / 2
	} else {
		fmt.Printf("odd length\n")
		if xs[midX] > ys[midY] {
			return float64(ys[midY])
		} else {
			return float64(xs[midX])
		}
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return median(nums1, nums2)
	} else {
		return median(nums2, nums1)
	}
}

func main() {

	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 4, 6, 8, 10, 12, 14}
	// median := (6 + 7) / 2

	// TODO: Does not work on this data
	// nums1 := []int{1, 3}
	// nums2 := []int{2}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("result: %6.2f\n", result)
}
