package main

import "fmt"

func search(nums []int, target int) int {
	notFound := -1
    n := len(nums)
    l,m,r := 0,0,n-1
	for l <= r {
        // first calculate & check new midpoint
        m = (l + r) / 2
        if nums[m] == target {
            return m
        }
        // if inside left sorted portion
        if nums[l] <= nums[m] {
            // if target outside range
            if target > nums[m] || target < nums[l] {
                // then only consider right half of slice
                l = m + 1
            } else {
                // only consider left half
                r = m - 1
            }
        // elif inside right sorted portion
        } else {
            if target < nums[m] || target > nums[r] {
                r = m - 1
            } else {
                l = m + 1
            }
        }
    }
	return notFound
}


func main() {
	nums := []int{4,5,6,7,0,1,2}
	ans := search(nums, 0)
	fmt.Printf("Answer: %v", ans)
}