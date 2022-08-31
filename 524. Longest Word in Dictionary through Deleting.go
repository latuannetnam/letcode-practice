// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 524. Longest Word in Dictionary through Deleting
// Letcode link: https://leetcode.com/problems
// Level: medium
// Topics: Array, String, Two pointers, Sorting
// Problem detail:
//

package main

import (
	"fmt"
	"sort"
)

// Runtime: 21 ms, faster than 71.43% of Go online submissions for Longest Word in Dictionary through Deleting.
// Memory Usage: 8.2 MB, less than 11.09% of Go online submissions for Longest Word in Dictionary through Deleting.
func findLongestWordNotSort(s string, dictionary []string) string {
	result := ""
	for i := range dictionary {
		item := dictionary[i]
		dLen := len(item)
		rLen := len(result)
		if dLen >= rLen {
			sub := isSubsequence(item, s)
			fmt.Printf("Check subsequence of %s->%s:%t\n", dictionary[i], s, sub)
			if sub && ((rLen < dLen) || (rLen == dLen && result > item)) {
				result = dictionary[i]
				fmt.Printf("---> new Result:%s\n", result)
			}
		}
	}
	return result
}

// Sort dictionary first
// Runtime: 32 ms, faster than 21.43% of Go online submissions for Longest Word in Dictionary through Deleting.
// Memory Usage: 8.3 MB, less than 16.67% of Go online submissions for Longest Word in Dictionary through Deleting.
func findLongestWord(s string, dictionary []string) string {
	// Sort dictionary by length and lexico order
	sort.SliceStable(dictionary, func(i, j int) bool {
		return len(dictionary[i]) > len(dictionary[j]) || (len(dictionary[i]) == len(dictionary[j]) && dictionary[i] < dictionary[j])
	})

	fmt.Printf("sorted dictionary:%v\n", dictionary)

	for i := range dictionary {
		fmt.Printf("Check subsequence of %s->%s\n", dictionary[i], s)
		if isSubsequence(dictionary[i], s) {
			return dictionary[i]
		}
	}
	return ""
}

func mainFindLongestWord() {
	s := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}

	s = "abpcplea"
	dictionary = []string{"a", "b", "c"}

	s = "abce"
	dictionary = []string{"abe", "abc"}

	s = "aewfafwafjlwajflwajflwafj"
	dictionary = []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}

	fmt.Printf("s:%s\n", s)
	fmt.Printf("dictionary:%v\n", dictionary)
	fmt.Printf("longest word:%s\n", findLongestWord(s, dictionary))

}
