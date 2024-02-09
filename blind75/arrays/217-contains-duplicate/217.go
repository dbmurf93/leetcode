package main

import "fmt"

func containsDuplicate(nums []int) bool {
    hashMap := make(map[int]int)
    for _, num := range nums {
        if _, ok := hashMap[num]; ok{
            return true
        } else {
            hashMap[num] = 0
        }
    }
    return false
}

func main() {
	nums := []int{2,11,7,15, 15}
	ans := containsDuplicate(nums)
	fmt.Printf("Answer: %v", ans)
}
