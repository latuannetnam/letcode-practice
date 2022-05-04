// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/4sum/
// Level: medium
// Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//  - 0 <= a, b, c, d < n
//  - a, b, c, and d are distinct.
//  - nums[a] + nums[b] + nums[c] + nums[d] == target
// You may return the answer in any order.

package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) [][]int {
	result := [][]int{}
	nums_len := len(nums)
	index_1 := 0
	index_2 := nums_len - 1
	for index_1 < index_2 {
		sum_temp := nums[index_1] + nums[index_2]
		if sum_temp < target {
			index_1++
		} else if sum_temp > target {
			index_2--
		} else {
			triplet := []int{nums[index_1], nums[index_2]}
			result = append(result, triplet)
			index_1++
			index_2--
		}
	}
	return result

}

func threeSum(nums []int, target int) [][]int {
	result := [][]int{}
	nums_len := len(nums)
	for index_1 := 0; index_1 < nums_len-2; index_1++ {
		target_1 := target - nums[index_1]
		result_1 := twoSum(nums[index_1+1:], target_1)
		for twoSum_index := 0; twoSum_index < len(result_1); twoSum_index++ {
			triplet := result_1[twoSum_index]
			triplet = append([]int{nums[index_1]}, triplet...)
			result = append(result, triplet)
		}
		fmt.Printf("index_2:[%d]:%d target_1:%d 3sum result:%v\n", index_1, nums[index_1], target, result)

	}
	return result
}

func makeKey(triplet []int) string {
	key := ""
	for index := range triplet {
		key += fmt.Sprint(triplet[index])
	}
	return key
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	var result_map = map[string][]int{}
	// If not enough leng to make triplet
	nums_len := len(nums)
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	fmt.Printf("target:%d sorted nums:%v \n", target, nums)
	for index_1 := 0; index_1 < nums_len-3; index_1++ {
		target_1 := target - nums[index_1]
		result_1 := threeSum(nums[index_1+1:], target_1)
		for threeSum_index := 0; threeSum_index < len(result_1); threeSum_index++ {
			triplet := result_1[threeSum_index]
			triplet = append([]int{nums[index_1]}, triplet...)
			key := makeKey(triplet)
			result_map[key] = triplet
		}
		fmt.Printf("index_1: [%d]:%d target:%d result:%v\n", index_1, nums[index_1], target, result)
	}

	for key := range result_map {
		result = append(result, result_map[key])
	}
	return result
}

// Best solution: https://leetcode.com/problems/4sum/discuss/2005800/Go-sliding-window
func fourSumBest(nums []int, target int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	uniq := map[string]bool{}

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1

		for j < k {
			l := j + 1
			for l < k {
				sum := nums[i] + nums[j] + nums[l] + nums[k]
				if sum == target {
					str := fmt.Sprintf("%d%d%d%d", nums[1], nums[j], nums[l], nums[k])
					if !uniq[str] {
						results = append(results, []int{nums[i], nums[j], nums[l], nums[k]})
						uniq[str] = true
					}

					k--
					l++
				} else if sum < target {
					l++
				} else {
					k--
				}
			}

			j++
		}
	}

	return results
}

func mainFourSum() {
	// nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	result := fourSum(nums, target)
	fmt.Printf("Result: %v\n", result)
}
