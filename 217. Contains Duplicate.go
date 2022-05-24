// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 5/24/2022
// Letcode problem: 217. Contains Duplicate
// Letcode link: https://leetcode.com/problems/contains-duplicate/
// Level: easy
//Given an integer array nums, return true if any value appears at least twice in the array,
//and return false if every element is distinct.
//Input: nums = [1,2,3,1]
//Output: true

package main

import (
	"fmt"
	"sort"
)

//Sorted array, then compare
//Runtime: 115 ms, faster than 30.88% of Go online submissions for Contains Duplicate.
//Memory Usage: 8.4 MB, less than 91.83% of Go online submissions for Contains Duplicate.
func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	prevNum := nums[0]
	if len(nums) <= 1 {
		return false
	}
	for index := 1; index < len(nums); index++ {
		if nums[index] == prevNum {
			return true
		} else {
			prevNum = nums[index]
		}
	}
	return false
}

//Hashtable, check count of key
//Runtime: 65 ms, faster than 93.10% of Go online submissions for Contains Duplicate.
//Memory Usage: 8.5 MB, less than 87.95% of Go online submissions for Contains Duplicate.
func containsDuplicate(nums []int) bool {
	numFreqs := map[int]int{}
	for _, num := range nums {
		numFreqs[num]++
		if numFreqs[num] > 1 {
			return true
		}
	}
	return false
}

func mainContainsDuplicate() {
	nums := []int{1, 2, 3, 1}
	fmt.Printf("nums:%v\n", nums)
	result := containsDuplicate(nums)
	fmt.Printf("containsDuplicate:%t\n", result)

	nums = []int{1, 2, 3, 4}
	fmt.Printf("nums:%v\n", nums)
	result = containsDuplicate(nums)
	fmt.Printf("containsDuplicate:%t\n", result)
}
