// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 392. Is Subsequence
// Letcode link: https://leetcode.com/problems/is-subsequence/
// Level: easy
// Topics: String, Two pointer, Dynamic Progamming
// Problem detail:
// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters
// without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

package main

import "fmt"

// Two pointers
// Runtime: 1 ms, faster than 59.77% of Go online submissions for Is Subsequence.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Is Subsequence.
func isSubsequence(s string, t string) bool {
	sP := 0
	tP := 0
	for sP < len(s) && tP < len(t) {
		fmt.Printf("s[%d]:%s -> t[%d]:%s\n", sP, string(s[sP]), tP, string(t[tP]))
		if s[sP] == t[tP] {
			sP++
		}
		tP++
	}

	fmt.Printf("sP:%d/%d  tP:%d/%d\n", sP, len(s), tP, len(t))
	return sP == len(s)
}

func mainIsSubsequence() {
	s := "abc"
	t := "ahbgdc"

	s = "axc"
	t = "ahbgdc"

	fmt.Printf("s:%s t:%s\n", s, t)
	fmt.Printf("is subsequence:%t\n", isSubsequence(s, t))
}
