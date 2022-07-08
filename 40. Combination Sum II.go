// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/7/2022
// Letcode problem: 40. Combination Sum II
// Letcode link: https://leetcode.com/problems/combination-sum-ii/
// Level: medium
// Topics: Array, Backtracking
//Given a collection of candidate numbers (candidates) and a target number (target),
//find all unique combinations in candidates where the candidate numbers sum to target.
////Each number in candidates may only be used once in the combination.//
//Note: The solution set must not contain duplicate combinations.

package main

import (
	"fmt"
	"sort"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Combination Sum II.
//Memory Usage: 2.5 MB, less than 84.29% of Go online submissions for Combination Sum II.
func combinationSumII(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	var backTracking func(int, int, []int)
	backTracking = func(start int, sum int, combination []int) {
		for i := start; i < len(candidates); i++ {
			prevCandidate := candidates[i]
			sum += candidates[i]
			combination = append(combination, candidates[i])
			fmt.Printf("level:%d start:%d sum:%d target:%d combination:%v\n", len(combination), i, sum, target, combination)

			if sum >= target {
				if sum == target {
					subCombination := make([]int, len(combination))
					copy(subCombination, combination)
					result = append(result, subCombination)
					fmt.Printf("   Found combination: %v\n", subCombination)
				}
				fmt.Printf("  Backtracking ...\n")
				break
			}

			backTracking(i+1, sum, combination)
			//	Backtracking
			sum -= candidates[i]
			combination = combination[:len(combination)-1]

			//Skip duplicate
			for i < len(candidates)-1 && candidates[i+1] == prevCandidate {
				i++
			}
			fmt.Printf(" Backtrack :start:%d sum:%d target:%d combination:%v\n", i, sum, target, combination)

		}
	}
	backTracking(0, 0, []int{})
	return result
}

func mainCombinationSumII() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8

	//candidates = []int{2, 5, 2, 1, 2}
	//target = 5
	//
	//candidates = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	//target = 30

	//candidates = []int{1, 1, 1, 1, 1, 1}
	//target = 4

	fmt.Printf("Candidates: %d:%v target:%d\n", len(candidates), candidates, target)
	fmt.Printf("Combination sum: %v\n", combinationSumII(candidates, target))
}
