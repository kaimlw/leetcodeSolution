package main

import "fmt"

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	for i := 0; i < len(nums)-1; i++ {
		if i == 0 {
			left[i] = 1
			right[len(nums)-1] = 1
		}
		left[i+1] = left[i] * nums[i]
		right[len(nums)-2-i] = right[len(nums)-1-i] * nums[len(nums)-1-i]
	}
	
	for i := 0; i < len(nums); i++ {
		left[i] = left[i] * right[i]
	}
	
	return left
}

func main() {
	input := []int{1,2,3,4}
	// input := []int{-1,1,0,-3,3}
	fmt.Println(productExceptSelf(input))
}