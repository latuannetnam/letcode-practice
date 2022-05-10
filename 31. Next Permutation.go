// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/next-permutation/
// Level: medium
// Given an array of integers nums, find the next permutation of nums
// For example, the next permutation of arr = [1,2,3] is [1,3,2].
// Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
// While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func buildPermutation(nums []int) []string {
	result := []string{}
	result_map := map[string]bool{}
	if len(nums) == 2 {
		result = []string{strconv.Itoa(nums[0]) + "-" + strconv.Itoa(nums[1]), strconv.Itoa(nums[1]) + "-" + strconv.Itoa(nums[0])}

	} else {
		for index := range nums {
			item := nums[index]
			reduce_nums := removeSliceItemIndex(nums, index)
			fmt.Printf("item:%d reduce num:%v\n", item, reduce_nums)
			permutations := buildPermutation(reduce_nums)
			for permu_index := range permutations {
				key := strconv.Itoa(item) + "-" + permutations[permu_index]
				if !result_map[key] {
					result = append(result, key)
					result_map[key] = true
				}

			}

		}
	}
	return result
}

func arrayToString(arr []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprintf("%03v", arr), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func intSlice2String(nums []int) string {
	return strings.Trim(fmt.Sprintf("%03v", nums), "[]")
}

func string2IntSlice(str string) []int {
	s_slice := strings.Split(str, " ")
	var nums []int
	for index := range s_slice {
		num, err := strconv.Atoi(s_slice[index])
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func minSlice2String(nums []int) string {
	sorted_nums := sortSlice(nums)
	return intSlice2String(sorted_nums)
}

func maxSlice2String(nums []int) string {
	sorted_nums := sortSliceDesc(nums)
	return intSlice2String(sorted_nums)
}

func minSliceWithIndex2String(nums []int, index int) string {
	reduced_nums := removeSliceItemIndex(nums, index)
	min_reduced_nums_str := minSlice2String(reduced_nums)
	min_index_sorted_num_str := fmt.Sprintf("%03d", nums[index]) + " " + min_reduced_nums_str
	return min_index_sorted_num_str
}

func removeSliceItemIndex(s []int, i int) []int {
	var s1 []int
	s1 = append(s1, s...)
	return append(s1[:i], s1[i+1:]...)
}

func sortSliceDesc(nums []int) []int {
	sorted_nums := make([]int, len(nums))
	copy(sorted_nums, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted_nums)))
	return sorted_nums
}

func sortSlice(nums []int) []int {
	sorted_nums := make([]int, len(nums))
	copy(sorted_nums, nums)
	sort.Ints(sorted_nums)
	return sorted_nums
}
func minSliceWithIndex(nums []int, index int) []int {
	reduced_nums := removeSliceItemIndex(nums, index)
	new_nums := []int{nums[index]}
	new_nums = append(new_nums, sortSlice(reduced_nums)...)
	return new_nums
}

// Compare 2 slice: 0: equal, 1: nums1> nums2, -1: nums1< nums2
//  2 slice must be equal in size
func compareSlice(nums1, nums2 []int) int {
	flag := 0
	for index := range nums1 {
		if nums1[index] > nums2[index] {
			return 1
		} else if nums1[index] < nums2[index] {
			return -1
		}
	}
	return flag
}

func findNextPermutation(nums []int) []int {
	sorted_nums := sortSliceDesc(nums)
	next_permutation := []int{}
	left_num := 0
	left := 0
	fmt.Printf("-----findNextPermutation---\n")
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("sorted nums: %v\n", sorted_nums)
	for left < len(sorted_nums) {
		min_index_sorted_num := minSliceWithIndex(sorted_nums, left)
		fmt.Printf("for left: %d/%d left_num:%d  \n", left, len(sorted_nums), sorted_nums[left])
		if compareSlice(min_index_sorted_num, nums) > 0 {
			left++
			next_permutation = min_index_sorted_num
			fmt.Printf(" increase left: %d/%d next_permutation:%v\n", left, len(sorted_nums), next_permutation)
		} else {
			fmt.Printf(" %v <= %v \n", min_index_sorted_num, nums)
			break
		}
	}

	if left == len(sorted_nums) {
		return next_permutation
	} else {
		reduced_nums := nums[1:]
		fmt.Printf("after for left: %d/%d left_num:%d  \n", left, len(sorted_nums), sorted_nums[left])
		if compareSlice(reduced_nums, sortSliceDesc(reduced_nums)) == 0 {
			return next_permutation
		} else {
			left_num = sorted_nums[left]
			temp_next_permutation := findNextPermutation(reduced_nums)
			new_nums := []int{left_num}
			return append(new_nums, temp_next_permutation...)
		}

	}
}

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	flag := compareSlice(nums, sortSliceDesc(nums))
	if flag == 0 {
		copy(nums, sortSlice(nums))
		fmt.Printf("nums is already max. Use min:%v\n", nums)
	} else {
		next_permutation := findNextPermutation(nums)
		copy(nums, next_permutation)
	}
}

func mainNextPermutation() {
	// nums := []int{4, 3, 2, 1}
	// nums := []int{1, 2, 3, 4}
	// nums2 := []int{1, 2, 4, 2}
	// nums := []int{1, 1, 5}
	// nums := []int{1}
	// nums := []int{6, 7, 5, 3, 5, 6, 2, 9, 1, 2, 7, 0, 9}
	nums := []int{16, 27, 25, 23, 25, 16, 12, 9, 1, 2, 7, 20, 19, 23, 16, 0, 6, 22, 16, 11, 8, 27, 9, 2, 20, 2, 13, 7, 25, 29, 12, 12, 18, 29, 27, 13, 16, 1, 22, 9, 3, 21, 29, 14, 7, 8, 14, 5, 0, 23, 16, 1, 20}
	// nums1 := []int{16, 27, 25, 23, 25, 16, 12, 9, 1, 2, 7, 20, 19, 23, 16, 0, 6, 22, 16, 11, 8, 27, 9, 2, 20, 2, 13, 7, 25, 29, 12, 12, 18, 29, 27, 13, 16, 1, 22, 9, 3, 21, 29, 14, 7, 8, 14, 5, 0, 23, 16, 20, 1}
	fmt.Printf("nums:%v\n", nums)
	// fmt.Printf("nums2:%v compare:%d\n", nums2, compareSlice(nums, nums2))
	nextPermutation(nums)
	fmt.Printf("Next permutation:%v\n", nums)
}
