// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 287. Find the Duplicate Number
// Letcode link: https://leetcode.com/problems/find-the-duplicate-number/
// Level: medium
// Topics: Array, Two pointers, Binary Search, Bit manupilation
// Problem detail:
// Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and uses only constant extra space.

package main

import "fmt"

// Runtime: 131 ms, faster than 65.44% of Go online submissions for Find the Duplicate Number.
// Memory Usage: 9.2 MB, less than 48.62% of Go online submissions for Find the Duplicate Number.
func findDuplicateMap(nums []int) int {
	n := len(nums)
	numMap := make([]bool, n)
	for i := range nums {
		if numMap[nums[i]] {
			return nums[i]
		} else {
			numMap[nums[i]] = true
		}
	}

	return 1
}

// Neetcode
//  Link list cycle + Floyd algorithm
// Runtime: 153 ms, faster than 52.30% of Go online submissions for Find the Duplicate Number.
// Memory Usage: 8.6 MB, less than 87.56% of Go online submissions for Find the Duplicate Number.
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		fmt.Printf("Next slow:%d. Next fast:%d\n", slow, fast)
		if slow == fast {
			break
		}
	}

	slow2 := 0

	for {
		slow = nums[slow]
		slow2 = nums[slow2]
		fmt.Printf("Next slow:%d. Next slow2:%d\n", slow, slow2)
		if slow == slow2 {
			return slow
		}
	}

}

func mainFindDuplicate() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("repeated num:%d\n", findDuplicate(nums))

}
