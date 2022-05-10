package main

import (
	"fmt"
)

type characteristics struct {
	repeating []string
	numLower  int
	numUpper  int
	length    int
}

func updateRepeating(p *characteristics, repeating string, numRepeats int) {
	tempStr := ""
	for i := 0; i < numRepeats+1; i++ {
		tempStr += repeating
	}
	p.repeating = append(p.repeating, tempStr)
}

func strongPasswordChecker(password string) int {

	minLength := 6
	maxLength := 20
	p := new(characteristics)
	p.repeating = make([]string, 0)
	p.length = len(password)

	repeating := ""
	numRepeats := 0
	for i, c := range password {
		if i > 0 {
			if string(password[i]) == string(password[i-1]) {
				if repeating == "" {
					repeating = string(c)
				}

				numRepeats++
				fmt.Printf("numRepeats 1: %d i: %d password[i-1]: %s password[i]: %s\n",
					numRepeats, i, string(password[i-1]), string(password[i]))
			} else if repeating != "" && numRepeats > 1 {
				fmt.Printf("numRepeats: %d\n", numRepeats)
				updateRepeating(p, repeating, numRepeats)
				repeating = ""
				numRepeats = 0
			} else if repeating != "" {
				repeating = ""
				numRepeats = 0
			}
		}
		if int(c) >= int('A') && int(c) <= int('Z') {
			p.numUpper++
		}

		if int(c) >= int('a') && int(c) <= int('z') {
			p.numLower++
		}
	}

	if repeating != "" && numRepeats > 1 {
		fmt.Printf("numRepeats 2: %d\n", numRepeats)
		updateRepeating(p, repeating, numRepeats)
		repeating = ""
		numRepeats = 0
	}

	fmt.Printf("p: %v\n", p)

	numChanges := 0
	lenIncreases := 0
	/*
	   lenDecreases := 0

	   if p.length > maxLength {
	       numChanges += p.length - maxLength
	       lenDecreases = p.length - maxLength
	   }
	*/

	if p.length < minLength {
		numChanges += minLength - p.length
		lenIncreases = minLength - p.length
	}

	if p.numLower == 0 {
		if maxLength-p.length > 0 {
			// can add lowercase value
		} else {
			// must update one value to lowercase
		}
		if lenIncreases == 0 {
			numChanges++
		}
	}

	if p.numUpper == 0 {
		if maxLength-p.length > 0 {
			// can add uppercase value
		} else {
			// must update one value to uppercase
		}
		if lenIncreases-numChanges > 0 {
			numChanges++
		}
	}

	for _, rep := range p.repeating {
		l := len(rep)
		if l > 2 {
			if l%2 == 0 {
				numChanges += int(l/2) - 1
			} else {
				numChanges += int(l / 2)
			}
		}
	}

	fmt.Printf("lenIncreases: %d\n", lenIncreases)
	fmt.Printf("len(repeating): %d\n", len(p.repeating))
	if lenIncreases > 0 && (len(p.repeating) > 0) {
		numChanges--
	}

	return numChanges
}

func main() {
	s := "aaAA11"
	fmt.Printf("numChanges: %d\n", strongPasswordChecker(s))
}
