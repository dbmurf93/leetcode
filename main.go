package main

import (
	"fmt"
	"slices"
)

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

func fourSum(nums []int, target int) [][]int {
    slices.Sort(nums)
	res := [][]int{}
	quad := []int{}

	// declare and create helper func
	var kSum func(k, start, target int)[]int
	kSum = func (k, start, target int)[]int{
		// for all recursive cases
		if k != 2 {
			// test all start numbers up to k from the end of the list
			for i := start; i < len(nums) - k; i++{
				if i > start && nums[i] == nums[i - 1]{
					// skip duplicates
					continue
				}
				// add number to quad
				quad = append(quad, nums[i])
				kSum(k-1, i+1, target - nums[i])
				// remove last element and loop
				quad = quad[:len(quad)-1]
			}
		} else {
			// base case 2Sum prob
			l, r := start, len(nums)-1

		}

		// for base case k=2
	return quad
	}
}

func main() {
	// nums := []int{11,13,15,17}
	nums := []int{2,7,11,15}
	target := 9
	ans := fourSum(nums, target)
	// fmt.Printf("Answer: %v\nTarget: %v\n", ans, target)
	fmt.Printf("Answer: %v", ans)
}
