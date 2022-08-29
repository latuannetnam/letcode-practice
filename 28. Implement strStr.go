// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 28. Implement strStr()
// Letcode link: https://leetcode.com/problems/implement-strstr/
// Level: easy
// Topics: String, Two Pointers
// Problem detail:
// Implement strStr().
// Given two strings needle and haystack,
// return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack

package main

import "fmt"

// Two pointer
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
// Memory Usage: 2 MB, less than 61.91% of Go online submissions for Implement strStr().
func strStrSimulation(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	hPointer := 0
	nPointer := 0
	index := -1

	for hPointer < len(haystack) {
		fmt.Printf("haystack[%d]:%s needle[%d]:%s\n", hPointer, string(haystack[hPointer]), nPointer, string(needle[nPointer]))
		if needle[nPointer] == haystack[hPointer] {
			if index == -1 {
				index = hPointer
			}
			if nPointer == n-1 {
				return index
			}
			nPointer++
			hPointer++

		} else {
			nPointer = 0
			if index >= 0 {
				hPointer = index + 1
				index = -1
			} else {
				hPointer++
			}

		}

	}

	return -1
}

func mainStrStrSimulation() {
	haystack := "hello"
	needle := "ll"

	haystack = "aaaaa"
	needle = "abba"

	haystack = "aab"
	needle = "ab"

	fmt.Printf("string:%s substr:%s\n", haystack, needle)
	fmt.Printf("substring position:%d\n", strStrSimulation(haystack, needle))
}
