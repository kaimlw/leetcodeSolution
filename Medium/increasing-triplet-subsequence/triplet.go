package main

import (
	"fmt"
)

func increaseTriplet(nums []int) bool {
	min_num := 0
	min2_num := 0
	seq_flag := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			min_num = nums[i]
			continue
		}

		if nums[i] > min2_num && seq_flag == 1 {
			return true
		}

		if nums[i] < min_num {
			min_num = nums[i]
		} else if nums[i] > min_num {
			min2_num = nums[i]
			seq_flag =1
		} 
	}

	return false
}

func main() {
	testcases := [][]int{
		{1, 2, 3, 4, 5},
		{5,4,3,2,1},
		{2,1,5,0,4,6},
		{20,100,10,12,5,13},
		{0,4,2,1,0,-1,-3},
		{1,5,0,4,1,3},
	}

	for _, test := range testcases {
		fmt.Println(increaseTriplet(test))
	}
}