// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 189. Rotate Array
// Letcode problem: Given an array, rotate the array to the right by k steps, where k is non-negative.
// Level: medium
// Example:
// nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
package main

import "fmt"

// Runtime= 62ms, faster: 20.48%
// Space: 8.3M, < 71.28%
func rotateArray2(nums []int, k int) {
	n_len := len(nums)
	if k >= n_len {
		k = k % n_len
	}
	if k == 0 {
		return
	}

	nums2 := append(nums[n_len-k:], nums[:n_len-k]...)
	copy(nums, nums2)
}

// in-place and O(1) space
func shiftArrayRightOne(nums []int) {
	if len(nums) == 1 {
		return
	}
	temp := nums[len(nums)-1]
	index := len(nums) - 1
	for index > 0 {
		nums[index] = nums[index-1]
		index--
	}
	nums[0] = temp
}

// Runtime= 3319ms, faster: 5.43%
// Space: 7.7M, < 99.28%
func rotateArrayInPlace(nums []int, k int) {
	n_len := len(nums)
	k = k % n_len
	if k == 0 {
		return
	}

	for i := 0; i < k; i++ {
		shiftArrayRightOne(nums)
	}
}

// Rotate by swap part of array
// https://leetcode.com/problems/rotate-array/discuss/2049239/Python-O(n)-time-Solution-w-Comments
func rotateArray(nums []int, k int) {
	n_len := len(nums)
	k = k % n_len
	if k == 0 {
		return
	}
	reverseArray(nums, n_len-k, n_len)
	fmt.Printf("nums reverse n_len -k, n: %v\n", nums)
	reverseArray(nums, 0, n_len)
	fmt.Printf("nums reverse n: %v\n", nums)
	reverseArray(nums, k, n_len)

}

func mainRotateArray() {
	var nums []int
	var k int
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	k = 3
	fmt.Printf("nums:%v\n", nums)
	rotateArray(nums, k)
	fmt.Printf("rotated nums:%v\n", nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	fmt.Printf("nums:%v\n", nums)
	rotateArray(nums, k)
	fmt.Printf("rotated nums:%v\n", nums)

	nums = []int{-1}
	k = 2
	fmt.Printf("nums:%v\n", nums)
	rotateArray(nums, k)
	fmt.Printf("rotated nums:%v\n", nums)

	nums = []int{1, 2}
	k = 5
	fmt.Printf("nums:%v\n", nums)
	rotateArray(nums, k)
	fmt.Printf("rotated nums:%v\n", nums)
}
