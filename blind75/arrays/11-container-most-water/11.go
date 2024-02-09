package main

import "fmt"


func max(x,y int)int{
	if x >= y {
		return x
	}
	return y
}

func min(x,y int)int{
	if x <= y {
		return x
	}
	return y
}

func maxArea(heights []int) int {
    currentArea := 0
	maxArea := 0
	n := len(heights)

	// start ptrs at ends
	i, j := 0, n - 1
	for i < j {
		minHeight :=  min(heights[i], heights[j])
		currentArea = (j-i) * minHeight
		maxArea = max(maxArea, currentArea)

		// shift lowest ptr until a bigger height is found
		if heights[j] <= heights[i] {
			j--
			for i < j && heights[j] <= minHeight{
				j--
			}
		} else {
			i++
			for i < j && heights[i] <= minHeight {
				i++
			}
		}

	}

	return maxArea
}


func main() {
	nums := []int{1,8,6,2,5,4,8,3,7}
	ans := maxArea(nums)
	fmt.Printf("Answer: %v", ans)
}