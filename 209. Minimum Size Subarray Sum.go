// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/27/2022
// Letcode problem: 209. Minimum Size Subarray Sum
// Letcode link: https://leetcode.com/problems/minimum-size-subarray-sum/
// Level: medium
// Topics: Binary Search, Sliding Window, Prefix Sum
//Given an array of positive integers nums and a positive integer target,
//return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target.
//If there is no such subarray, return 0 instead.
package main

import (
	"fmt"
	"math"
)

//Time Limit Exeeded
//Sliding Window
func minSubArrayLen_old(target int, nums []int) int {
	result := 0
	size := len(nums)
	firstSum := 0

	//Slide window
	for windowSize := 1; windowSize <= size; windowSize++ {
		fmt.Printf("\n----windowSize:%d\n", windowSize)
		//Build first window with windowSize=N
		if firstSum == 0 {
			for i := 0; i < windowSize; i++ {
				firstSum += nums[i]
			}
		} else {
			firstSum += nums[windowSize-1]
		}
		fmt.Printf("First sum:%d target:%d\n", firstSum, target)
		sum := firstSum
		if sum >= target {
			fmt.Printf("Match sum>target\n")
			result = windowSize
			return result
		}

		//	Slide window with windowSize=N
		for windowStart := 1; windowStart <= size-windowSize; windowStart++ {
			if nums[windowStart-1] != nums[windowStart+windowSize-1] {
				//	Decrease arrProduct
				sum -= nums[windowStart-1]
				//	Increase arrProduct
				sum += nums[windowStart+windowSize-1]
			}

			fmt.Printf("windowStart:%d Next sum:%d target:%d\n", windowStart, sum, target)
			if sum >= target {
				fmt.Printf("Match sum>target\n")
				result = windowSize
				return result
			}
		}
		fmt.Printf("windowSize:%d result:%d\n", windowSize, result)
	}
	return result
}

//Time Limit Exeeded
//Prefix sum
//https://leetcode.com/problems/minimum-size-subarray-sum/solution/
func minSubArrayLen_prefixsum(target int, nums []int) int {
	result := math.MaxInt
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	fmt.Printf("Prefix sum: %v\n", sums)

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := sums[j] - sums[i] + nums[i]
			fmt.Printf("i:%d j:%d sum:%d target:%d\n", i, j, sum, target)
			if sum >= target {
				result = min(result, j-i+1)
				fmt.Printf("Found min length: %d\n", result)
				break
			}
		}
	}

	if result == math.MaxInt {
		result = 0
	}
	return result
}

//Prefix sum + Binary search
//https://leetcode.com/problems/minimum-size-subarray-sum/discuss/2187213/Java-Solution-oror-BinarySearch-oror-PrefixSum
func minSubArrayLen_PrefixSum_BinarySearch(target int, nums []int) int {
	result := math.MaxInt
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	fmt.Printf("Prefix sum: %v\n", sums)

	//Return 0 if sum of array < target
	if sums[len(nums)-1] < target {
		return 0
	}
	//Binary search for min sub-array length
	for i := 0; i < len(nums); i++ {
		num := sums[i] - target
		idx := binarySearch(sums, num, 0, i)
		if idx >= 0 {
			result = min(result, i-idx)
		}
	}

	if result == math.MaxInt {
		result = 0
	}
	return result
}

//https://leetcode.com/problems/minimum-size-subarray-sum/solution/
//2-pointers, sliding window
//Runtime: 36 ms, faster than 26.99% of Go online submissions for Minimum Size Subarray Sum.
//Memory Usage: 8.3 MB, less than 8.19% of Go online submissions for Minimum Size Subarray Sum.
func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt

	left := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			result = min(result, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if result == math.MaxInt {
		result = 0
	}
	return result
}

func mainMinSubArrayLen() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	//nums = []int{1, 4, 4}
	//target = 4
	//nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	//target = 11
	fmt.Printf("nums:%v target:%d\v", nums, target)
	fmt.Printf("target:%d min_length:%d\n", target, minSubArrayLen(target, nums))
}
