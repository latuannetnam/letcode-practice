// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/binary-search/
// Level: easy

// Given an array of integers nums which is sorted in ascending order, and an integer target,
// write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.
// You must write an algorithm with O(log n) runtime complexity.

package main

import "fmt"

func search2(nums []int, target int) int {
	right := len(nums)
	var left int = len(nums) / 2
	for left < right {
		fmt.Printf("left:%d right:%d  nums[%d]:%d target:%d\n", left, right, left, nums[left], target)
		if nums[left] == target {
			return left
		} else if nums[left] < target {
			if right-left > 1 {
				left += (right - left) / 2
			} else {
				break
			}

		} else {
			right = left
			left = right / 2
		}
		fmt.Printf(" new left:%d new right:%d\n", left, right)
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func search(nums []int, target int) int {
	right := len(nums) - 1
	var left int = 0
	for left <= right {
		middle := left + (right-left)/2
		fmt.Printf("left:%d right:%d  nums[%d]:%d target:%d\n", left, right, middle, nums[middle], target)
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
		fmt.Printf(" new left:%d new right:%d\n", left, right)
	}

	return -1

}

func mainSearch() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	// nums := []int{5}
	target := 2
	fmt.Printf("target:%d nums:%v \n", target, nums)
	// var n int = len(nums) / 2
	// fmt.Print(n)
	fmt.Printf("result:%d\n", search(nums, target))
}
