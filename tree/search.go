// Breadth-first and depth-first search with
// in-order, pre-order, and post-order variations.
package main

import "fmt"

type value struct {
	v int
}

type node struct {
	v        *value
	children []*node
}

func newNode(v int) *node {
	return &node{
		v:        &value{v: v},
		children: make([]*node, 0),
	}
}

func (n *node) add(child *node) {
	n.children = append(n.children, child)
}

type nlElem struct {
	next *nlElem
	prev *nlElem
	n    *node
}

type nodesList struct {
	head *nlElem
	tail *nlElem
}

func newNodesList() *nodesList {
	return &nodesList{nil, nil}
}

func (nl *nodesList) add(n *node) {

	if nl.head == nil {
		nl.head = &nlElem{
			next: nil,
			prev: nil,
			n:    n,
		}
		nl.tail = nl.head
		return
	}

	tail := nl.tail
	nl.tail = &nlElem{
		next: nil,
		prev: tail,
		n:    n,
	}
	tail.next = nl.tail
}

func (nl *nodesList) getHead() (*nlElem, error) {
	if nl.head == nil {
		return nil, fmt.Errorf("No head element")
	}

	if nl.head == nl.tail {
		head := nl.head
		nl.head = nil
		nl.tail = nil
		return head, nil
	}

	head := nl.head
	nl.head = nl.head.next
	return head, nil
}

func (nl *nodesList) notEmpty() bool {
	if nl.head == nil {
		return false
	}
	return true
}

func bfs(n *node, f func(*node)) error {

	nList := newNodesList()
	nList.add(n)

	for nList.notEmpty() {
		nlE, err := nList.getHead()
		if err != nil {
			return fmt.Errorf("Could not get head nodelist element: %w", err)
		}

		if (nlE.n != nil) && (nlE.n.children != nil) {
			for _, child := range nlE.n.children {
				nList.add(child)
			}
			f(nlE.n)
		}
	}

	return nil
}

func dfs(n *node, f func(*node)) error {

	nList := newNodesList()
	nList.add(n)

	for nList.notEmpty() {
		nlE, err := nList.getHead()
		if err != nil {
			return fmt.Errorf("Could not get head nodelist element: %w", err)
		}

		f(nlE.n)
		for _, child := range nlE.n.children {
			dfs(child, f)
		}
	}

	return nil
}

func tree(edges [][]int) *node {

	created := make(map[int]*node)
	var topNode *node
	for _, edge := range edges {
		start, end := edge[0], edge[1]
		nStart, ok := created[start]
		if !ok {
			nStart = newNode(start)
			created[start] = nStart
		}
		if topNode == nil {
			topNode = nStart
		}

		nEnd, ok := created[end]
		if !ok {
			nEnd = newNode(end)
			created[end] = nEnd
		}

		nStart.add(nEnd)
	}

	return topNode
}

func main() {

	//      1
	//      |
	//      2
	//      |\
	//      3 5
	//      |  \
	//      4   6
	edges := [][]int{
		{1, 2}, {3, 4}, {5, 6},
		{2, 3}, {2, 5},
	}

	top := tree(edges)
	err := bfs(top, func(n *node) {
		fmt.Printf("%d ", n.v.v)
	})
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}

	fmt.Printf("\n\n")

	dfs(top, func(n *node) {
		fmt.Printf("%d ", n.v.v)
	})
	fmt.Printf("\n\n")
}
