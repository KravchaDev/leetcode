package main

func defangIPaddr(address string) string {
	var output string
	for i := 0; i < len(address); i++ {
		if string(address[i]) == "." {
			output += "[.]"
		} else {
			output += string(address[i])
		}
	}
	println(output)
	return output
}

func main() {
	defangIPaddr("255.100.50.0")
	defangIPaddr("95.248.195.127")
	defangIPaddr("39.27.141.110")
}
