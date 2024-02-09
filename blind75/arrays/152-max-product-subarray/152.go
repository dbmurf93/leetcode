package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProduct(nums []int) int {
	// init vars
	numsLen := len(nums)
	maxProduct := nums[0]
	leftProduct := 1
	rightProduct := 1

	for i := 0; i < numsLen; i++ {
		// reset after any zeroes, & continue
		if leftProduct == 0 {
			leftProduct = 1
		}
		if rightProduct == 0 {
			rightProduct = 1
		}

		// trace array both directions
		leftProduct = leftProduct * nums[i]
		rightProduct = rightProduct * nums[numsLen-1-i]

		maxProduct = max(maxProduct, max(leftProduct, rightProduct))
	}
	return maxProduct
}

func main() {
	nums := []int{2, 3, -2, 4}
	ans := maxProduct(nums)
	fmt.Printf("Answer: %v", ans)
}