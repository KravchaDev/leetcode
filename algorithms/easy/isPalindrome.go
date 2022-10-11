func isPalindrome(x int) bool {
	if math.Signbit(float64(x)) {
		return false
	}
	rev := 0
	remainder := 0
	argz := x
	for argz > 0 {
		rev *= 10
		remainder = argz % 10
		rev += remainder
		argz /= 10
	}
	if rev == x {
		return true
	}
	return false
