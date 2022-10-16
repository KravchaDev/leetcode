package main

import (
	"fmt"
)

func smallestEvenMultiple(n int) int {
	switch n % 2 {
	// в случае если оно четное
	case 0:
		return n
	// в любых других случаях оно нечетное
	default:
		return n * 2
	}
}

func main() {
	for i := 0; i < 10; i++ {
		x := smallestEvenMultiple(i)
		fmt.Printf("%d == %d\n", i, x)
	}
}
