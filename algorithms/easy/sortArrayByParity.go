package main

func sortArrayByParity(nums []int) []int {
	evenNum := 0
    for i := range nums {
		if nums[i]%2 == 0{
			nums[i], nums[evenNum] =  nums[evenNum], nums[i]
			evenNum++
		}
	}
	return nums[:]
}

func inputArray(nums []int) {
	for i := 0; i < len(nums); i++{
		println(nums[i])
	}
}

func main(){
	nums := [4]int{3,1,2,4}
	inputArray(sortArrayByParity(nums[:]))
}