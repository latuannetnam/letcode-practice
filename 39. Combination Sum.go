// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/6/2022
// Letcode problem: 39. Combination Sum
// Letcode link: https://leetcode.com/problems/combination-sum/
// Level: medium
// Topics: Array, Backtracking
//Given an array of distinct integers candidates and a target integer target,
//return a list of all unique combinations of candidates where the chosen numbers sum to target.
//You may return the combinations in any order.
//The same number may be chosen from candidates an unlimited number of times.
//Two combinations are unique if the frequency of at least one of the chosen numbers is different.
//It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.
package main

import (
	"fmt"
	"sort"
)

//Backtrack 1
//Runtime: 6 ms, faster than 54.04% of Go online submissions for Combination Sum.
//Memory Usage: 3.4 MB, less than 55.73% of Go online submissions for Combination Sum.
func combinationSum3(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var backTracking func(int, int, int, []int) bool
	backTracking = func(start int, sum int, count int, combination []int) bool {
		fmt.Printf("Begin of start:%d count:%d\n", start, count)
		if sum >= target {
			if sum == target {
				subCombination := make([]int, len(combination))
				copy(subCombination, combination)
				result = append(result, subCombination)
				fmt.Printf("   Found combination: %v\n", subCombination)
			}
			fmt.Printf("  Backtracking ....\n")
			return false
		}
		i := start
		for i < len(candidates) {
			count++
			sum += candidates[i]
			combination = append(combination, candidates[i])
			fmt.Printf(" start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)
			if !backTracking(i, sum, count, combination) {
				//	Backtracking
				count = 0
				sum -= candidates[i]
				combination = combination[:len(combination)-1]
				//	Expand i
				i++
			}
			//fmt.Printf("  After start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)
		}
		fmt.Printf("  End of start:%d count:%d\n", start, count)
		return false
	}

	backTracking(0, 0, 0, []int{})
	return result
}

//Backtrack 3
//Runtime: 4 ms, faster than 73.31% of Go online submissions for Combination Sum.
//Memory Usage: 3.6 MB, less than 48.18% of Go online submissions for Combination Sum.
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var backTracking func(int, int, []int)
	backTracking = func(start int, sum int, combination []int) {
		//fmt.Printf("Begin of start:%d sum:%d target:%d\n", start, sum, target)

		//for i < len(candidates) && sum+candidates[i] <= target {
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			combination = append(combination, candidates[i])
			fmt.Printf("start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)
			if sum >= target {
				if sum == target {
					subCombination := make([]int, len(combination))
					copy(subCombination, combination)
					result = append(result, subCombination)
					fmt.Printf("   Found combination: %v\n", subCombination)
				}
				fmt.Printf("  Backtracking ....\n")
				break
			}

			backTracking(i, sum, combination)
			//	Backtracking
			sum -= candidates[i]
			combination = combination[:len(combination)-1]
			fmt.Printf(" Backtrack :start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)

		}
		return
	}

	backTracking(0, 0, []int{})
	return result
}

//https://leetcode.com/problems/combination-sum/discuss/1777569/FULL-EXPLANATION-WITH-STATE-SPACE-TREE-oror-Recursion-and-Backtracking-oror-Well-Explained-oror-C%2B%2B
func combinationSum5(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var backTracking func(int, int, []int)
	backTracking = func(start int, sum int, combination []int) {
		if sum >= target {
			if sum == target {
				subCombination := make([]int, len(combination))
				copy(subCombination, combination)
				result = append(result, subCombination)
				fmt.Printf("   Found combination: %v\n", subCombination)
			}
			fmt.Printf("  Backtracking ....\n")
			return
		}
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			combination = append(combination, candidates[i])
			if sum > target {
				break
			}
			fmt.Printf("start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)
			backTracking(i, sum, combination)
			//	Backtracking
			sum -= candidates[i]
			combination = combination[:len(combination)-1]
			fmt.Printf(" Backtrack :start:%d-%d sum:%d target:%d combination:%v\n", start, i, sum, target, combination)
		}
	}
	backTracking(0, 0, []int{})
	return result
}
func mainCombinationSum() {
	candidates := []int{2, 3, 6, 7}
	target := 7

	candidates = []int{2, 3, 5}
	target = 8

	//candidates = []int{2}
	//target = 1
	fmt.Printf("Candidates: %v target:%d\n", candidates, target)
	fmt.Printf("Combination sum: %v\n", combinationSum(candidates, target))

}
