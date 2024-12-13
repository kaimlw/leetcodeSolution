package main

import "slices"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	greatest := slices.Max(candies)

	for i, item := range candies {
		if item + extraCandies >= greatest  {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func main() {
	kidsWithCandies([]int{1,2,3,5}, 3)
}