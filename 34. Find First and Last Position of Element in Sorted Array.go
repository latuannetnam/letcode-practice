// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/13/2022
// Letcode problem: 34. Find First and Last Position of Element in Sorted Array
// Letcode link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Level: medium
//Given an array of integers nums sorted in non-decreasing order,
//find the starting and ending position of a given target value.
//If target is not found in the array, return [-1, -1].
//You must write an algorithm with O(log n) runtime complexity.
package main

import "fmt"

type TargetResult struct {
	left  int
	mid   int
	right int
}

//Runtime: 9 ms, faster than 60.86% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//Memory Usage: 3.9 MB, less than 100.00% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange2(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left:%d mid:%d right:%d\n", left, mid, right)
		if nums[mid] == target {
			fmt.Printf("Found target %d at pos:%d\n", target, mid)
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Printf("left:%d right:%d\n", left, right)
	if left <= right {
		for nums[left] < target {
			left++
		}
		for nums[right] > target {
			right--
		}
		result[0] = left
		result[1] = right
	}

	return result
}

func searchTarget(nums []int, start, end, target int) TargetResult {
	result := TargetResult{-1, -1, -1}
	if len(nums) == 0 {
		return result
	}
	left := start
	right := end
	fmt.Printf("searchTarget start:%d end:%d %v\n", start, end, nums[start:end+1])
	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("left:%d mid:%d right:%d\n", left, mid, right)
		if nums[mid] == target {
			fmt.Printf("Found target %d at pos:%d\n", target, mid)
			result.left = left
			result.right = right
			result.mid = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("searchTarget result:%v\n", result)
	return result
}

//Runtime: 9 ms, faster than 60.86% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//Memory Usage: 4 MB, less than 14.67% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange3(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	start := 0
	end := len(nums) - 1
	tempResult := searchTarget(nums, start, end, target)
	if tempResult.mid >= 0 {
		result[0] = tempResult.mid
		result[1] = tempResult.mid
		tempResult2 := TargetResult{tempResult.left, tempResult.mid, tempResult.mid}
		for tempResult2.left < tempResult2.mid {
			tempResult2 = searchTarget(nums, tempResult2.left, tempResult2.mid, target)
			if tempResult2.mid >= 0 {
				result[0] = tempResult2.mid
			}
		}

		tempResult2 = TargetResult{tempResult.mid + 1, tempResult.mid, tempResult.right}
		for tempResult2.right > tempResult2.mid {
			tempResult2 = searchTarget(nums, tempResult2.mid+1, tempResult2.right, target)
			if tempResult2.mid >= 0 {
				result[1] = tempResult2.mid
			}
		}

	}
	return result
}

func searchTargetLeftOrRight(nums []int, target int, start, end int, isLeft bool) int {
	result := -1
	left := start
	right := end
	// fmt.Printf("searchTarget start:%d end:%d %v\n", start, end, nums[start:end+1])
	for left <= right {
		mid := left + (right-left)/2
		// fmt.Printf("left:%d mid:%d right:%d\n", left, mid, right)
		if nums[mid] == target {
			// fmt.Printf("Found target %d at pos:%d\n", target, mid)
			result = mid
			if isLeft {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// fmt.Printf("searchTarget result:%v\n", result)
	return result
}

//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/discuss/2146017/Java-binary-search-soln-oror-100-faster
//Runtime: 11 ms, faster than 36.85% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//Memory Usage: 4 MB, less than 14.66% of Go online submissions for Find First and Last Position of Element in Sorted Array.
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	start := 0
	end := len(nums) - 1
	left := -1
	right := -1
	// fmt.Println("Search left pos of target")
	left = searchTargetLeftOrRight(nums, target, start, end, true)
	if left >= 0 {
		start = left + 1
		// fmt.Println("Search right pos of target")
		right = searchTargetLeftOrRight(nums, target, start, end, false)
		if right == -1 {
			right = left
		}
	}

	result[0] = left
	result[1] = right
	return result
}

func mainSearchRange() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	//nums := []int{1, 2, 3, 8, 8, 8, 8, 9, 10, 11}
	nums := []int{1}
	target := 1
	fmt.Printf("nums:%v target:%d\n", nums, target)
	result := searchRange(nums, target)
	fmt.Printf("Target:%d Range:%d\n", target, result)

}
