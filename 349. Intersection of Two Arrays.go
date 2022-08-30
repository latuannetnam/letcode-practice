// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 349. Intersection of Two Arrays
// Letcode link: https://leetcode.com/problems/intersection-of-two-arrays/
// Level: medium
// Topics: Array, Two pointers, Hash, Sort
// Problem detail:
// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.
// Example 1:
// - Input: nums1 = [1,2,2,1], nums2 = [2,2]
// - Output: [2]

package main

import (
	"fmt"
	"sort"
)

// Hashtable
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Intersection of Two Arrays.
// Memory Usage: 3.2 MB, less than 12.53% of Go online submissions for Intersection of Two Arrays.
func intersectionHash1(nums1 []int, nums2 []int) []int {
	hash1 := make(map[int]bool)
	hash2 := make(map[int]bool)

	// Build hash of nums1's item
	for i := range nums1 {
		hash1[nums1[i]] = true
	}

	// Build hash of nums2's item
	for i := range nums2 {
		hash2[nums2[i]] = true
	}

	fmt.Printf("hash1:%v  hash2:%v\n", hash1, hash2)

	// Compare 2 hashes
	result := []int{}
	for key := range hash1 {
		if hash2[key] {
			result = append(result, key)
		}
	}
	return result
}

// Runtime: 12 ms, faster than 5.20% of Go online submissions for Intersection of Two Arrays.
// Memory Usage: 3.1 MB, less than 59.57% of Go online submissions for Intersection of Two Arrays.
func intersection(nums1 []int, nums2 []int) []int {
	hash1 := make(map[int]bool)

	// Build hash of nums1's item
	for i := range nums1 {
		hash1[nums1[i]] = true
	}

	fmt.Printf("hash1:%v \n", hash1)

	// Compare nums2 -> hash
	result := []int{}
	sort.Ints(nums2)
	prev := -1
	for i := range nums2 {
		if nums2[i] != prev {
			prev = nums2[i]
			if hash1[prev] {
				result = append(result, prev)
			}
		}

	}
	return result
}

func mainIntersection() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}

	fmt.Printf("nums1:%v nums2:%v\n", nums1, nums2)
	fmt.Printf("interesection:%v\n", intersection(nums1, nums2))

}
