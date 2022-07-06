package main

import "fmt"

const (
	byte0 byte = 0
	byte1 byte = 1
)

func sameSizeSliceVisited(xs [][]byte, v bool) [][]bool {

	result := make([][]bool, 0)
	l := len(xs[0])
	for i := 0; i < len(xs); i++ {
		result = append(result, make([]bool, l))
		for j := 0; j < l; j++ {
			result[i][j] = v
		}
	}

	return result
}

func topLeft(i, j byte, board [][]byte) bool {
	if i == byte0 && j == byte0 {
		return true
	}
	return false
}

func topRight(i, j byte, board [][]byte) bool {
	if i == byte0 && int(j) == len(board[0])-1 {
		return true
	}
	return false
}

func bottomLeft(i, j byte, board [][]byte) bool {
	if int(i) == len(board)-1 && j == byte0 {
		return true
	}
	return false
}

func bottomRight(i, j byte, board [][]byte) bool {
	if int(i) == len(board[0])-1 && int(j) == len(board[0])-1 {
		return true
	}
	return false
}

func left(i, j byte, board [][]byte) bool {
	if j == byte0 {
		return true
	}
	return false
}

func right(i, j byte, board [][]byte) bool {
	if int(j) == len(board[0])-1 {
		return true
	}
	return false
}

func top(i, j byte, board [][]byte) bool {
	if i == byte0 {
		return true
	}
	return false
}

func bottom(i, j byte, board [][]byte) bool {
	if int(i) == len(board)-1 {
		return true
	}
	return false
}

func search(word string, board [][]byte, i, j byte, visited [][]bool, m []byte) bool {

	if len(word) == 0 {
		return false
	}

	fmt.Printf("word: %s, i: %d, j: %d, visited: %+v, m: %+v\n",
		word, i, j, visited, m)

	if word[0] == board[i][j] {

		if len(word) == 1 {
			return true
		}

		visited[i][j] = true
		m = append(m, board[i][j])

		if topLeft(i, j, board) {
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j+byte1, visited, m) == true {
				return true
			}
		} else if topRight(i, j, board) {
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
		} else if bottomLeft(i, j, board) {
			if search(word[1:], board, i-byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j+byte1, visited, m) == true {
				return true
			}
		} else if bottomRight(i, j, board) {
			if search(word[1:], board, i-byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
		} else if top(i, j, board) {
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j+byte1, visited, m) == true {
				return true
			}
		} else if right(i, j, board) {
			if search(word[1:], board, i-byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
		} else if left(i, j, board) {
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j+byte1, visited, m) == true {
				return true
			}
		} else if right(i, j, board) {
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
		} else {
			if search(word[1:], board, i-byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j-byte1, visited, m) == true {
				return true
			}
			if search(word[1:], board, i+byte1, j, visited, m) == true {
				return true
			}
			if search(word[1:], board, i, j+byte1, visited, m) == true {
				return true
			}
		}

		visited[i][j] = false
		m = m[:len(m)-1]
		return false
	} else {
		return false
	}
}

func exist(board [][]byte, word string) bool {

	visited := sameSizeSliceVisited(board, false)
	m := make([]byte, 0)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(word, board, byte(i), byte(j), visited, m) == true {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"

	fmt.Printf("exist: %t\n", exist(board, word))
}
