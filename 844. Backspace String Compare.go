// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/21/2022
// Letcode problem: 844. Backspace String Compare
// Letcode link: https://leetcode.com/problems/backspace-string-compare/
// Level: easy
//Given two strings s and t, return true if they are equal when both are typed into empty text editors.
//'#' means a backspace character.
//Note that after backspacing an empty text, the text will continue empty.
package main

import (
	"bytes"
	"fmt"
)

func cleanBackspace(s string) string {
	var result []rune
	for _, char := range s {
		if char != '#' {
			result = append(result, char)
		} else {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		}
	}
	return string(result)
}

func cleanBackspace2(s string) []byte {
	var result []byte
	for _, char := range s {
		if char != '#' {
			result = append(result, byte(char))
		} else {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		}
	}
	return result
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Backspace String Compare.
//Memory Usage: 2.1 MB, less than 38.74% of Go online submissions for Backspace String Compare.
func backspaceCompare(s string, t string) bool {
	s = cleanBackspace(s)
	t = cleanBackspace(t)
	fmt.Printf("After clean s:%s t:%s\n", s, t)
	return s == t
}

//Runtime: 4 ms, faster than 16.32% of Go online submissions for Backspace String Compare.
//Memory Usage: 2 MB, less than 85.20% of Go online submissions for Backspace String Compare.
func backspaceCompare2(s string, t string) bool {
	sArr := cleanBackspace2(s)
	tArr := cleanBackspace2(t)
	fmt.Printf("After clean s:%v t:%v\n", s, t)
	return bytes.Equal(sArr, tArr)
}

func mainBackspaceCompare() {
	s := "ab#c"
	t := "ad#c"
	fmt.Printf("s:'%s' t:'%s'\n", s, t)
	fmt.Printf("Two string identical:%t\b", backspaceCompare(s, t))
}
