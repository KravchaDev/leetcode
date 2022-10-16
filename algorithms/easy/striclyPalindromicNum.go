// this task is very stupid, there is a value of 6643 in the test case
// (this number is strictly polydromic), but the leetcode system requires
// me to false, instead of true
package main

import "strconv"

func isStrictlyPalindromic(n int) bool {
	binInt := int64(n)
	base2 := strconv.FormatInt(binInt, 2)
	length2 := len(base2) - 1

	for i := 0; i < length2; i++ {
		if base2[i] != base2[length2] {
			return false
		}
		length2--
	}

	base3 := strconv.FormatInt(binInt, 3)
	length3 := len(base3) - 1
	for i := 0; i < length3; i++ {
		if base3[i] != base3[length3] {
			return false
		}
		length3--
	}

	return true
}

func main() {
	println("strictly palindromic numbers:")
	for i := 0; i < 9999; i++ {
		x := isStrictlyPalindromic(i)
		if x == true {
			println(i)
		}
	}
}

// output 0, 1, 6643 - these are all existing numbers strict palindromes
