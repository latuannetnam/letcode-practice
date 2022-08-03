// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 49. Group Anagrams
// Letcode link: https://leetcode.com/problems/group-anagrams/
// Level: medium
// Topics: Array, Hashtable, String
// Problem detail:
//Given an array of strings strs, group the anagrams together.
// You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

package main

import (
	"fmt"
	"strconv"
)

// Time limit Excedded
func groupAnagramsBrustForce(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}

	anagrams := [][]string{}
	processMap := make(map[string]bool)
	for i := range strs {
		if !processMap[strconv.Itoa(i)+strs[i]] {
			fmt.Printf("Find anagram group for [%d]:%s\n", i, strs[i])
			processMap[strconv.Itoa(i)+strs[i]] = true
			subAnagram := []string{strs[i]}
			for j := range strs {

				if !processMap[strconv.Itoa(j)+strs[j]] && isAnagramArrayMap(strs[i], strs[j]) {
					fmt.Printf("  found anagram[%d]:%s\n", j, strs[j])
					processMap[strconv.Itoa(j)+strs[j]] = true
					subAnagram = append(subAnagram, strs[j])
				}
			}
			fmt.Printf(" subAnagram:%v\n", subAnagram)
			anagrams = append(anagrams, subAnagram)
		}
	}
	return anagrams
}

// Runtime: 797 ms, faster than 5.05% of Go online submissions for Group Anagrams.
// Memory Usage: 7.2 MB, less than 96.76% of Go online submissions for Group Anagrams.
func groupAnagramsBrustForce2(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	anagrams := [][]string{}
	for i := range strs {
		fmt.Printf("Find anagram group for [%d]:%s\n", i, strs[i])
		found := false
		for j := range anagrams {
			if isAnagramArrayMap(anagrams[j][0], strs[i]) {
				found = true
				anagrams[j] = append(anagrams[j], strs[i])
				break
			}
		}
		if !found {
			anagrams = append(anagrams, []string{strs[i]})
		}
		fmt.Printf(" After:%v\n", anagrams)
	}
	return anagrams
}

// Runtime: 26 ms, faster than 89.16% of Go online submissions for Group Anagrams.
// Memory Usage: 8.1 MB, less than 69.80% of Go online submissions for Group Anagrams.
func groupAnagramsSort(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	anagrams := [][]string{}
	anagramsMap := make(map[string][]string)
	for i := range strs {
		fmt.Printf("Find anagram group for [%d]:%s\n", i, strs[i])
		sorted := sortString(strs[i])
		anagramsMap[sorted] = append(anagramsMap[sorted], strs[i])
		fmt.Printf(" After:%v\n", anagramsMap)
	}

	for key := range anagramsMap {
		anagrams = append(anagrams, anagramsMap[key])
	}
	return anagrams
}

func mainGroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// strs = []string{"", ""}
	fmt.Printf("strs:%v\n", strs)
	fmt.Printf("is Anagram:%t\n", isAnagramArrayMap("", ""))
	// fmt.Printf("anagram group:%v\n", groupAnagramsBrustForce2(strs))
	fmt.Printf("anagram group:%v\n", groupAnagramsSort(strs))

}
