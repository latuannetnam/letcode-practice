// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/4/2022
// Letcode problem: 78. Subsets
// Letcode link: https://leetcode.com/problems/subsets/
// Level: medium
// Topics: Backtracking, Array, Bit manipulation
//Given an integer array nums of unique elements, return all possible subsets (the power set).
//The solution set must not contain duplicate subsets. Return the solution in any order.
package main

import (
	"fmt"
	"sort"
	"strconv"
)

var backTrackingResultHash78 map[string][][]int = make(map[string][][]int)

func deepCopy2DArray(src [][]int) [][]int {
	var dst [][]int
	for i := range src {
		arr := make([]int, len(src[i]))
		copy(arr, src[i])
		dst = append(dst, arr)
	}
	return dst
}

//Runtime: 2 ms, faster than 50.24% of Go online submissions for Subsets.
//Memory Usage: 2.6 MB, less than 5.10% of Go online submissions for Subsets.
func subsetsBackTracking(item int, nums []int) [][]int {
	subKey := ""
	if len(nums) > 0 {
		//subKey = "-nums-" + arrayToString(nums, "-")
		subKey = "-nums-" + strconv.Itoa(nums[0]) + "-" + strconv.Itoa(nums[len(nums)-1])
	}
	key := strconv.Itoa(item) + subKey
	//key := arrayToString(orgNums, "-")

	result2 := backTrackingResultHash78[key]
	fmt.Printf("Backtracking: %s:%d %v: %v\n", key, item, nums, result2)
	if len(result2) == 0 {
		if len(nums) > 0 {
			for i := range nums {
				tempResult := deepCopy2DArray(subsetsBackTracking(nums[i], nums[i+1:]))
				for idx := range tempResult {
					tempResult[idx] = append(tempResult[idx], item)
				}
				result2 = append(result2, tempResult...)
			}
		}

		result2 = append(result2, []int{item})
		backTrackingResultHash78[key] = result2
	}

	fmt.Printf("After Backtracking: %s: %d: %v\n", key, item, result2)
	//fmt.Printf("backTrackingResultHash78:%v\n", backTrackingResultHash78)

	return result2
}

func subsets(nums []int) [][]int {
	var result [][]int = [][]int{{}}
	sort.Ints(nums)
	for i := range nums {
		tempResult := subsetsBackTracking(nums[i], nums[i+1:])
		fmt.Printf("Result:%d %v\n", nums[i], tempResult)
		result = append(result, tempResult...)
	}
	return result
}

//https://leetcode.com/problems/subsets/solution/
//Runtime: 0 ms, faster than 100.00% of Go online submissions for Subsets.
//Memory Usage: 2.4 MB, less than 24.08% of Go online submissions for Subsets.
func subsetsCascading(nums []int) [][]int {
	var result [][]int = [][]int{{}}
	sort.Ints(nums)
	for i := range nums {

		for idx := range result {
			var tempResult []int
			if len(result[idx]) > 0 {
				tempResult = make([]int, len(result[idx]))
				copy(tempResult, result[idx])
				tempResult = append(tempResult, nums[i])
			} else {
				tempResult = []int{nums[i]}
			}
			fmt.Printf("nums[%d]:%d tempResult[%d]:%v\n", i, nums[i], idx, tempResult)
			result = append(result, tempResult)
		}
		fmt.Printf("nums[%d]:%d  result:%v\n", i, nums[i], result)

	}
	return result
}

//https://leetcode.com/problems/subsets/solution/
func subsetsBacktrack(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	//Backtracking function
	var backTracking func(int, []int)
	backTracking = func(start int, subResult []int) {
		result = append(result, subResult)
		fmt.Printf("backTracking start:%d %v: %v\n", start, subResult, result)
		for i := start; i < len(nums); i++ {
			subResult = append(subResult, nums[i])
			backTracking(i+1, subResult)
			//	Backtrack
			if len(subResult) > 0 {
				subResult = subResult[1:]
			}
		}
	}
	//------------------------
	backTracking(0, []int{})
	return result
}

func mainSubsets() {
	nums := []int{1, 2, 3}
	//nums = []int{9, 0, 3, 5, 7}
	fmt.Printf("nums:%v\n", nums)
	//fmt.Printf("subsets:%v\n", subsetsCascading(nums))
	fmt.Printf("subsets:%v\n", subsetsBacktrack(nums))

}
