// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 45. Jump Game II
// Letcode link: https://leetcode.com/problems/jump-game-ii/
// Level: medium
// Topics:  Array, Dynamic Programming
// Problem detail:
// Given an array of non-negative integers nums, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
// You can assume that you can always reach the last index.

package main

import "fmt"

func minimumJumpFailed(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	maxReach := 0
	i := 0
	for i = 0; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(i+nums[i], maxReach)
		fmt.Printf("index:%d  maxReach:%d\n", i, maxReach)
		if maxReach >= len(nums)-1 {
			return i + 1
		}
	}
	return i + 1
}

func minimumJumpFailed2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	minSteps := 100000

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	count := 0

	var jumpStep func(int, int) int
	jumpStep = func(index int, numSteps int) int {
		count++
		// if count > 500 {
		// 	return numSteps
		// }

		// if index <= n-1 && dp[index] >= 0 {
		// 	fmt.Printf(" index:%d -> dp[%d]:%d->%d\n", index, index, dp[index], minSteps)
		// 	dp[index] = min(dp[index], minSteps)
		// 	return dp[index]
		// }

		if index >= n-1 {
			return numSteps
		}

		maxReach := nums[index]
		fmt.Printf("Jump at index:%d maxReach:%d\n", index, maxReach)
		distance := maxReach
		finalSteps := numSteps
		for distance = maxReach; distance > 0; distance-- {
			fmt.Printf(" index:%d distance:%d numSteps:%d nextIndex:%d\n", index, distance, numSteps, index+distance)
			finalSteps = jumpStep(index+distance, numSteps+1)
			minSteps = min(minSteps, finalSteps)
			fmt.Printf("      After Jump at index:%d distance:%d numSteps:%d->%d \n", index, distance, finalSteps, minSteps)
		}

		// minSteps = min(minSteps, finalSteps)
		dp[index] = minSteps
		// fmt.Printf(" Jump at index:%d step:%d finalIndex:%d numSteps:%d \n", index, step, finalIndex, numSteps)
		return dp[index]
	}

	jumpStep(0, 0)
	fmt.Printf("dp:%v\n", dp)
	return minSteps
}

// https://towardsdatascience.com/tackling-jump-game-problems-leetcode-e0a718e7dfba
// Runtime: 673 ms, faster than 5.53% of Go online submissions for Jump Game II.
// Memory Usage: 5.9 MB, less than 99.83% of Go online submissions for Jump Game II.
func minimumJumpDP(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 100000
	}

	dp[0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] >= i-j {
				dp[i] = min(dp[i], dp[j]+1)
			}

		}
	}

	return dp[n-1]
}

// https://towardsdatascience.com/tackling-jump-game-problems-leetcode-e0a718e7dfba
func minimumJumpOptmized(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	maxReach := nums[0]
	steps := nums[0]
	jumps := 0
	for i := 1; i < n; i++ {
		maxReach = max(maxReach, nums[i]+i)
		steps--
		if steps == 0 {
			jumps++
			steps = maxReach - 1
		}
	}
	return jumps + 1
}

func mainMinumJump() {
	nums := []int{2, 3, 1, 1, 4}
	nums = []int{2, 3, 0, 1, 4}
	nums = []int{0}
	nums = []int{2, 1}
	nums = []int{1, 2, 1, 1, 1}
	nums = []int{9, 8, 2, 2, 0, 2, 2, 0, 4, 1, 5, 7, 9, 6, 6, 0, 6, 5, 0, 5}
	nums = []int{5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5}
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("min jumps:%d\n", minimumJumpDP(nums))
	// fmt.Printf("min jumps:%d\n", minimumJumpOptmized(nums))
}
