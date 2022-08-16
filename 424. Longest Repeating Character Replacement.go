// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 424. Longest Repeating Character Replacement
// Letcode link: https://leetcode.com/problems/longest-repeating-character-replacement/
// Level: medium
// Topics: Hashtable, String, Sliding Window
// Problem detail:
// You are given a string s and an integer k.
// You can choose any character of the string and change it to any other uppercase English character.
// You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations.

package main

import (
	"fmt"
	"sort"
)

func characterReplacementFailed(s string, k int) int {
	n := len(s)
	start := 0
	end := start
	size := 0
	count := 0
	maxSize := 0
	for start < n-maxSize && end < n {
		fmt.Printf("Start[%d]:%s End[%d]:%s Count:%d/%d size:%d\n", start, string(s[start]), end, string(s[end]), count, k, size)
		if s[start] == s[end] {
			size++
			end++
		} else {
			if count < k {
				count++
				size++
				end++
			} else {
				// Reset count
				count = 0
				size = 0
				start++
				end = start
			}
		}

		if size > maxSize {
			maxSize = size
		}
		fmt.Printf("  After Start:%d End:%d Count:%d/%d size:%d->%d\n", start, end, count, k, size, maxSize)

	}
	return maxSize
}

// Neetcode
// Runtime: 109 ms, faster than 7.65% of Go online submissions for Longest Repeating Character Replacement.
// Memory Usage: 7.6 MB, less than 5.10% of Go online submissions for Longest Repeating Character

// Runtime: 86 ms, faster than 8.84% of Go online submissions for Longest Repeating Character Replacement.
// Memory Usage: 6.8 MB, less than 5.10% of Go online submissions for Longest Repeating Character Replacement.
func characterReplacement2(s string, k int) int {
	n := len(s)
	maxSize := 0
	start := 0
	end := start
	charMap := make([]int, 26)
	maxFreq := 0
	maxCharFreq := func() int {
		arr := make([]int, 26)
		copy(arr, charMap)
		sort.Ints(arr)
		fmt.Printf("---->get maxFreg:%d\n", arr[len(arr)-1])
		return arr[len(arr)-1]
	}

	for start < n-maxSize && end < n {
		charMap[s[end]-65]++
		if charMap[s[end]-65] >= maxFreq {
			maxFreq = charMap[s[end]-65]
		} else {
			maxFreq = maxCharFreq()
		}

		fmt.Printf("Start[%d]:%s End[%d]:%s charMap:%v maxFreq:%d\n", start, string(s[start]), end, string(s[end]), charMap, maxFreq)

		// Check valid window size: maxSize - maxFreq <=k
		fmt.Printf(" check valid size:%d->%d\n", end-start+1-maxFreq, k)
		if end-start+1-maxFreq <= k {
			if end-start+1 > maxSize {
				maxSize = end - start + 1
			}
			end++
		} else {
			// move start
			charMap[s[start]-65]--
			charMap[s[end]-65]--
			start++
		}
		fmt.Printf(" After size:%d charMap:%v\n", maxSize, charMap)
	}
	return maxSize
}

// Neetcode
func characterReplacement(s string, k int) int {
	maxSize := 0
	start := 0
	maxFreq := 0
	var charMap [26]int

	for end := range s {
		charMap[s[end]-65]++
		if charMap[s[end]-65] >= maxFreq {
			maxFreq = charMap[s[end]-65]
		}

		fmt.Printf("Start[%d]:%s End[%d]:%s charMap:%v maxFreq:%d\n", start, string(s[start]), end, string(s[end]), charMap, maxFreq)
		if end-start+1-maxFreq > k {
			// move start
			charMap[s[start]-65]--
			start++
		}
		if end-start+1 > maxSize {
			maxSize = end - start + 1
		}

		fmt.Printf(" After size:%d \n", maxSize)

	}
	return maxSize
}

func mainCharacterReplacement() {
	s := "AABABBA"
	k := 1

	// s = "ABAB"
	// k = 2

	// s = "ABBB"
	// k = 2

	s = "EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH"
	k = 7

	fmt.Printf("string:%s k:%d\n", s, k)
	fmt.Printf("Longest string length:%d\n", characterReplacement(s, k))
}
