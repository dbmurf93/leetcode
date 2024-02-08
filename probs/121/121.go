package main

import (
	"fmt"
)

func max(x,y int)int{
	if x >= y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	maxProfit := 0
	buyPrice := prices[0]
    for _, price := range prices[1:]{
		if price < buyPrice {
			buyPrice = price
			continue
		} else {
			maxProfit = max(price-buyPrice, maxProfit)
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7,1,5,3,6,4}
	ans := maxProfit(prices)
	fmt.Printf("Answer: %v", ans)
}
