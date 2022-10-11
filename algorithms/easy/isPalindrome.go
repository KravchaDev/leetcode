package	main

func isPalindrome(x int) bool {
	// записываем изначально переданное число в другую переменную, для дальнейших действий с исходным X
    initial := x

	// ноль всегда палиндром
    if x == 0 {
        return true
    }
    
	// числа меньше нуля всегда не палиндром
    if x < 0 {
        return false
    }
	
    // создаем переменную для конечного числа
	finite := 0
	for ; x != 0 ; x/=10 {
		last := x % 10
		finite = finite*10 + last
	}
	return(initial==finite)
}

func main(){
	println("7554 is palindrome?  - ", isPalindrome(7554))
	println("0    is palindrome?  - ", isPalindrome(0))
	println("-121 is palindrome?  - ", isPalindrome(-121))
	println("121 is palindrome?   - ", isPalindrome(121))
	println("4224 is palindrome?  - ", isPalindrome(4224))
}