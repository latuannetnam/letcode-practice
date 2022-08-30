// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 350. Intersection of Two Arrays II
// Letcode link: https://leetcode.com/problems/intersection-of-two-arrays-ii/
// Level: easy
// Topics:Array, Two pointers, Hash, Sort
// Problem detail:
// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

package main

import (
	"fmt"
	"sort"
)

// Runtime: 6 ms, faster than 43.89% of Go online submissions for Intersection of Two Arrays II.
// Memory Usage: 3.5 MB, less than 9.09% of Go online submissions for Intersection of Two Arrays II.
func intersectIIHash(nums1 []int, nums2 []int) []int {

	hash1 := make(map[int]int)
	hash2 := make(map[int]int)

	// Build hash of nums1's item
	for i := range nums1 {
		hash1[nums1[i]]++
	}

	// Build hash of nums2's item
	for i := range nums2 {
		hash2[nums2[i]]++
	}

	fmt.Printf("hash1:%v  hash2:%v\n", hash1, hash2)

	// Compare 2 hashes
	result := []int{}
	for key := range hash1 {
		if hash2[key] > 0 {
			repeated := min(hash1[key], hash2[key])
			fmt.Printf("Intersect:%d count:%d\n", key, repeated)
			for repeated > 0 {
				result = append(result, key)
				repeated--
			}

		}
	}

	return result
}

// https://leetcode.com/problems/intersection-of-two-arrays-ii/discuss/2501482/C%2B%2B-oror-2-Pointers-Sort-Arrays
//
//	Two Pointers, Sort array
//
// Runtime: 2 ms, faster than 86.99% of Go online submissions for Intersection of Two Arrays II.
// Memory Usage: 2.8 MB, less than 94.20% of Go online submissions for Intersection of Two Arrays II.
func intersectIITwoPointers(nums1 []int, nums2 []int) []int {
	result := []int{}
	sort.Ints(nums1)
	sort.Ints(nums2)
	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			p1++
			p2++
		} else if nums1[p1] < nums2[p2] {
			p1++
		} else {
			p2++
		}

	}
	return result
}

func mainIntersectII() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	// nums1 = []int{4, 9, 5}
	// nums2 = []int{9, 4, 9, 8, 4}

	// nums1 = []int{1, 2, 2, 1}
	// nums2 = []int{2}

	fmt.Printf("nums1:%v nums2:%v\n", nums1, nums2)
	fmt.Printf("interesection II:%v\n", intersectIITwoPointers(nums1, nums2))
}
