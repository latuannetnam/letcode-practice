// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/14/2022
// Letcode problem: 153. Find Minimum in Rotated Sorted Array
// Letcode link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Level: medium

package main

import "fmt"

//https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/2148989/C%2B%2B-or-binary-search-or-2-versions
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left:%d mid:%d right:%d:%d\n", left, mid, right, nums[mid])
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		fmt.Printf("new left:%d right:%d\n", left, right)
	}
	return nums[left]
}

//Runtime: 4 ms, faster than 51.82% of Go online submissions for Find Minimum in Rotated Sorted Array.
//Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Find Minimum in Rotated Sorted Array.
func findMin2(nums []int) int {
	result := -1
	if len(nums) == 0 {
		return result
	}
	left := 0
	right := len(nums) - 1
	rotatePos := left
	rotateTarget := nums[left]
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("Search for rotate pos from %d-%d-%d:%d\n", left, mid, right, nums[mid])
		if nums[mid] >= rotateTarget {
			rotatePos = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("RotatePos:%d\n", rotatePos)
	if rotatePos == len(nums)-1 {
		result = nums[0]
	} else {
		result = nums[rotatePos+1]
	}
	return result
}
func mainFindMinInRotatedSortedArray() {
	nums := []int{3, 4, 5, 1, 2}
	nums = []int{4, 5, 6, 7, 0, 1, 2}
	//nums = []int{3, 1, 2}
	fmt.Printf("nums:%v\n", nums)
	result := findMin(nums)
	fmt.Printf("Min:%d\n", result)
}
