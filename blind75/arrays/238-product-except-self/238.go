package main

import "fmt"

func productExceptSelf(nums []int) []int {
    fwd := make([]int, len(nums))
	bkwd := make([]int, len(nums))
	ans := make([]int, len(nums))

	// seed snowball arrays
	fwd[0] = 1
	bkwd[len(nums)-1] = 1

	// multiply each number by its predecessor (beginning to end)
	for i := 1; i < len(fwd); i++ {
		fwd[i] = nums[i-1] * fwd[i - 1]
	}

	// multiply each number by those coming after it (end to beginning)
	for i := len(bkwd) - 2; i >= 0; i-- {
		bkwd[i] = nums[i + 1] * bkwd[i + 1]
	}

	// combine results for answer
	for i := range nums {
		ans[i] = fwd[i] * bkwd[i]
	}
	return ans
}

func main() {
	nums := []int{1,2,3,4}
	ans := productExceptSelf(nums)
	fmt.Printf("Answer: %v", ans)
}