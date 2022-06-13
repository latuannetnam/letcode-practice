// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/13/2022
// Letcode problem: 136. Single Number
// Letcode link: https://leetcode.com/problems/flood-fill/
// Level: easy
//Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
//You must implement a solution with a linear runtime complexity and use only constant extra space.

package main

import "fmt"

//Runtime: 16 ms, faster than 87.23% of Go online submissions for Single Number.
//Memory Usage: 7.2 MB, less than 10.18% of Go online submissions for Single Number.
func singleNumber2(nums []int) int {
	lookup := map[int]int{}
	for i := range nums {
		lookup[nums[i]] += 1
	}
	for key := range lookup {
		if lookup[key] == 1 {
			return key
		}
	}
	return nums[0]
}

//https://leetcode.com/problems/single-number/discuss/1771771/Think-it-through-oror-Time%3A-O(n)-Space%3A-O(1)-oror-Python-Explained/
//Using XOR: 5^5 = 0   0^4=4 => 5^5^4 = 4
func singleNumber(nums []int) int {
	result := 0
	for i := range nums {
		result = result ^ nums[i]
	}
	return result
}

func mainSingleNumber() {
	nums := []int{4, 1, 2, 1, 2}
	//nums := []int{2, 1, 2}
	fmt.Printf("nums:%v single number:%d\n", nums, singleNumber(nums))
}
