package main

import "fmt"

func max(x,y int)int{
	if x > y{
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	// set initial max to first entry
	maxSum := nums[0]
	// init currentSum var in memory
	currentSum := 0

	for i := range nums {
		// add next entry
        currentSum += nums[i]
		// check max
		maxSum = max(maxSum, currentSum)
		if currentSum < 0 {
			// then restart subarray sum from next entry
			currentSum = 0
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	ans := maxSubArray(nums)
	fmt.Printf("Answer: %v", ans)
}