// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 977. Squares of a Sorted Array
// Letcode problem: https://leetcode.com/problems/squares-of-a-sorted-array/
// Level: easy
// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.

// Best: https://leetcode.com/problems/squares-of-a-sorted-array/discuss/2030651/Java-O(n)-solution
package main

import (
	"fmt"
	"sort"
)

func sortedSquares2(nums []int) []int {
	output := []int{}
	for _, num := range nums {
		output = append(output, num*num)
	}
	sort.Ints(output)
	return output
}

func insertSlice(nums []int, target int) []int {
	var output []int
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

	if index >= len(nums) {
		output = append(nums, target)
	} else {
		output = append(nums[:index+1], nums[index:]...)
		output[index] = target
	}
	// fmt.Printf("nums:%v target:%d output:%v\n", nums, target, output)
	return output
}

func mergeSlices(nums1 []int, nums2 []int) []int {
	min_nums1 := nums1[0]
	max_nums1 := nums1[len(nums1)-1]
	min_nums2 := nums2[0]
	max_nums2 := nums2[len(nums2)-1]
	output := []int{}
	fmt.Printf("merge %v and %v\n", nums1, nums2)
	if min_nums2 >= max_nums1 {
		// Append nums2 to nums1
		output = append(nums1, nums2...)
	} else if min_nums1 >= max_nums2 {
		// Prepend nums2 to nums1
		output = append(nums2, nums1...)
	} else {
		min_index := searchInsert(nums1, min_nums2)
		max_index := searchInsert(nums1, max_nums2)
		fmt.Printf("min index:%d max_index:%d\n", min_index, max_index)
		if min_index >= len(nums1) {
			output = append(nums1, nums2...)
		} else if min_index == max_index {
			output = append(output, nums1[:min_index]...)
			output = append(output, nums2...)
			output = append(output, nums1[min_index:]...)

		} else {
			output = append(output, nums1[:min_index]...)
			if max_index > len(nums1) {
				max_index = len(nums1)
			}

			nums3 := []int{}
			nums3 = append(nums3, nums1[min_index:max_index]...)
			output = append(output, mergeSlices(nums2, nums3)...)
			if max_index < len(nums1) {
				output = append(output, nums1[max_index:]...)
			}
		}
	}

	fmt.Printf("merge result:%v\n", output)
	return output
}

func sortedSquares(nums []int) []int {
	output := []int{}
	output2 := []int{}
	for _, num := range nums {
		if num >= 0 {
			// append to ascending sort slice with num > 0
			output = append(output, num*num)
		} else {
			// prepend to ascending sort slice with num < 0
			output2 = append([]int{num * num}, output2...)
		}
	}
	if len(output) == 0 {
		output = output2
	} else if len(output2) > 0 {
		output = mergeSlices(output, output2)
	}

	return output
}

func bestSortedSquares(nums []int) []int {
	squares := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	pos := right
	for left <= right {
		if abs(nums[right]) > abs(nums[left]) {
			squares[pos] = nums[right] * nums[right]
			right--
		} else {
			squares[pos] = nums[left] * nums[left]
			left++
		}
		pos--
	}
	return squares
}

func mainSortedSquares() {
	// nums := []int{-4, -1, 0, 3, 10}
	nums := []int{-7, -3, 2, 3, 11}
	// nums := []int{-1}
	// nums2 := []int{1, 9}
	// fmt.Printf("nums:%v\n", nums)
	// fmt.Printf("nums2:%v\n", nums2)
	// output := mergeSlices(nums, nums2)
	// fmt.Printf("merge nums:%v\n", output)
	// output := sortedSquares(nums)
	output := bestSortedSquares(nums)
	fmt.Printf("sortedSquares output:%v\n", output)
}
