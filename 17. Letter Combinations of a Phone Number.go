// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/letter-combinations-of-a-phone-number/
// Level: medium
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.
// A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
// 1 = None
// 2 = abc
// 3 = def
// 4 = ghi
// 5 = jkl
// 6 = mno
// 7 = pqrs
// 8 = tuv
// 9 = wxyz

package main

import (
	"fmt"
	"strings"
)

var digit_2_letters = map[string][]string{
	"1": {""},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var digit_2_letters2 = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func convertDigit2(letters []string, digit_slice []string) []string {
	fmt.Printf(" digits: %v letters: %v \n", digit_slice, letters)
	if len(digit_slice) == 1 {
		letters = digit_2_letters[digit_slice[0]]
	} else {
		letters_final := []string{}
		letters_temp_arr := convertDigit2(letters, digit_slice[1:])
		digit_2_letter_arr := digit_2_letters[string(digit_slice[0])]
		for _, digit_2_letter := range digit_2_letter_arr {
			for _, letters_temp := range letters_temp_arr {
				letters_final = append(letters_final, digit_2_letter+letters_temp)
			}
		}
		letters = letters_final
	}

	fmt.Printf(" digits: %v letters after: %v \n", digit_slice, letters)
	return letters
}

func letterCombinations2(digits string) []string {
	letters := []string{}

	if digits == "" {
		return letters
	}

	digit_slice := strings.Split(digits, "")
	fmt.Printf("Digit: %v \n", digit_slice)
	fmt.Printf("Digit 2 letter: %v \n", digit_2_letters)
	letters = convertDigit2(letters, digit_slice)

	return letters
}

func convertDigit(letters []string, digits string) []string {
	fmt.Printf(" digits: %v letters: %v \n", digits, letters)
	if len(digits) == 1 {
		letters = strings.Split(digit_2_letters2[digits[0]-'0'], "")
	} else {
		letters_final := []string{}
		letters_temp_arr := convertDigit(letters, digits[1:])
		digit_2_letter_arr := strings.Split(digit_2_letters2[digits[0]-'0'], "")
		for _, digit_2_letter := range digit_2_letter_arr {
			for _, letters_temp := range letters_temp_arr {
				letters_final = append(letters_final, digit_2_letter+letters_temp)
			}
		}
		letters = letters_final
	}

	fmt.Printf(" digits: %v letters after: %v \n", digits, letters)
	return letters
}

func letterCombinations(digits string) []string {
	letters := []string{}

	if digits == "" {
		return letters
	}

	letters = convertDigit(letters, digits)

	return letters
}

func mainLetterCombinations() {
	digits := "23"
	letters := letterCombinations(digits)
	fmt.Printf("digit: %s total: %d letters: %v", digits, len(letters), letters)
}
