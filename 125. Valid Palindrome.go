// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 125. Valid Palindrome
// Letcode link: https://leetcode.com/problems/valid-palindrome/
// Level: easy
// Topics: Two pointers, String
// Problem detail:
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

package main

import "fmt"

// Runtime: 9 ms, faster than 42.79% of Go online submissions for Valid Palindrome.
// Memory Usage: 2.7 MB, less than 84.99% of Go online submissions for Valid Palindrome.
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
		} else if !isAlphanumeric(s[right]) {
			right--
		} else if charLowerCase(s[left]) == charLowerCase(s[right]) {
			left++
			right--
		} else {
			fmt.Printf("mismatch [%d]:%s -> [%d]:%s\n", left, string(s[left]), right, string(s[right]))
			return false
		}
	}
	return true
}

func mainValidPalindrome() {
	s := "A man, a plan, a canal: Panama"
	// s = "race a car"
	// s = " "
	fmt.Printf("String:%s\n", s)
	fmt.Printf("is valid palindrome:%t\n", isPalindrome(s))
}
