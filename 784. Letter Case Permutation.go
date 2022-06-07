// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/7/2022
// Letcode problem: 784. Letter Case Permutation
// Letcode link: https://leetcode.com/problems/letter-case-permutation/
// Level: medium
//Given a string s, you can transform every letter individually to be lowercase or uppercase to create another string.
//Return a list of all possible strings we could create. Return the output in any order.
//Input: s = "a1b2"
//Output: ["a1b2","a1B2","A1b2","A1B2"]
package main

import (
	"fmt"
	"strings"
)

func isLowercase(char rune) bool {
	if char >= 97 && char <= 122 {
		return true
	}
	return false
}

func isUppercase(char rune) bool {
	if char >= 65 && char <= 90 {
		return true
	}
	return false
}

//Runtime: 31 ms, faster than 7.75% of Go online submissions for Letter Case Permutation.
//Memory Usage: 7.3 MB, less than 7.75% of Go online submissions for Letter Case Permutation.
func makeLetterCasePermutation2(arr []rune, start int, end int, results *map[string]bool) {
	fmt.Printf("Start:%d end:%d arr[%d]:%d arr[%d:]:%s\n", start, end, start, arr[start], start, string(arr[start:]))
	if start == end-1 {
		//fmt.Printf("Result:%s\n", string(arr))
		(*results)[string(arr)] = true
		//temp := arr[start]
		if isLowercase(arr[start]) {
			//Convert to Uppercase
			arr[start] = arr[start] - 32
			(*results)[string(arr)] = true
		} else if isUppercase(arr[start]) {
			//Convert to Lowercase
			arr[start] = arr[start] + 32
			(*results)[string(arr)] = true
		}
		//fmt.Printf("Temp Result:%v\n", *results)
		//arr[start] = temp
	} else {
		for i := start; i < end; i++ {

			//temp := arr[start]
			if !(*results)[string(arr)] {
				makeLetterCasePermutation2(arr, start+1, end, results)
			}

			if isLowercase(arr[start]) {
				//Convert to Uppercase
				arr[start] = arr[start] - 32
				if !(*results)[string(arr)] {
					fmt.Printf("Uppercase:%s  -> Start:%d end:%d arr[start]:%s\n", string(arr[start]), start, end, string(arr[start:]))
					makeLetterCasePermutation2(arr, start+1, end, results)
				}

			} else if isUppercase(arr[start]) {
				//Convert to Lowercase
				arr[start] = arr[start] + 32
				if !(*results)[string(arr)] {
					fmt.Printf("LowerCase:%s -> Start:%d end:%d arr[start]:%s\n", string(arr[start]), start, end, string(arr[start:]))
					makeLetterCasePermutation2(arr, start+1, end, results)
				}
			}
			//arr[start] = temp
		}
	}

}

func letterCasePermutation2(s string) []string {
	var results []string
	tempResult := map[string]bool{}
	arr := []rune(s)
	makeLetterCasePermutation2(arr, 0, len(s), &tempResult)
	for key := range tempResult {
		results = append(results, key)
	}

	return results
}

//https://leetcode.com/problems/letter-case-permutation/discuss/2102271/Python-oror-Letter-case-permutation-oror-faster-than-80-solution
var results []string

//Runtime: 9 ms, faster than 80.99% of Go online submissions for Letter Case Permutation.
//Memory Usage: 6.9 MB, less than 16.90% of Go online submissions for Letter Case Permutation.
func makeLetterCasePermutation(s string, sPart string) {
	if len(s) == 0 {
		results = append(results, sPart)
	} else {
		if s[0] >= '0' && s[0] <= '9' {
			makeLetterCasePermutation(s[1:], sPart+string(s[0]))
		} else {
			makeLetterCasePermutation(s[1:], sPart+strings.ToLower(string(s[0])))
			makeLetterCasePermutation(s[1:], sPart+strings.ToUpper(string(s[0])))
		}
	}
}

func letterCasePermutation(s string) []string {
	results = []string{}
	makeLetterCasePermutation(s, "")
	return results
}

func mainLetterCasePermutation() {
	s := "a1b2"
	//permutation := letterCasePermutation(s)
	//fmt.Printf("%v", permutation)

	s = "L2tcnhpZK"
	//s = "abc"
	results = letterCasePermutation(s)
	fmt.Printf("%v", results)
}
