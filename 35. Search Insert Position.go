// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 35. Search Insert Position
// Letcode problem: https://leetcode.com/problems/search-insert-position/
// Level: easy
// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

// You must write an algorithm with O(log n) runtime complexity.

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	index := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
			index = left
		} else {
			right = mid - 1
		}
	}
	return index
}

func mainSearchInsert() {
	nums := []int{1, 3, 5, 6}
	target := 5
	index := searchInsert(nums, target)
	fmt.Printf("nums:%v target:%d\n", nums, target)
	fmt.Printf("Insert position:%d\n", index)
}
