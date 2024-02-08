package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	pair := 0

	// for each entry, check if it's pair is in the map, 
	// or else add it to the map and continue
	for i := range nums {
		pair = target - nums[i]
		if pairIndex, ok := hashMap[pair]; ok{
			// if found:
			return []int{pairIndex, i}
		}
		hashMap[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2,11,7,15}
	target := 9
	ans := twoSum(nums, target)
	fmt.Printf("Answer: %v", ans)
}
