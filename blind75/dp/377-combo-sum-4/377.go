package main

import (
	"fmt"
	"slices"
)

func combinationSum4(nums []int, target int) int {
	slices.Sort(nums)
	// target : # of ways
	cacheMap := make(map[int]int)
	// base case
	cacheMap[0] = 1

	var dp func(target int)int
	dp = func(target int) int {
		// check cache first
		if val, ok := cacheMap[target]; ok {
			return val
		}

		// reset result
		res := 0

		// try each num in array
		for _, num := range nums {
			if target < num {
				break
			}
			// call recursively for remaining target
			res += dp(target - num)
		}
		// cache and return result
		cacheMap[target] = res
		return res
	}
	return dp(target)
}


func main() {
	nums := []int{1,2,3}
	ans := combinationSum4(nums, 4)
	fmt.Printf("Answer: %v", ans)
}