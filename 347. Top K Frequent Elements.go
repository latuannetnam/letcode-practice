// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 347. Top K Frequent Elements
// Letcode link: https://leetcode.com/problems/top-k-frequent-elements/
// Level: medium
// Topics: Array, Divide and Conquer, Hash, Sort
// Problem detail:
// Given an integer array nums and an integer k, return the k most frequent elements.
// You may return the answer in any order.

package main

import (
	"fmt"
	"sort"
)

// Runtime: 27 ms, faster than 20.68% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.5 MB, less than 83.93% of Go online submissions for Top K Frequent Elements.
func topKFrequentHash(nums []int, k int) []int {
	freq := []int{}
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
	}

	fmt.Printf(" freqMap:%v\n", freqMap)
	// Sort freqMap by frequency
	keys := make([]int, 0)
	for key := range freqMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	fmt.Printf("keys:%v\n", keys)
	for i := 0; i < k; i++ {
		freq = append(freq, keys[i])
	}

	return freq
}

// Runtime: 46 ms, faster than 6.00% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 6 MB, less than 36.93% of Go online submissions for Top K Frequent Elements.
func topKFrequentHashAndBinarySearch(nums []int, k int) []int {
	freq := []int{}
	freqMap := make(map[int]int)
	sort.Ints(nums)

	lastIndex := 0
	for lastIndex <= len(nums)-1 {
		target := nums[lastIndex]
		rangeTarget := searchRange(nums[lastIndex:], target)
		fmt.Printf("target:%d firstIndex:%d range:%v\n", target, lastIndex, rangeTarget)
		freqMap[target] = rangeTarget[1] + 1
		lastIndex += rangeTarget[1] + 1
	}

	fmt.Printf(" freqMap:%v\n", freqMap)
	// Sort freqMap by frequency
	keys := make([]int, 0)
	for key := range freqMap {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})

	fmt.Printf("keys:%v\n", keys)
	for i := 0; i < k; i++ {
		freq = append(freq, keys[i])
	}
	return freq
}

// https://neetcode.io/
// Runtime: 21 ms, faster than 42.66% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.9 MB, less than 43.12% of Go online submissions for Top K Frequent Elements.
func topKFrequentHashAndBucketSort(nums []int, k int) []int {
	freq := []int{}
	freqMap := make(map[int]int)
	for i := range nums {
		freqMap[nums[i]]++
	}
	fmt.Printf("freqMap:%v\n", freqMap)
	bucket := make([][]int, len(nums)+1)
	for key := range freqMap {
		bucket[freqMap[key]] = append(bucket[freqMap[key]], key)
	}

	fmt.Printf("bucket:%v\n", bucket)
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) <= k {
			freq = append(freq, bucket[i]...)
			k -= len(bucket[i])
		} else {
			freq = append(freq, bucket[i][:k]...)
			k = 0
		}
		if k == 0 {
			break
		}

	}

	return freq
}

// Runtime: 21 ms, faster than 42.66% of Go online submissions for Top K Frequent Elements.
// Memory Usage: 5.9 MB, less than 36.93% of Go online submissions for Top K Frequent Elements.
func topKFrequentHashAndBinarySearchAndBucketSort(nums []int, k int) []int {
	freq := []int{}
	freqMap := make(map[int]int)
	sort.Ints(nums)

	lastIndex := 0
	for lastIndex <= len(nums)-1 {
		target := nums[lastIndex]
		rangeTarget := searchRange(nums[lastIndex:], target)
		fmt.Printf("target:%d firstIndex:%d range:%v\n", target, lastIndex, rangeTarget)
		freqMap[target] = rangeTarget[1] + 1
		lastIndex += rangeTarget[1] + 1
	}

	fmt.Printf(" freqMap:%v\n", freqMap)
	bucket := make([][]int, len(nums)+1)
	for key := range freqMap {
		bucket[freqMap[key]] = append(bucket[freqMap[key]], key)
	}

	fmt.Printf("bucket:%v\n", bucket)
	for i := len(bucket) - 1; i > 0; i-- {
		if len(bucket[i]) <= k {
			freq = append(freq, bucket[i]...)
			k -= len(bucket[i])
		} else {
			freq = append(freq, bucket[i][:k]...)
			k = 0
		}
		if k == 0 {
			break
		}

	}
	return freq
}

func mainTopKFrequent() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	nums = []int{1, 2}
	k = 2
	nums = []int{1, 1, 2, 2, 3, 3, 4}
	k = 2

	fmt.Printf("nums:%v k:%d\n", nums, k)
	fmt.Printf("topKFrequent:%v\n", topKFrequentHashAndBinarySearchAndBucketSort(nums, k))
}
