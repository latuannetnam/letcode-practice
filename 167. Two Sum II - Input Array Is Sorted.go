// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 167. Two Sum II - Input Array Is Sorted
// Letcode problem:
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number.
// Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.
// Level: medium

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

package main

import "fmt"

// Runtime: 23 ms, faster than 24.84% of Go online submissions for Two Sum II - Input Array Is Sorted.
// Memory Usage: 5.4 MB, less than 50.04% of Go online submissions for Two Sum II - Input Array Is Sorted.
func twoSumSortedArray(numbers []int, target int) []int {
	numbers_len := len(numbers)
	index_1 := 0
	index_2 := numbers_len - 1
	for index_1 < index_2 {
		sum_temp := numbers[index_1] + numbers[index_2]
		if sum_temp < target {
			index_1++
		} else if sum_temp > target {
			index_2--
		} else {
			return ([]int{index_1 + 1, index_2 + 1})
		}
	}
	return []int{}
}

func mainTwoSumSortedArray() {
	var numbers []int
	var target int
	numbers = []int{2, 7, 11, 15}
	target = 9
	fmt.Printf("numbers:%v target: %d result:%v\n", numbers, target, twoSumSortedArray(numbers, target))

	numbers = []int{2, 3, 4}
	target = 6
	fmt.Printf("numbers:%v target: %d result:%v\n", numbers, target, twoSumSortedArray(numbers, target))

	numbers = []int{-1, 0}
	target = -1
	fmt.Printf("numbers:%v target: %d result:%v\n", numbers, target, twoSumSortedArray(numbers, target))

}
