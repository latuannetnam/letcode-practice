// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/14/2022
// Letcode problem: 81. Search in Rotated Sorted Array II
// Letcode link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Level: medium
//There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
//Before being passed to your function,
//nums is rotated at an unknown pivot index k (0 <= k < nums.length)
//such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
//For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].
//Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.
//You must decrease the overall operation steps as much as possible.
package main

import "fmt"

//https://leetcode.com/problems/search-in-rotated-sorted-array-ii/solution/
//Runtime: 7 ms, faster than 41.13% of Go online submissions for Search in Rotated Sorted Array II.
//Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Search in Rotated Sorted Array II.
func searchRotatedSortedArrayII(nums []int, target int) bool {
	result := -1
	if len(nums) == 0 {
		return false
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left:%d mid:%d right:%d:%d\n", left, mid, right, nums[mid])
		if nums[mid] == target {
			fmt.Printf("Found target:%d at:%d\n", target, mid)
			result = mid
			break
		}
		//Find left or right half to search
		if nums[left] < nums[mid] {
			//	rotate pos in right half
			if nums[left] <= target && nums[mid] >= target {
				//Search in left half
				right = mid - 1
			} else {
				//Search in right half
				left = mid + 1
			}
		} else if nums[left] == nums[mid] {
			//Can not determine which half to search
			//Move left to next
			left++
		} else {
			//	rotate pos in left half
			if nums[mid] <= target && nums[right] >= target {
				//Search in right half
				left = mid + 1
			} else {
				//Search in left half
				right = mid - 1
			}

		}
		fmt.Printf("new left:%d right:%d\n", left, right)

	}
	return result >= 0
}

func mainSearchRotatedSortedArrayII() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	//nums = []int{1, 0, 1, 1, 1}
	target := 0
	fmt.Printf("nums:%v target:%d\n", nums, target)
	result := searchRotatedSortedArrayII(nums, target)
	fmt.Printf("target search result:%t\n", result)
}
