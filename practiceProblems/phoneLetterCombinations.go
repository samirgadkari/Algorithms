package main

import "fmt"

type rFn func([]string, []int, int, string)

func recurseFn() (rFn, chan string) {

	result := make(chan string, 0)
	var recurse rFn

	recurse = func(s []string, lengths []int, level int, partial string) {

		for _, c := range s[level] {

			if level == len(s)-1 {
				s2 := partial + string(c)
				result <- s2
			} else if level < len(s)-1 {

				recurse(s, lengths, level+1, partial+string(c))
			}
		}

		if level == 0 {
			close(result)
		}
	}

	return recurse, result
}

func lc(digits string) chan string {

	m := make(map[int]string, 8)
	m[2] = "abc"
	m[3] = "def"
	m[4] = "ghi"
	m[5] = "jkl"
	m[6] = "mno"
	m[7] = "pqrs"
	m[8] = "tuv"
	m[9] = "wxyz"

	lengths := make([]int, 0)
	s := make([]string, 0)
	for _, i := range digits {

		chars := m[int(byte(i)-byte('0'))]
		lengths = append(lengths, len(chars))
		s = append(s, chars)
	}

	recurse, result := recurseFn()
	go recurse(s, lengths, 0, "")

	return result
}

func letterCombinations(digits string) []string {

	result := lc(digits)
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	for s := range result {
		res = append(res, s)
	}

	return res
}

func main() {

	number := "2859"

	result := letterCombinations(number)
	for _, s := range result {
		fmt.Printf("%s\n", s)
	}
}
