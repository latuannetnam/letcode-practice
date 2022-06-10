// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/9/2022
// Letcode problem: 198. House Robber
// Letcode link: https://leetcode.com/problems/house-robber/
// Level: medium
//https://www.geeksforgeeks.org/solve-dynamic-programming-problem/
//You are a professional robber planning to rob houses along a street.
//Each house has a certain amount of money stashed,
//the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected
//and it will automatically contact the police if two adjacent houses were broken into on the same night.
////Given an integer array nums representing the amount of money of each house,
//return the maximum amount of money you can rob tonight without alerting the police.
package main

import "fmt"

//Dynamic Programing with Memorization
//Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
//Memory Usage: 2.1 MB, less than 37.29% of Go online submissions for House Robber.
func robFromFirstHouse(nums []int, firstHouse int, lookup []int) int {

	if lookup[firstHouse] < 0 {
		fmt.Printf("Rob from %d/%d\n", firstHouse, len(nums)-1)
		maxMoney := 0
		for house := firstHouse; house < len(nums); house++ {
			if house+2 < len(nums) {
				//fmt.Printf("Submoney from %d\n", house+2)
				lookup[house+2] = robFromFirstHouse(nums, house+2, lookup)
				if lookup[house+2] > maxMoney {
					maxMoney = lookup[house+2]
				}
				//fmt.Printf("After Submoney from %d:%d max:%d\n", house+2, lookup[house+2], maxMoney)
			}
		}
		lookup[firstHouse] = nums[firstHouse] + maxMoney
		fmt.Printf("After Rob from %d/%d:%d\n---\n", firstHouse, len(nums)-1, lookup[firstHouse])
	}

	return lookup[firstHouse]
}

func robHouse(nums []int) int {
	maxMoney := 0
	lookup := make([]int, len(nums))
	for i := range lookup {
		lookup[i] = -1
	}
	for i := range nums {
		fmt.Printf("\n-----\nFirst house:%d\n", i)
		lookup[i] = robFromFirstHouse(nums, i, lookup)
		if lookup[i] > maxMoney {
			maxMoney = lookup[i]
		}
		fmt.Printf("\nAfter First house:%d:%d - max:%d\n", i, lookup[i], maxMoney)
	}
	return maxMoney
}

func mainRobHouse() {
	//nums := []int{1, 2, 3, 1}
	//money := robHouse(nums)
	//fmt.Printf("nums:%v  money:%d\n", nums, money)
	//
	fmt.Printf("\n------\n")
	nums := []int{2, 7, 9, 3, 1}
	money := robHouse(nums)
	fmt.Printf("nums:%v  money:%d\n", nums, money)
	//
	//fmt.Printf("\n------\n")
	//nums = []int{2, 1, 1, 2}
	//money = robHouse(nums)
	//fmt.Printf("nums:%v  money:%d\n", nums, money)

	//nums := []int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}
	////nums := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//money := robHouse(nums)
	//fmt.Printf("nums:%v  money:%d\n", nums, money)

}
