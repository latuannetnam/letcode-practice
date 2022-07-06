// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/5/2022
// Letcode problem: 47. Permutations II
// Letcode link: https://leetcode.com/problems/permutations-ii/
// Level: medium
// Topics: Array, Backtrack
//Given a collection of numbers, nums, that might contain duplicates,
//return all possible unique permutations in any order.
package main

import (
	"fmt"
	"sort"
)

//Backtracking + Hash array
//Runtime: 100 ms, faster than 7.17% of Go online submissions for Permutations II.
//Memory Usage: 7.5 MB, less than 5.82% of Go online submissions for Permutations II.
func permuteUniqueBacktrackWithHashKey(nums []int) [][]int {
	var result [][]int
	var keyHash map[string]bool = make(map[string]bool)
	sort.Ints(nums)
	var backTracking func(int, int)
	backTracking = func(start int, end int) {
		fmt.Printf("start:%d end:%d\n", start, end)
		if start == end {
			key := arrayToString(nums, "-")
			if !keyHash[key] {
				keyHash[key] = true
				newNums := make([]int, len(nums))
				copy(newNums, nums)
				fmt.Printf("Permutation:%v\n", newNums)
				result = append(result, newNums)
			}

		} else {
			for i := start; i < end; i++ {
				if i != start && nums[i] == nums[start] {
					continue
				}
				nums[i], nums[start] = nums[start], nums[i]
				fmt.Printf("Make permutation:%d-%d:%d\n", start, i, nums)
				backTracking(start+1, end)
				nums[start], nums[i] = nums[i], nums[start]
			}
		}
	}
	backTracking(0, len(nums))
	return result
}

//https://leetcode.com/problems/permutations-ii/discuss/2234110/C%2B%2B-Simple-Backtracking
func checkSwap(nums []int, start, curr int) bool {
	for i := start; i < curr; i++ {
		if nums[i] == nums[curr] {
			return false
		}
	}
	return true
}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	var backTracking func(int, int)
	backTracking = func(start int, end int) {
		fmt.Printf("start:%d end:%d\n", start, end)
		if start == end {
			newNums := make([]int, len(nums))
			copy(newNums, nums)
			fmt.Printf("Permutation:%v\n", newNums)
			result = append(result, newNums)

		} else {
			for i := start; i < end; i++ {
				if checkSwap(nums, start, i) {
					nums[i], nums[start] = nums[start], nums[i]
					fmt.Printf("Make permutation:%d-%d:%d\n", start, i, nums)
					backTracking(start+1, end)
					nums[start], nums[i] = nums[i], nums[start]
				}

			}
		}
	}
	backTracking(0, len(nums))
	return result
}

func mainPermuteUnique() {
	nums := []int{1, 1, 2}
	//nums = []int{1, 2, 3}
	nums = []int{2, 2, 1, 1}
	fmt.Printf("nums:%v\n", nums)
	fmt.Printf("Permutation:%v\n", permuteUnique(nums))
}
