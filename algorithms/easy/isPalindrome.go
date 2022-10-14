package main

func isPalindrome(x int) bool {
	// записываем изначально переданное число в другую переменную, для дальнейших действий с исходным X
	initial := x

	switch {
	// ноль всегда палиндром
	case x == 0:
		return true
	// число меньше нуля всегда не палиндром
	case x < 0:
		return false
	default:
		// создаем переменную для конечного числа
		finite := 0
		for ; x != 0; x /= 10 {
			last := x % 10
			finite = finite*10 + last
		}
		// возвращаем true/false (bool)
		return (initial == finite)
	}
}

func main() {
	println("7554 is palindrome?  - ", isPalindrome(7554)) // false
	println("0    is palindrome?  - ", isPalindrome(0))    // true
	println("-121 is palindrome?  - ", isPalindrome(-121)) // false
	println("121 is palindrome?   - ", isPalindrome(121))  // true
	println("4224 is palindrome?  - ", isPalindrome(4224)) // true
}
