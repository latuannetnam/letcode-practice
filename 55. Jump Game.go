// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 55. Jump Game
// Letcode link: https://leetcode.com/problems/jump-game/
// Level: medium
// Topics:  Array, Dynamic Programming
// Problem detail:
// You are given an integer array nums.
// You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.

package main

import "fmt"

// Runtime: 220 ms, faster than 27.03% of Go online submissions for Jump Game.
// Memory Usage: 9 MB, less than 13.01% of Go online submissions for Jump Game.
func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}

	var jumpStep func(int) int
	jumpStep = func(index int) int {

		if index <= n-1 && dp[index] >= 0 {
			// fmt.Printf(" index:%d -> dp[%d]:%d\n", index, index, dp[index])
			return dp[index]
		}

		if index >= n-1 {
			return index
		}

		maxStep := nums[index]
		fmt.Printf("Jump at index:%d maxStep:%d\n", index, maxStep)
		step := maxStep
		finalIndex := index
		for step = maxStep; step > 0; step-- {
			fmt.Printf(" index:%d step:%d finalIndex:%d\n", index, step, index+step)
			finalIndex = jumpStep(index + step)
			if finalIndex == n-1 {
				break
			}

		}

		dp[index] = finalIndex
		fmt.Printf(" Jump at index:%d step:%d dp:%v->%d\n", index, step, dp, finalIndex)

		return dp[index]
	}

	finalIndex := jumpStep(0)
	return finalIndex == n-1
}

// https://towardsdatascience.com/tackling-jump-game-problems-leetcode-e0a718e7dfba
// Runtime: 85 ms, faster than 68.77% of Go online submissions for Jump Game.
// Memory Usage: 7.5 MB, less than 61.86% of Go online submissions for Jump Game.
func canJumpLinearForward(nums []int) bool {
	maxReach := 0
	i := 0
	for i = 0; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(i+nums[i], maxReach)
		fmt.Printf("index:%d  maxReach:%d\n", i, maxReach)
		if maxReach == len(nums) {
			return true
		}
	}
	return i == len(nums)
}

func mainJumpGame() {
	nums := []int{2, 3, 1, 1, 4}
	nums = []int{3, 2, 1, 0, 4}
	nums = []int{2, 0}
	fmt.Printf("nums:%v\n", nums)
	// fmt.Printf("Can jumps:%t\n", canJump(nums))
	fmt.Printf("Can jumps:%t\n", canJumpLinearForward(nums))
}
