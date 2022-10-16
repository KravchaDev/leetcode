package main

func outArray(nums []int) {
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}

func main() {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	runningSum(nums[:])
	outArray(nums[:])
}
