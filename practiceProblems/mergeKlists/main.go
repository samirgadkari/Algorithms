package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertToChs(lists []*ListNode) []chan *ListNode {

	resChs := make([]chan *ListNode, 0)

	chNum := 0

	for _, node := range lists {

		ch := make(chan *ListNode, 0)
		resChs = append(resChs, ch)

		go func(head *ListNode, cNum int) {

			defer close(ch)

			n := head

			for n != nil {
				// fmt.Printf("Putting n {%+v} in channel: %d\n", n, cNum)
				ch <- n
				n = n.Next
			}
		}(node, chNum)

		chNum++
	}

	return resChs
}

func merge2Chs(ch1 chan *ListNode, ch2 chan *ListNode) chan *ListNode {

	resCh := make(chan *ListNode, 0)

	var v1 *ListNode
	var v2 *ListNode
	var saveV1 *ListNode
	var saveV2 *ListNode

	go func() {
		defer close(resCh)
	LOOP:
		for {
			// fmt.Printf("Getting values from channel\n")
			if saveV1 == nil {
				v1 = <-ch1
			} else {
				v1 = saveV1
			}
			if saveV2 == nil {
				v2 = <-ch2
			} else {
				v2 = saveV2
			}

			// fmt.Printf("Got values: %+v %+v from channels: %+v %+v\n",
			//	v1, v2, ch1, ch2)

			switch {
			case v1 == nil && v2 == nil:
				// fmt.Printf("Both channels have nil values in them\n")
				break LOOP
			case v1 == nil:
				resCh <- v2
				// fmt.Printf("Sent v2: %+v\n", v2)
				saveV1 = v1
				saveV2 = nil
			case v2 == nil:
				resCh <- v1
				// fmt.Printf("Sent v1: %+v\n", v1)
				saveV1 = nil
				saveV2 = v2
			case v1.Val <= v2.Val:
				resCh <- v1
				// fmt.Printf("2 Sent v1: %+v\n", v1)
				saveV1 = nil
				saveV2 = v2
			default:
				resCh <- v2
				// fmt.Printf("2 Sent v2: %+v\n", v2)
				saveV1 = v1
				saveV2 = nil
			}
		}
	}()

	return resCh
}

func mergeKChs(chs []chan *ListNode) []chan *ListNode {

	l := len(chs)
	resChs := make([]chan *ListNode, 0)

	for i := 0; i < l; i += 2 {

		if i+1 >= l {
			break
		}

		resChs = append(resChs, merge2Chs(chs[i], chs[i+1]))
		// fmt.Printf("Merged 2 channels: resChs: %+v\n", resChs)
	}

	if l%2 == 1 {
		ch := merge2Chs(chs[l-1], resChs[len(resChs)-1])
		// fmt.Printf("len(resChs): %d\n", len(resChs))
		resChs = append(resChs[:len(resChs)-1], ch)
		// fmt.Printf("Merged 2 channels: resChs: %+v\n", resChs)
	}

	// fmt.Printf("returning resChs: %+v\n", resChs)
	return resChs
}

func convertToList(ch chan *ListNode) *ListNode {

	// fmt.Printf("converting to list\n")

	var last *ListNode
	var head *ListNode

	for p := range ch {
		// fmt.Printf("p.Val: %d\n", p.Val)
		if head == nil {
			head = p
		}
		if last != nil {
			last.Next = p
		}
		last = p
	}

	return head
}

func mergeKLists(lists []*ListNode) *ListNode {

	chs := convertToChs(lists)

	for len(chs) > 1 {
		chs = mergeKChs(chs)
		// fmt.Printf("len(chs): %d\n", len(chs))
	}

	return convertToList(chs[0])
}

func main() {

	d := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}

	allData := make([]*ListNode, len(d))
	for i, s := range d {
		data := make([]*ListNode, len(s))
		for j, v := range s {
			data[j] = &ListNode{Val: v}
			if j > 0 {
				data[j-1].Next = data[j]
			}
		}

		allData[i] = data[0]
	}

	fmt.Printf("allData: %+v\n", allData)
	result := mergeKLists(allData)
	fmt.Printf("[[1,4,5],[1,3,4],[2,6]]: %v\n", result)

	r := result
	for r != nil {
		fmt.Printf("r.Val: %d\n", r.Val)
		r = r.Next
	}
}

/* Leetcode has problems with this code, however the go compiler does not.
*  Leetcode prints:
*  Line 167: Char 19: undefined: Deserializer (solution.go)
*  Line 168: Char 33: undefined: Deserializer (solution.go)
*  Line 200: Char 34: undefined: Serializer (solution.go)
 */
