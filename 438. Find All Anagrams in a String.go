// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/21/2022
// Letcode problem: 438. Find All Anagrams in a String
// Letcode link: https://leetcode.com/problems/find-all-anagrams-in-a-string/
// Level: medium
//Given two strings s and p, return an array of all the start indices of p's anagrams in s.
//You may return the answer in any order.
//
//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
//typically using all the original letters exactly once.
package main

import "fmt"

//Runtime: 18 ms, faster than 58.46% of Go online submissions for Find All Anagrams in a String.
//Memory Usage: 5.2 MB, less than 57.54% of Go online submissions for Find All Anagrams in a String.
func findAnagrams2(s string, p string) []int {
	var result []int
	size := len(p)
	pointer := 0

	//Build hash map of every char in p
	anagrams := map[byte]int{}
	anagramCount := 0
	for _, char := range p {
		key := byte(char)
		anagrams[key]++
		anagramCount += int(key) * int(key)
	}

	fmt.Printf("Anagrams map:%v Count:%d\n", anagrams, anagramCount)

	//Slide window to scan matching s[pointer] in anagrams hash map
	lastPointer := pointer
	sourceCount := 0
	for lastPointer < len(s) {
		key := s[lastPointer]
		if anagrams[key] > 0 {
			sourceCount += int(key) * int(key)
		} else {
			pointer = lastPointer
			//Find first pointer contain anagrams chars
			for pointer < len(s)-size && anagrams[s[pointer]] == 0 {
				pointer++
			}
			lastPointer = pointer
			key = s[lastPointer]
			//Clean origin Count
			sourceCount = int(key) * int(key)
		}
		fmt.Printf("Pointer:%d lastPointer:%d key:%s orgCount:%d org:%s anagrams:%s\n", pointer, lastPointer, string(s[lastPointer]), sourceCount, s[pointer:lastPointer+1], p)

		if lastPointer-pointer+1 < size {
			lastPointer++
		} else {
			if sourceCount == anagramCount {
				fmt.Printf("Found position: %d org:%s anagrams:%s\n", pointer, s[pointer:lastPointer+1], p)
				result = append(result, pointer)
			}

			//	Increase pointer and decrease sourceCount
			orgKey := s[pointer]
			if anagrams[orgKey] > 0 {
				sourceCount -= int(orgKey) * int(orgKey)
			}
			orgKey = s[lastPointer]
			if anagrams[orgKey] > 0 {
				sourceCount -= int(orgKey) * int(orgKey)
			}
			pointer++
			if lastPointer <= pointer {
				//Clean origin count
				lastPointer = pointer
				sourceCount = 0
			}
		}
		fmt.Printf("After Pointer:%d lastPointer:%d orgCount:%d\n", pointer, lastPointer, sourceCount)
	}
	return result
}

//https://leetcode.com/problems/find-all-anagrams-in-a-string/discuss/2189256/C%2B%2B-or-Sliding-Window-or-Super-Clean-and-Concise-Code-or-Commented-or-Easy-to-understand
//Runtime: 23 ms, faster than 47.71% of Go online submissions for Find All Anagrams in a String.
//Memory Usage: 5.1 MB, less than 58.41% of Go online submissions for Find All Anagrams in a String.
func findAnagrams(s string, p string) []int {
	var result []int
	size := len(p)
	if size > len(s) {
		return result
	}
	pointer := 0

	//Build hash map of every char in p
	anagrams := map[byte]int{}
	anagramCount := 0
	for _, char := range p {
		key := byte(char)
		anagrams[key]++
		anagramCount += int(key) * int(key)
	}

	fmt.Printf("Anagrams map:%v Count:%d\n", anagrams, anagramCount)

	lastPointer := pointer
	sourceCount := 0
	//Build first window
	for lastPointer = pointer; lastPointer < size; lastPointer++ {
		key := s[lastPointer]
		if anagrams[key] > 0 {
			sourceCount += int(key) * int(key)
		}
	}
	fmt.Printf("Windows[%d-%d]:%s sourceCount:%d anagramCount:%d\n", pointer, lastPointer, s[pointer:lastPointer], sourceCount, anagramCount)
	if sourceCount == anagramCount {
		result = append(result, pointer)
	}

	//Slide window to scan matching s[pointer] in anagrams hash map
	for lastPointer < len(s) {
		//Decrease source count for pointer
		orgKey := s[pointer]
		if anagrams[orgKey] > 0 {
			sourceCount -= int(orgKey) * int(orgKey)
		}
		//Increase source count for lastPointer
		orgKey = s[lastPointer]
		if anagrams[orgKey] > 0 {
			sourceCount += int(orgKey) * int(orgKey)
		}
		pointer++
		lastPointer++
		fmt.Printf("Windows[%d-%d]:%s sourceCount:%d anagramCount:%d\n", pointer, lastPointer, s[pointer:lastPointer], sourceCount, anagramCount)
		if sourceCount == anagramCount {
			result = append(result, pointer)
		}
	}

	return result
}

func mainFindAnagrams() {
	s := "cbaebabacd"
	p := "abc"
	s = "abab"
	p = "ab"
	//s = "baa"
	//p = "aa"
	s = "acdcaeccde"
	p = "c"
	fmt.Printf("s:%s  p:%s\n", s, p)
	fmt.Printf("Anagrams:%v\n", findAnagrams(s, p))
}
