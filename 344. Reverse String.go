// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 344. Reverse String
// Letcode problem: Write a function that reverses a string. The input string is given as an array of characters s.
//  You must do this by modifying the input array in-place with O(1) extra memory.
// Level: easy
// Input: s = ["h","e","l","l","o"]
// Output: ["o","l","l","e","h"]

package main

import "fmt"

// Runtime: 47 ms, faster than 26.09% of Go online submissions for Reverse String.
// Memory Usage: 6.7 MB, less than 57.52% of Go online submissions for Reverse String.
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func mainReverseString() {
	var s []byte
	s = []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("Original string: %v\n", s)
	reverseString(s)
	fmt.Printf("Reverse string: %v\n", s)

	s = []byte{'h', 'e', 'l', 'l'}
	fmt.Printf("Original string: %v\n", s)
	reverseString(s)
	fmt.Printf("Reverse string: %v\n", s)

}
