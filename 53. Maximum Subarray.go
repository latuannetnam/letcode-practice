// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 5/24/2022
// Letcode problem: 53. Maximum Subarray
// Letcode link: https://leetcode.com/problems/maximum-subarray/
// Level: easy
// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
// A subarray is a contiguous part of an array.

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
package main

import "fmt"

//Runtime: O(N^2) -> Time Limit Exceeded
func maxSubArrayBrustForce(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	for start := 0; start < len(nums); start++ {
		tempSum := nums[start]
		if tempSum > maxSum {
			maxSum = tempSum
		}
		for end := start + 1; end < len(nums); end++ {
			tempSum += nums[end]
			if tempSum > maxSum {
				maxSum = tempSum
			}
			fmt.Printf("start:%d end:%d tempSum:%d maxSum:%d\n", start, end, tempSum, maxSum)
		}
	}

	return maxSum
}

//Divide-Conquer: based on chapter 4 of book Introduction to Algorith 3rd Edition
type SubArraySum struct {
	low  int
	high int
	sum  int
}

func findMaxCrossingSubArray(nums []int, low int, mid int, high int) SubArraySum {
	subArraySum := SubArraySum{0, 0, 0}
	fmt.Printf("findMaxCrossingSubArray:%d -> %d -> %d\n", low, mid, high)
	leftSum := -100000
	sum := 0
	maxLeft := low
	for i := mid; i >= low; i-- {
		sum = sum + nums[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := -100000
	sum = 0
	maxRight := high
	for i := mid + 1; i <= high; i++ {
		sum = sum + nums[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}

	subArraySum.low = maxLeft
	subArraySum.high = maxRight
	subArraySum.sum = leftSum + rightSum

	fmt.Printf("findMaxCrossingSubArray:%d -> %d -> %d:%d\n", low, mid, high, subArraySum.sum)
	return subArraySum

}

func findMaxSubArray(nums []int, low, high int) SubArraySum {
	subArraySum := SubArraySum{0, 0, 0}
	fmt.Printf("findMaxSubArray:%d->%d\n", low, high)
	if high == low {
		subArraySum.low = low
		subArraySum.high = high
		subArraySum.sum = nums[low]
	} else {
		mid := (high + low) / 2
		leftSubArraySum := findMaxSubArray(nums, low, mid)
		rightSubArraySum := findMaxSubArray(nums, mid+1, high)
		crossSubArraySum := findMaxCrossingSubArray(nums, low, mid, high)
		if leftSubArraySum.sum >= rightSubArraySum.sum && leftSubArraySum.sum >= crossSubArraySum.sum {
			subArraySum.low = leftSubArraySum.low
			subArraySum.high = leftSubArraySum.high
			subArraySum.sum = leftSubArraySum.sum
		} else if rightSubArraySum.sum >= leftSubArraySum.sum && rightSubArraySum.sum >= crossSubArraySum.sum {
			subArraySum.low = rightSubArraySum.low
			subArraySum.high = rightSubArraySum.high
			subArraySum.sum = rightSubArraySum.sum
		} else {
			subArraySum.low = crossSubArraySum.low
			subArraySum.high = crossSubArraySum.high
			subArraySum.sum = crossSubArraySum.sum
		}
	}
	fmt.Printf("findMaxSubArray:%d->%d: %d\n", low, high, subArraySum.sum)
	return subArraySum
}

//Runtime: 141 ms, faster than 35.75% of Go online submissions for Maximum Subarray.
//Memory Usage: 9.9 MB, less than 19.76% of Go online submissions for Maximum Subarray.
func maxSubArray3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := findMaxSubArray(nums, 0, len(nums)-1)
	return maxSum.sum
}

//Best solution: 72ms -> From Leetcode forums: Kadane's algorithm
//https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	currentMax := nums[0]

	for i := 1; i < l; i++ {
		nums[i] = nums[i] + max(nums[i-1], 0)
		currentMax = max(currentMax, nums[i])
	}

	return currentMax
}

func mainMaxSubArray() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("nums:%v\n", nums)
	maxSum := maxSubArray(nums)
	fmt.Printf("Maxsum:%d\n", maxSum)

	fmt.Println("-----")
	nums = []int{5, 4, -1, 7, 8}
	fmt.Printf("nums:%v\n", nums)
	maxSum = maxSubArray(nums)
	fmt.Printf("Maxsum:%d\n", maxSum)

	fmt.Println("-----")
	nums = []int{-2, 1}
	fmt.Printf("nums:%v\n", nums)
	maxSum = maxSubArray(nums)
	fmt.Printf("Maxsum:%d\n", maxSum)

	fmt.Println("-----")
	nums = []int{5528, -444, -2520, 8588, -8658, 4122, -5493, 1748, 1306, 7475, 4488, -9532, 5605, -4589, 1711, 3504, 5213, -9923, 7613, 2445, -5999, -1594, 3961, -3236, 5057, 3745, 9724, 6429, -6768, 6502, -529, -6117, 9874, -7928, -6328, 9552, 5863, 8174, 8238, 3890, -4307, -4256, 3411, -7121, -4098, 7853, 6534, 4876, -8443, -4887, 8818, -9352, 1477, -9181, 5329, -8534, -8808, -4475, 6565, -4828, 7372, -7175, -7481, 9502, -9537, 6098, 7319, -670, 3462, -6303, 7642, -1855, 2606, -5868, -9450, -7795, 9992, -3483, 2826, -4342, 3735, -7497, -2607, -1179, 6110, 8429, 9537, 9818, -7671, -3510, 8399, -2240, -925, 5212, 129, 1647, 257, -8737, 661, 7999, 9755, 9728, -2180, -7810, -5312, 6902, 5031, -9149, 9385, 7185, -524, 8180, -3649, -9095, 2009, 9828, -2649, -4426, 7238, -123, 7480, 9958, -8045, 8470, -8384, -2623, -5963, -6873, -2489, -4476, -9587, 4577, 3901, -9167, 8196}

	//fmt.Printf("nums:%v\n", nums)
	maxSum = maxSubArray(nums)
	fmt.Printf("Maxsum:%d\n", maxSum)

}
