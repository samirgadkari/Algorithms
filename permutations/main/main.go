package main

import "fmt"

// 1. find the index of the decreasing value from the right (index of 3)
// 2. find the index with the value slightly larger than it from the decreasing value onwards (index of 4)
// 3. swap those values
// 4. reverse the right hand string
func permute(digits []uint8, res [][]uint8) ([][]uint8, error) {

	for {
		decIdx := -1
		for i := len(digits); i >= 1; i-- {
			if digits[i-1] < digits[i] {
				decIdx = i - 1
				break
			}
		}
		if decIdx == -1 {
			return res, nil
		}

		nextBiggestIdx := -1
		for i, digit := range digits[decIdx+1:] {
			if digit > digits[decIdx] &&
				digit < digits[nextBiggestIdx] {
				nextBiggestIdx = i
			}
		}
		if nextBiggestIdx == -1 {
			return nil, fmt.Errorf("Error during permute")
		}

		digits[nextBiggestIdx], digits[decIdx] = digits[decIdx], digits[nextBiggestIdx]
		for left, right := decIdx+1, len(digits)-1; left < right; left, right = left+1, right-1 {
			digits[left], digits[right] = digits[right], digits[left]
		}

		res = append(res, digits)
	}

	return nil, fmt.Errorf("Should not reach here")
}

func main() {

	var digits = []uint8{0, 1, 2, 3, 4}
	var res [][]uint8

	res, err := permute(digits, res)
	if err != nil {
		fmt.Printf("Error permuting %v: %s\n", digits, err.Error())
		return
	}

	fmt.Printf("Result: %v\n", res[:10])
}
