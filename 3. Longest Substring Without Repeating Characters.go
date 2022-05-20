// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 3. Longest Substring Without Repeating Characters
// Letcode problem: Given a string s, find the length of the longest substring without repeating characters.
// Level: medium

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

package main

import "fmt"

// Use sliding window
// Runtime: 7 ms, faster than 77.54% of Go online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 2.7 MB, less than 75.59% of Go online submissions for Longest Substring Without Repeating Characters.
func lengthOfLongestSubstring(s string) int {
	max_lenght := 0
	start := 0
	end := 0
	char_hash := map[byte]bool{}
	// Start window with size = 0
	for end < len(s) {
		if !char_hash[s[end]] {
			char_hash[s[end]] = true
			// Increase windows size
			end++
			// Store window size and compare with max length
			window_size := end - start
			if max_lenght < window_size {
				max_lenght = window_size
			}
			fmt.Printf("substr: %s max len:%d\n", s[start:end], max_lenght)
		} else {
			// Repeated char
			// Remove char_hash[s[start]]
			char_hash[s[start]] = false
			// Slide window
			start++
		}
	}
	return max_lenght
}

// Best solution
func lengthOfLongestSubstringBest(s string) int {
	location := [256]int{}

	for i := range location {
		location[i] = -1
	}

	left, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}
func mainLengthOfLongestSubstring() {
	s := "abcabcbb"
	fmt.Printf("string:%s\n", s)
	fmt.Printf("max len:%d\n", lengthOfLongestSubstring(s))

	s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	fmt.Printf("string:%s\n", s)
	fmt.Printf("max len:%d\n", lengthOfLongestSubstring(s))
}
