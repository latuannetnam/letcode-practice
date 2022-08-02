// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 673. Number of Longest Increasing Subsequence
// Letcode link: https://leetcode.com/problems/number-of-longest-increasing-subsequence/
// Level: medium
// Topics: Array, Dynamic Programing
// Problem detail:
// Given an integer array nums, return the number of longest increasing subsequences.
// Notice that the sequence has to be strictly increasing.

package main

import (
	"fmt"
	"sort"
)

func findNumberOfLIS(nums []int) int {
	numOfLongest := 0
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	lis := make([]int, n)
	for i := range lis {
		lis[i] = 1
	}

	for i := 1; i < n; i++ {
		fmt.Printf("Find longest at index[%d]\n", i)
		for j := 0; j < i; j++ {
			// if nums[i] > nums[j] && dp[i] < dp[j]+1 {
			if nums[i] > nums[j] {
				lis[i] = lis[j] + 1
				dp[i][j] = lis[i]
				fmt.Printf(" path from %d->%d:%d\n", j, i, dp[i][j])
			}
		}
		fmt.Printf(" longest at index[%d]:%v->%d\n", i, dp[i], lis[i])
	}

	sort.Ints(lis)
	fmt.Printf("lis:%v\n", lis)
	fmt.Printf("dp:%v\n", dp)
	longest := lis[n-1]
	fmt.Printf("longest sequence:%d\n", longest)

	if longest > 1 {
		for i := range dp {
			sort.Ints(dp[i])
			for j := n - 1; j >= 0 && dp[i][j] == longest; j-- {
				numOfLongest++
			}
		}
	} else {
		numOfLongest = n
	}

	return numOfLongest
}

func mainFindNumberOfLIS() {
	nums := []int{1, 3, 5, 4, 7}
	nums = []int{2, 2, 2, 2, 2}
	nums = []int{1, 2, 4, 3, 5, 4, 7, 2}
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("Number of longest sequences:%d\n", findNumberOfLIS(nums))
}
