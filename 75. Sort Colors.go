// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 75. Sort Colors
// Letcode link: https://leetcode.com/problems/sort-colors/
// Level: medium
// Topics: Array, Two pointers, Sort
// Problem detail:
// Given an array nums with n objects colored red, white, or blue,
// sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.
// Follow up: Could you come up with a one-pass algorithm using only constant extra space?
package main

import "fmt"

// Counting + 2 pass
// Runtime: 2 ms, faster than 53.69% of Go online submissions for Sort Colors.
// Memory Usage: 2.1 MB, less than 69.29% of Go online submissions for Sort Colors.

func sortColors2Pass(nums []int) {
	var countArr [3]int
	// Counting number 0, 1, 2 in nums
	for i := range nums {
		countArr[nums[i]]++
	}

	fmt.Printf("countArr:%v\n", countArr)

	// Overwrite nums with counting of 0, 1, 2
	pointer := 0
	for i := range countArr {
		for j := 0; j < countArr[i]; j++ {
			nums[pointer+j] = i
		}
		pointer += countArr[i]

	}

}

func mainSortColors() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Printf("nums:%v\n", nums)
	sortColors2Pass(nums)
	fmt.Printf("sorted nums:%v\n", nums)

}
