// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 128. Longest Consecutive Sequence
// Letcode link: https://leetcode.com/problems/longest-consecutive-sequence/
// Level: medium
// Topics: Array, Hash, Union Find
// Problem detail:
//Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.

package main

import (
	"fmt"
	"sort"
)

// Sort arrays and find longest
// Runtime: 65 ms, faster than 98.09% of Go online submissions for Longest Consecutive Sequence.
// Memory Usage: 9.7 MB, less than 79.71% of Go online submissions for Longest Consecutive Sequence.
func longestConsecutiveSort(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	}
	longest := 0
	first := 0
	i := 1
	skip := 0
	fmt.Printf("Sorted nums:%d\n", nums)
	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			skip++
		} else if nums[i] != nums[i-1]+1 {
			fmt.Printf("first:%d->%d  longest:%d->%d\n", first, i, longest, i-first-skip)
			if longest < i-first-skip {
				longest = i - first - skip
			}
			first = i
			skip = 0
		}
	}

	if longest < i-first-skip {
		longest = i - first - skip
	}
	return longest
}

func longestConsecutiveHash(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	longest := 0

	return longest
}

func mainLongestConsecutive() {
	nums := []int{100, 4, 200, 1, 3, 2}
	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	// nums = []int{1, 2, 0, 1}
	fmt.Printf("Nums:%v\n", nums)
	fmt.Printf("longestConsecutive:%d\n", longestConsecutiveSort(nums))
	// fmt.Printf("longestConsecutive:%d\n", longestConsecutiveHash(nums))
}
