// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/14/2022
// Letcode problem: 33. Search in Rotated Sorted Array
// Letcode link: https://leetcode.com/problems/search-in-rotated-sorted-array/
// Level: mediun
//There is an integer array nums sorted in ascending order (with distinct values).
//Prior to being passed to your function,
//nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
//such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
//For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

//Given the array nums after the possible rotation and an integer target,
//return the index of target if it is in nums, or -1 if it is not in nums.
//You must write an algorithm with O(log n) runtime complexity.
package main

import (
	"fmt"
)

func binarySearch(nums []int, target int, start int, end int) int {
	result := -1
	left := start
	right := end
	fmt.Printf("binarySearch start:%d end:%d %v:%d\n", start, end, nums[start:end+1], target)
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left:%d mid:%d right:%d:%d\n", left, mid, right, nums[mid])
		if nums[mid] == target {
			fmt.Printf("Found target %d at pos:%d\n", target, mid)
			result = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("binarySearch result:%v\n", result)
	return result
}

//Runtime: 3 ms, faster than 64.77% of Go online submissions for Search in Rotated Sorted Array.
//Memory Usage: 2.6 MB, less than 65.40% of Go online submissions for Search in Rotated Sorted Array.
func searchRotatedSortedArray2(nums []int, target int) int {
	result := -1
	if len(nums) == 0 {
		return result
	}

	end := len(nums) - 1
	rotatePos := -1

	//Find rotate position
	if nums[0] > nums[end] {
		left := 0
		right := len(nums) - 1
		rotatePos = left
		rotateTarget := nums[left]
		for left <= right {
			mid := left + (right-left)/2
			fmt.Printf("Search for rotate pos from %d-%d-%d:%d\n", left, mid, right, nums[mid])
			if nums[mid] >= rotateTarget {
				rotatePos = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		fmt.Printf("RotatePos:%d\n", rotatePos)
	}

	if rotatePos >= 0 {
		if nums[0] <= target && nums[rotatePos] >= target {
			result = binarySearch(nums, target, 0, rotatePos)
		} else if nums[rotatePos+1] <= target && nums[end] >= target {
			result = binarySearch(nums, target, rotatePos+1, end)
		} else {
			return result
		}
	} else if nums[0] <= target && nums[end] >= target {
		result = binarySearch(nums, target, 0, end)
	}

	return result
}

//https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/2144045/Best-and-Easy-C%2B%2B-code-or-Binary-Search-or-11ms
//Runtime: 4 ms, faster than 52.21% of Go online submissions for Search in Rotated Sorted Array.
//Memory Usage: 2.6 MB, less than 65.40% of Go online submissions for Search in Rotated Sorted Array.
func searchRotatedSortedArray(nums []int, target int) int {
	result := -1
	if len(nums) == 0 {
		return result
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
		if nums[left] <= nums[mid] {
			//	rotate pos in right half
			if nums[left] <= target && nums[mid] >= target {
				//Search in left half
				right = mid - 1
			} else {
				//Search in right half
				left = mid + 1
			}
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
	return result
}

func mainSearchRotatedSortedArray() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{1}
	//nums := []int{3, 1}
	//nums := []int{7, 8, 1, 2, 3, 4, 5, 6}
	//nums := []int{4, 5, 1, 2, 3}

	target := 5
	fmt.Printf("nums:%v target:%d\n", nums, target)
	result := searchRotatedSortedArray(nums, target)
	fmt.Printf("target:%d position:%d\n", target, result)
}
