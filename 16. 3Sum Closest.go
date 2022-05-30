// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/3sum-closest/
// Level: medium
// Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
// Return the sum of the three integers.

package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	var first_sum bool = true
	var sum_closest int = 0
	var distance int
	var nums_len int = len(nums)
	var index_1, index_2, index_3 int
	// If not enough leng to make triplet
	if nums_len < 3 {
		return 0
	}

	sort.Ints(nums)
	fmt.Printf("sorted nums:%v \n", nums)
	for index_1 = 0; index_1 < nums_len-2; index_1++ {
		index_2 = index_1 + 1
		index_3 = nums_len - 1
		for index_2 < index_3 {
			var sum_temp int = nums[index_1] + nums[index_2] + nums[index_3]
			var distance_temp int = abs(target - sum_temp)

			if first_sum == true || distance_temp < distance {
				first_sum = false
				distance = distance_temp
				sum_closest = sum_temp
			}

			if sum_temp < target {
				index_2++
			} else if sum_temp > target {
				index_3--
			} else {
				return sum_closest
			}
		}
	}

	return sum_closest
}

func mainThreeSumClosest() {
	// var nums = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println("threeSumClosest")
	var nums = []int{82597, -9243, 62390, 83030, -97960, -26521, -61011, 83390, -38677, 12333, 75987}
	var target int = -1
	var sum_closest = threeSumClosest(nums, target)
	fmt.Printf("target: %d sum_closest:%d", target, sum_closest)
}
