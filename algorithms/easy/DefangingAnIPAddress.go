package main

import (
	"fmt"
	"strings"
)

func main() {
	refString := "255.100.50.0"
	fmt.Println(strings.Replace(refString, ".", "[.]", -1))
}
