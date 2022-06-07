// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/7/2022
// Letcode problem: 46. Permutations
// Letcode link: https://leetcode.com/problems/permutations/
// Level: medium
//Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
package main

import "fmt"

//Backtrack & Recursion
//Runtime: 3 ms, faster than 50.89% of Go online submissions for Permutations.
//Memory Usage: 3 MB, less than 17.60% of Go online submissions for Permutations.
func makeArrayPermutations(nums []int, start, end int) [][]int {
	var newResult [][]int
	fmt.Printf("start:%d end:%d\n", start, end)
	if start == end-1 {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		fmt.Printf("Permutation:%v\n", newNums)
		newResult = append(newResult, newNums)
	} else {
		for i := start; i < end; i++ {
			fmt.Printf("Make permutation:%d-%d\n", start+1, end)
			nums[i], nums[start] = nums[start], nums[i]
			newResult = append(newResult, makeArrayPermutations(nums, start+1, end)...)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	return newResult
}

func mainArrayPermutations() {
	nums := []int{1, 2, 3}
	fmt.Printf("nums:%v\n", nums)
	result := makeArrayPermutations(nums, 0, len(nums))
	fmt.Printf("Permutations:%v\n", result)
}
