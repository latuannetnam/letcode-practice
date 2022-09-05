// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 532. K-diff Pairs in an Array
// Letcode link: https://leetcode.com/problems/k-diff-pairs-in-an-array/
// Level: medium
// Topics: Array, Hashtable, Two pointer, Binary search, Sorting
// Problem detail:
// Given an array of integers nums and an integer k, return the number of unique k-diff pairs in the array.
// A k-diff pair is an integer pair (nums[i], nums[j]), where the following are true:
// 		0 <= i, j < nums.length
// 		i != j
// 		nums[i] - nums[j] == k
// Notice that |val| denotes the absolute value of val.

package main

import (
	"fmt"
	"sort"
	"strconv"
)

// Runtime: 36 ms, faster than 5.00% of Go online submissions for K-diff Pairs in an Array.
// Memory Usage: 3.9 MB, less than 94.00% of Go online submissions for K-diff Pairs in an Array.
func findPairsBrustForce(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	fmt.Printf("Sorted nums:%v\n", nums)

	count := 0
	prev := nums[0]
	for i := 0; i < n-1; i++ {
		if i > 0 {
			if nums[i] != prev {
				prev = nums[i]
			} else {
				continue
			}
		}

		fmt.Printf("Check from index:%d/%d\n", i, n-2)
		for j := i + 1; j < n; j++ {
			fmt.Printf(" Scan from index:%d/%d\n", j, n-1)
			if nums[j]-nums[i] == k {
				fmt.Printf("found pair:[%d:%d]\n", nums[i], nums[j])
				count++
				break
			}
		}
	}
	return count
}

// Two poiters + Sort + Hash
// Runtime: 26 ms, faster than 12.00% of Go online submissions for K-diff Pairs in an Array.
// Memory Usage: 6.7 MB, less than 29.00% of Go online submissions for K-diff Pairs in an Array.
func findPairsTwoPointers(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	fmt.Printf("Sorted nums:%v\n", nums)
	pairHash := make(map[string]bool)

	count := 0
	left := 0
	right := 1
	for right < n {
		// for left < right && right<n {
		fmt.Printf("left[%d]:%d right[%d]:%d\n", left, nums[left], right, nums[right])
		if nums[right]-nums[left] == k {
			pairKey := strconv.Itoa(nums[left]) + ":" + strconv.Itoa(nums[right])
			if !pairHash[pairKey] {
				count++
				pairHash[pairKey] = true
			}

			fmt.Printf("found pair:[%d:%d] count:%d\n", nums[left], nums[right], count)

			right++
		} else if nums[right]-nums[left] < k {
			right++
		} else {
			left++
		}
		if left == right {
			right++
		}
	}
	return count
}

// https://leetcode.com/problems/k-diff-pairs-in-an-array/discuss/1756874/C%2B%2B-MULTIPLE-APPROACHES-%3A-MAPS-TWO-POINTER
// Hash Map
func findPairs(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	count := 0
	// Build frequency map for each number in array
	numFreq := make(map[int]int)
	for i := range nums {
		numFreq[nums[i]]++
	}

	// Find num + k in map
	for i := range nums {
		if k == 0 {
			if numFreq[nums[i]] > 1 {
				count++
				fmt.Printf("Found pair nums[%d]:%d count:%d\n", i, nums[i], count)
			}
		} else if numFreq[nums[i]+k] > 0 {
			count++
			fmt.Printf("Found pair nums[%d]:%d count:%d\n", i, nums[i], count)
		}
	}

	return count
}

func mainFindPairs() {
	nums := []int{3, 1, 4, 1, 5}
	k := 2

	nums = []int{1, 3, 1, 5, 4}
	k = 0

	nums = []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}
	k = 3
	fmt.Printf("nums:%v k:%d\n", nums, k)
	fmt.Printf("num pairs:%d\n", findPairs(nums, k))

}
