package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	i, j, k := 0, 1, 2
	base := 0
	sum := 0

	// sort list in ascending order
	sort.Ints(nums)

	// for each seed (up to the third to last entry)
	for i = 0; i < len(nums)-2; i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// search for solution given a fixed base val
		base = nums[i]

		// set left and right pointers
		j, k = i+1, len(nums)-1
		
		// converge ptrs based on sum
		for j < k {
			// get vals
			sum = base + nums[j] + nums[k]

			switch {
			case sum > 0:
				// move nums[k] pointer down
				k--
			case sum < 0:
				// move nums[j] pointer up
				j++
			case sum == 0:
				ans = append(ans, []int{base, nums[j], nums[k]})
				j++
				// skip duplicates after answer found
				for j < k && nums[j] == nums[j-1]{
					j++
				}
			}
		}
	}
	return ans
}



func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	ans := threeSum(nums)
	fmt.Printf("Answer: %v", ans)
}