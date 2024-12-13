package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	bedIndex := []int{}
	flag := false
	for i, item := range flowerbed{
		if flag {
			flag = false
			continue
		}

		if item == 1 {
			flag = true
			continue			
		}

		if item == 0 {
			if i < len(flowerbed)-1 {
				if flowerbed[i+1] == 1 {
					continue
				}
			}
			bedIndex = append(bedIndex, i)
			flag = true
		}
	}

	return len(bedIndex) > 0 && n <= len(bedIndex)
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		prev := 0
		next := 0
		current := flowerbed[i]

		if i > 0 {
			prev = flowerbed[i-1]
		}

		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}

		if prev == 0 && next == 0 && current == 0 {
			if n <= 0 {
				break
			}

			n--
			i++
		}
	}

	return n <= 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,1,0,0}, 2))
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,1,0,1}, 1))
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,0,1}, 2))
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,0,1}, 2))
	// fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 1))
	fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 2))
}