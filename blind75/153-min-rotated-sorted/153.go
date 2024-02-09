package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	ans := nums[0]

	// init ptrs to ends
	leftPrev := nums[0]
	rightPrev := nums[n-1]
	

    for i := 1; i < n; i++ {
		// if descending found from left
        if nums[i] < leftPrev {
			ans = nums[i]
			break
		}
		// if ascending found from right
		if nums[n-1-i] > rightPrev{
			ans = rightPrev
			break
		}

		// collapse pointers inwards, ignore overlap bc of breaks
		leftPrev = nums[i]
		rightPrev = nums[n-1-i]
    }
	return ans
}


func main() {
	nums := []int{4,5,6,7,0,1,2}
	ans := findMin(nums)
	fmt.Printf("Answer: %v", ans)
}