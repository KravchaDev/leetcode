package main

func numIdenticalPairs(nums []int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs += 1
			}
		}
	}
	return (pairs)
}

func main() {
	nums := [6]int{1, 2, 3, 1, 1, 3}
	println(numIdenticalPairs(nums[:]))
}
