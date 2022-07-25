// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 557. Reverse Words in a String III
// Letcode problem: Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace and initial word order.
// Level: easy
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

package main

import (
	"fmt"
	"strings"
)

// Runtime: 3 ms, faster than 95.24% of Go online submissions for Reverse Words in a String III.
// Memory Usage: 6.6 MB, less than 59.60% of Go online submissions for Reverse Words in a String III.
func reverseWords(s string) string {
	s_list := strings.Split(s, " ")
	new_s := []byte{}
	for _, sub_s := range s_list {
		fmt.Printf("sub s: %s\n", sub_s)
		sub_s_arr := []byte(sub_s)
		reverseStringByte(sub_s_arr)
		if len(new_s) > 0 {
			new_s = append(new_s, ' ')
		}
		new_s = append(new_s, sub_s_arr...)

	}
	return string(new_s)

}

func mainReverseWords() {
	var s string
	s = "Let's take LeetCode contest"
	fmt.Printf("Original string:%s\n", s)
	s = reverseWords(s)
	fmt.Printf("Reverse string:%s\n", s)
}
