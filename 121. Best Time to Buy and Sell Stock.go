// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 121. Best Time to Buy and Sell Stock
// Letcode link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Level: easy
// Topics: Array, Dynamic Programing
// Problem detail:
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//  You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

package main

import "fmt"

// O(N^2): Time Limit Exceeded
func maxProfitBrustforce(prices []int) int {
	profit := 0
	for buy := 0; buy < len(prices)-1; buy++ {
		fmt.Printf("Buy from: %d:%d\n", buy, prices[buy])
		for sell := buy + 1; sell < len(prices); sell++ {
			subProfit := prices[sell] - prices[buy]
			if subProfit > profit {
				profit = subProfit
			}
		}
		fmt.Printf(" Profit if Buy from: %d:%d->%d\n", buy, prices[buy], profit)
	}
	return profit
}

// Neetcode
func maxProfitTwopointersFail(prices []int) int {
	profit := 0
	buy := 0
	sell := buy + 1
	for buy < len(prices)-1 && buy < sell && sell < len(prices) {
		fmt.Printf("Buy:%d Sell:%d\n", buy, sell)
		subProfit := prices[sell] - prices[buy]
		if subProfit <= 0 {
			buy++
			sell++
		} else {
			if subProfit > profit {
				profit = subProfit
			}
			if sell < len(prices)-1 {
				sell++
			} else {
				buy++
				sell = buy + 1
			}
		}
		fmt.Printf(" After Buy:%d Sell:%d profit:%d->%d\n", buy, sell, subProfit, profit)
	}
	return profit
}

// Neetcode
// Runtime: 145 ms, faster than 75.05% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.9 MB, less than 62.29% of Go online submissions for Best Time to Buy and Sell Stock.
func maxProfitTwopointers(prices []int) int {
	profit := 0
	buy := 0
	for sell := buy + 1; sell < len(prices); sell++ {
		fmt.Printf("Buy:%d Sell:%d\n", buy, sell)
		subProfit := prices[sell] - prices[buy]
		if subProfit < 0 {
			buy = sell
		} else {
			if subProfit > profit {
				profit = subProfit
			}
		}
		fmt.Printf(" profit:%d->%d\n", subProfit, profit)

	}

	return profit
}

func mainMaxProfit() {
	prices := []int{7, 1, 5, 3, 6, 4}

	prices = []int{7, 6, 4, 3, 1}
	prices = []int{2, 4, 1}

	prices = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}

	fmt.Printf("Prices:%v\n", prices)
	fmt.Printf("Max profit:%d\n", maxProfitTwopointers(prices))

}
