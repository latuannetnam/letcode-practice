// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 567. Permutation in String
// Letcode problem: Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
// In other words, return true if one of s1's permutations is the substring of s2.
// Level: medium

// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").

package main

import (
	"fmt"
	"reflect"
)

// Sliding Window, Two pointers, Hash
// Runtime: 163 ms, faster than 20.92% of Go online submissions for Permutation in String.
// Memory Usage: 8 MB, less than 6.56% of Go online submissions for Permutation in String.
func checkInclusion2(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	// Convert s1 to hash
	char_map := map[byte]int{}
	for _, char := range s1 {
		char_map[byte(char)]++
	}
	fmt.Printf("s1 hash:%v\n", char_map)
	start := 0
	end := 0
	temp_char_map := map[byte]int{}
	for start < len(s2) && end < len(s2) {
		fmt.Printf("start:%d end:%d slide:%s\n", start, end, s2[start:end+1])
		if char_map[s2[end]] > 0 {
			temp_char_map[s2[end]]++
		}

		if char_map[s2[start]] == 0 || char_map[s2[end]] == 0 {
			if char_map[s2[start]] > 0 {
				temp_char_map[s2[start]]--
			}
			start++
		}

		fmt.Printf("s1_map:%v s2_map:%v\n", char_map, temp_char_map)

		if end-start == len(s1)-1 {
			fmt.Println("Check permutation ...")
			if reflect.DeepEqual(char_map, temp_char_map) {
				fmt.Println("Found ...")
				return true
			} else {
				if char_map[s2[start]] > 0 {
					temp_char_map[s2[start]]--
				}
				start++
			}

		}

		end++

	}
	return false
}

// Runtime: 39 ms, faster than 31.03% of Go online submissions for Permutation in String.
// Memory Usage: 6.9 MB, less than 19.33% of Go online submissions for Permutation in String.
func checkInclusion3(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	// Convert s1 to hash
	char_map := map[byte]int{}
	for _, char := range s1 {
		char_map[byte(char)]++
	}
	fmt.Printf("s1 hash:%v\n", char_map)
	start := 0
	temp_char_map := map[byte]int{}
	for end := 0; end < len(s2); end++ {
		fmt.Printf("before start:%d end:%d slide:%s\n", start, end, s2[start:end+1])
		if char_map[s2[end]] > 0 {
			if len(temp_char_map) == 0 {
				start = end
			}

			temp_char_map[s2[end]]++
		} else {
			temp_char_map = map[byte]int{}
			start = end
		}

		fmt.Printf("after slide:%s s1_map:%v s2_map:%v\n", s2[start:end+1], char_map, temp_char_map)

		if end-start == len(s1)-1 {
			fmt.Println("Check permutation ...")
			if reflect.DeepEqual(char_map, temp_char_map) {
				fmt.Println("Found ...")
				return true
			} else {
				if char_map[s2[start]] > 0 {
					temp_char_map[s2[start]]--
				}
				start++
			}

		}
	}
	return false
}

// Runtime: 3 ms, faster than 86.70% of Go online submissions for Permutation in String.
// Memory Usage: 2.6 MB, less than 46.81% of Go online submissions for Permutation in String.
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	// Convert s1 to hash
	char_map := map[byte]int{}
	total := 0
	for _, char := range s1 {
		char_map[byte(char)]++
		total += int(char) * int(char)
	}
	fmt.Printf("s1 hash:%v total:%d\n", char_map, total)
	start := 0
	count := 0
	for end := 0; end < len(s2); end++ {
		fmt.Printf("before start:%d end:%d slide:%s\n", start, end, s2[start:end+1])
		if char_map[s2[end]] > 0 {
			if count == 0 {
				start = end
			}
			count += int(s2[end]) * int(s2[end])
		} else {
			count = 0
			start = end
		}

		fmt.Printf("after slide:%s count:%d\n", s2[start:end+1], count)
		if end-start == len(s1)-1 {
			fmt.Printf("Check permutation count:%d total:%d\n", count, total)
			if count == total {
				fmt.Println("Found ...")
				return true
			} else {
				if char_map[s2[start]] > 0 {
					count -= int(s2[start]) * int(s2[start])
				}
				start++
			}
		}

	}
	return false
}

func mainCheckInclusion() {
	s1 := "bab"
	s2 := "aeidbaabboooaa"
	fmt.Printf("s1:%s s2:%s\n", s1, s2)
	fmt.Printf("isPermutation: %t\n", checkInclusion(s1, s2))

	s1 = "adc"
	s2 = "dcda"
	fmt.Printf("\n\ns1:%s s2:%s\n", s1, s2)
	fmt.Printf("isPermutation: %t\n", checkInclusion(s1, s2))

	s1 = "abc"
	s2 = "ccccbbbbaaaa"
	fmt.Printf("\n\ns1:%s s2:%s\n", s1, s2)
	fmt.Printf("isPermutation: %t\n", checkInclusion(s1, s2))

}
