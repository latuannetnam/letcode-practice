// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 242. Valid Anagram
// Letcode link: https://leetcode.com/problems/valid-anagram/
// Level: easy
// Topics: String, Hashtable
// Problem detail:
//Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.

package main

import (
	"fmt"
	"math"
)

// Runtime: 15 ms, faster than 25.76% of Go online submissions for Valid Anagram.
// Memory Usage: 3.9 MB, less than 9.91% of Go online submissions for Valid Anagram.
func isAnagramSorting(s string, t string) bool {
	s1 := sortString(s)
	t1 := sortString(t)
	return s1 == t1
}

// Runtime: 9 ms, faster than 55.16% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 82.00% of Go online submissions for Valid Anagram.
func isAnagramHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Build hash of all chars in source string
	srcHash := make(map[byte]int)
	for i := range s {
		srcHash[s[i]]++
	}
	fmt.Printf("srcHash:%v\n", srcHash)

	// Compare with dest string
	for i := range t {
		if srcHash[t[i]] > 0 {
			srcHash[t[i]]--
		} else {
			return false
		}
	}

	// Verify source hash
	for key := range srcHash {
		if srcHash[key] > 0 {
			return false
		}
	}

	return true
}

// Runtime: 14 ms, faster than 30.03% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 82.00% of Go online submissions for Valid Anagram
func isAnagramHash2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Build hash of all chars in source string and dest string
	srcHash := make(map[byte]int)
	for i := range s {
		srcHash[s[i]]++
		srcHash[t[i]]--
	}
	fmt.Printf("srcHash:%v\n", srcHash)

	// Verify source hash
	for key := range srcHash {
		if srcHash[key] != 0 {
			return false
		}
	}

	return true
}

// Failed
func isAnagramTrick(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := float64(0)
	for i := range s {
		// freq += int(s[i] * s[i] * s[i])
		// freq -= int(t[i] * t[i] * t[i])
		freq += math.Log(float64(s[i]))
		freq -= math.Log(float64(t[i]))
	}

	fmt.Printf("freq:%f\n", freq)
	return freq == 0
}

// https://leetcode.com/submissions/detail/763721073/
// Best solution: array[26] -> map of all chars in ASCII table
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

func isAnagramArrayMap(s string, t string) bool {
	alphabet := make([]int, 26)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		alphabet[s[i]-'a']++
		alphabet[t[i]-'a']--
	}

	for i := range alphabet {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

func mainValidAnagram() {
	s := "anagram"
	t := "nagaram"

	// s = "car"
	// t = "cat"

	// s = "hqbqo"
	// t = "lsnma"

	// s = "vbnxkji"
	// t = "wqdtegp"

	fmt.Printf("s:%s t:%s\n", s, t)
	fmt.Printf("is anagram:%t\n", isAnagramArrayMap(s, t))
}
