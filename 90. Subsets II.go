// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/4/2022
// Letcode problem: 90. Subsets II
// Letcode link: https://leetcode.com/problems/subsets-ii/
// Level: medium
// Topics:Backtracking, Array, Bit manipulation

package main

import (
	"fmt"
	"sort"
	"strconv"
)

//Cascade, remove duplicate by comparing duplicate sub array
//Runtime: 8 ms, faster than 6.34% of Go online submissions for Subsets II.
//Memory Usage: 2.6 MB, less than 30.99% of Go online submissions for Subsets II.
func subsetsWithDupCascade(nums []int) [][]int {
	var result [][]int = [][]int{{}}

	sort.Ints(nums)
	for i := range nums {
		var tempResult []int
		for idx := range result {
			tempResult = make([]int, len(result[idx]))
			copy(tempResult, result[idx])
			tempResult = append(tempResult, nums[i])
			if !checkSubArrayExits(tempResult, result) {
				result = append(result, tempResult)
			}
		}
		fmt.Printf("nums[%d]:%d result:%v\n", i, nums[i], result)
	}
	return result
}

//Cascade, remove duplicate by hashing array key
//Runtime: 6 ms, faster than 11.62% of Go online submissions for Subsets II.
//Memory Usage: 3.3 MB, less than 7.04% of Go online submissions for Subsets II.
func subsetsWithDupHash(nums []int) [][]int {
	var result [][]int = [][]int{{}}
	var arrayHash map[string]bool = make(map[string]bool)
	sort.Ints(nums)
	for i := range nums {
		var tempResult []int
		var subResult [][]int
		for idx := range result {
			tempResult = make([]int, len(result[idx]))
			copy(tempResult, result[idx])
			tempResult = append(tempResult, nums[i])
			key := arrayToString(tempResult, "-")
			if !arrayHash[key] {
				subResult = append(subResult, tempResult)
				arrayHash[key] = true
			}
		}
		result = append(result, subResult...)
		fmt.Printf("nums[%d]:%d subResult:%v\n", i, nums[i], subResult)
	}
	return result
}

//Runtime: 4 ms, faster than 32.65% of Go online submissions for Subsets II.
//Memory Usage: 3.4 MB, less than 5.84% of Go online submissions for Subsets II.
func subsetsWithDupHash2(nums []int) [][]int {
	var result [][]int = [][]int{{}}
	var resultKey []string = []string{""}
	var arrayHash map[string]bool = make(map[string]bool)
	sort.Ints(nums)
	for i := range nums {
		var tempResult []int
		var subResult [][]int
		for idx := range result {
			key := resultKey[idx] + "-" + strconv.Itoa(nums[i])
			resultKey = append(resultKey, key)
			tempResult = make([]int, len(result[idx]))
			copy(tempResult, result[idx])
			tempResult = append(tempResult, nums[i])
			subResult = append(subResult, tempResult)
		}
		result = append(result, subResult...)
		fmt.Printf("nums[%d]:%d resultKey:%v\n", i, nums[i], resultKey)
		//fmt.Printf("nums[%d]:%d subResult:%v\n", i, nums[i], subResult)
	}

	var result2 [][]int
	for i := range resultKey {
		key := resultKey[i]
		if !arrayHash[key] {
			result2 = append(result2, result[i])
			arrayHash[key] = true
		}
	}
	return result2
}

//Cascade, remove duplicate
func subsetsWithDupCascade2(nums []int) [][]int {
	var result [][]int = [][]int{{}}
	sort.Ints(nums)

	for i := range nums {
		var tempResult []int
		var subResult [][]int

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for idx := range result {
			tempResult = make([]int, len(result[idx]))
			copy(tempResult, result[idx])
			tempResult = append(tempResult, nums[i])
			subResult = append(subResult, tempResult)
		}
		result = append(result, subResult...)
		fmt.Printf("nums[%d]:%d subResult:%v\n", i, nums[i], subResult)
	}
	return result
}

//https://leetcode.com/problems/subsets/solution/
//https://leetcode.com/problems/subsets-ii/discuss/2238100/SUBSETS-II-%3A%3A-easy-approach
func subsetsBacktrackWithDup(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	//Backtracking function
	var backTracking func(int, []int)
	backTracking = func(start int, subResult []int) {
		tempResult := make([]int, len(subResult))
		copy(tempResult, subResult)
		result = append(result, tempResult)
		for i := start; i < len(nums); i++ {
			if i != start && nums[i] == nums[i-1] {
				continue
			}
			subResult = append(subResult, nums[i])
			fmt.Printf("backTracking start:%d %v: %v\n", start, subResult, result)
			backTracking(i+1, subResult)
			//	Backtrack
			subResult = subResult[:len(subResult)-1]
			fmt.Printf("After backTracking start:%d %v:%v\n", start, subResult, result)
		}
	}
	//------------------------
	backTracking(0, []int{})
	return result
}

func mainSubsetsWithDup() {
	nums := []int{1, 2, 2}
	//nums = []int{1, 2, 3}
	//nums = []int{5, 5, 5, 5, 5}
	//nums = []int{1, 2, 2, 3, 5, 5}
	fmt.Printf("nums:%v\n", nums)
	//fmt.Printf("subsets:%v\n", subsetsWithDupHash(nums))
	//fmt.Printf("\nsubsets:%v\n", subsetsWithDupHash2(nums))
	fmt.Printf("\nsubsets:%v\n", subsetsBacktrackWithDup(nums))
	//fmt.Printf("\nsubsets:%v\n", subsetsWithDupCascade2(nums))
}
