package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testInOut struct {
	edges    [][]int
	bfsNodes []int
	dfsNodes []int
	topNode  *node
}

type SearchTreeSuite struct {
	suite.Suite
	tests []testInOut
}

func newSearchTreeSuite() *SearchTreeSuite {

	s := new(SearchTreeSuite)
	s.tests = []testInOut{
		{
			// simple tree
			//      1
			//      |
			//      2
			//      |\
			//      3 5
			//      |  \
			//      4   6
			edges: [][]int{{1, 2}, {3, 4}, {5, 6},
				{2, 3}, {2, 5}},
			bfsNodes: []int{1, 2, 3, 5, 4, 6},
			dfsNodes: []int{1, 2, 3, 4, 5, 6},
			topNode:  nil,
		},
		{
			// simple tree
			//      1 - 2 - 3
			//      |       |
			//      6 - 5 - 4

			edges: [][]int{{1, 2}, {3, 4}, {5, 6},
				{2, 3}, {4, 5}, {6, 1}},
			bfsNodes: []int{1, 2, 3, 4, 5, 6},
			dfsNodes: []int{1, 2, 3, 4, 5, 6},
			topNode:  nil,
		},
		{
			// 2-element tree
			//      1 - 2
			edges:    [][]int{{1, 2}},
			bfsNodes: []int{1, 2},
			dfsNodes: []int{1, 2},
			topNode:  nil,
		},
	}

	for i, test := range s.tests {
		s.tests[i].topNode = tree(test.edges)
	}

	return s
}

func (s *SearchTreeSuite) TestSearchTree() {
	r := s.Require()
	for _, test := range s.tests {

		result := make([]int, 0)

		err := bfs(test.topNode, func(n *node) {
			result = append(result, n.v.v)
		})
		r.Equal(err, nil)

		r.Equal(test.bfsNodes, result)
	}
}

func TestSearchTreeSuite(t *testing.T) {

	s := newSearchTreeSuite()
	suite.Run(t, s)
}
