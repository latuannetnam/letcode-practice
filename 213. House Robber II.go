// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 213. House Robber II
// Letcode link: https://leetcode.com/problems/house-robber-ii/
// Level: medium
// Topics: Array, Dynamic Programming
// Problem detail:
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
// That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

package main

import "fmt"

func robHouseIITabular2(nums []int) int {
	robfromFirstHouseII := func(nums []int) int {
		n := len(nums)
		dp := make([]int, n)
		dp[0] = nums[0]

		// If Rob from 1st house
		for i := 1; i < n; i++ {
			rob := nums[i]
			if i > 1 {
				rob += dp[i-2]
			}
			notRob := dp[i-1]
			dp[i] = max(rob, notRob)
		}
		fmt.Printf("   dp:%v->%d\n", dp, dp[n-1])
		return dp[n-1]
	}

	// ------- Main func
	if len(nums) == 1 {
		return nums[0]
	}
	// Rob from 1st house
	fmt.Println("Rob from 1st house to last house-1")
	robFromFirstHouse := robfromFirstHouseII(nums[:len(nums)-1])
	fmt.Println("Rob from 2nd house to last house")
	robFromSecondHouse := robfromFirstHouseII(nums[1:])

	return max(robFromFirstHouse, robFromSecondHouse)
}

// https://leetcode.com/problems/house-robber-ii/discuss/2296761/Java-or-Memoization-or-Tabulation-or-Space-Optimization
// Runtime: 2 ms, faster than 45.37% of Go online submissions for House Robber II.
// Memory Usage: 2 MB, less than 85.78% of Go online submissions for House Robber II.
func robHouseIIMemorize2(nums []int) int {
	var maxMoney func(int, []int, []int) int
	maxMoney = func(currentHouse int, houses []int, dp []int) int {
		if currentHouse >= len(houses) {
			return 0
		}

		fmt.Printf("Rob house: %d\n", currentHouse)
		if dp[currentHouse] >= 0 {
			return dp[currentHouse]
		}
		robHouse := maxMoney(currentHouse+2, houses, dp) + houses[currentHouse]
		leaveHouse := maxMoney(currentHouse+1, houses, dp)
		if robHouse > leaveHouse {
			dp[currentHouse] = robHouse
		} else {
			dp[currentHouse] = leaveHouse
		}

		fmt.Printf("After Rob house: %d:%v -> %d\n", currentHouse, dp, dp[currentHouse])
		return dp[currentHouse]
	}

	// ------- Main func
	if len(nums) == 1 {
		return nums[0]
	}
	dp1 := make([]int, len(nums))
	for i := range dp1 {
		dp1[i] = -1
	}
	fmt.Println("Rob from first house to last house -1")
	robFromFirstHouse := maxMoney(0, nums[:len(nums)-1], dp1)

	fmt.Println("Rob from second house to last house")
	dp2 := make([]int, len(nums))
	for i := range dp2 {
		dp2[i] = -1
	}
	robFromSecondHouse := maxMoney(0, nums[1:], dp2)

	fmt.Printf(" After first:%d second:%d\n", robFromFirstHouse, robFromSecondHouse)
	if robFromFirstHouse > robFromSecondHouse {
		return robFromFirstHouse
	} else {
		return robFromSecondHouse
	}
}

// Dynamic Programing

func robHouseIITabular(nums []int) int {
	subRobHouseIITabular := func(houses []int) int {
		n := len(houses)
		if n == 1 {
			return houses[0]
		}

		dp := make([]int, n)
		dp[0] = houses[0]
		dp[1] = max(dp[0], houses[1])
		for i := 2; i < n; i++ {
			moneyFromHouse := houses[i] + dp[i-2]
			moneyFromAdjHouse := dp[i-1]
			dp[i] = max(moneyFromHouse, moneyFromAdjHouse)
			fmt.Printf("Rob from house[%d]:%d-%d =>%d\n", i, moneyFromHouse, moneyFromAdjHouse, dp[i])
		}
		fmt.Printf("dp:%v\n", dp)

		return dp[n-1]
	}

	// ------- Main func
	if len(nums) == 1 {
		return nums[0]
	}
	// Rob from 1st house
	fmt.Println("Rob from 1st house to last house-1")
	robFromFirstHouse := subRobHouseIITabular(nums[:len(nums)-1])
	fmt.Println("\n----\nRob from 2nd house to last house")
	robFromSecondHouse := subRobHouseIITabular(nums[1:])

	return max(robFromFirstHouse, robFromSecondHouse)

}

func robHouseIIMemorize(nums []int) int {
	// Rob from 1st house
	fmt.Println("Rob from 1st house to last house-1")
	robFromFirstHouse := robHouseMemorize(nums[:len(nums)-1])
	fmt.Println("\n----\nRob from 2nd house to last house")
	robFromSecondHouse := robHouseMemorize(nums[1:])

	return max(robFromFirstHouse, robFromSecondHouse)
}

func mainRobHouseII() {
	nums := []int{2, 3, 2}
	// nums = []int{1, 2, 3, 1}
	// nums = []int{1, 2, 3}
	// nums = []int{1, 3, 1, 3, 100}
	// nums = []int{4, 1, 2, 7, 5, 3, 1}
	nums = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	// nums = []int{1, 2}
	fmt.Printf("Houses:%v\n", nums)
	// fmt.Printf("Max money to Rob:%d\n", robHouseIITabular(nums))
	fmt.Printf("Max money to Rob:%d\n", robHouseIIMemorize(nums))
}
