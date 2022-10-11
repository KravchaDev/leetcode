func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	r := 0
	for z := x; z > 0; z /= 10 {
		 tmp := z % 10
		 r *= 10
		 r += tmp

	}
	return r == x
}
