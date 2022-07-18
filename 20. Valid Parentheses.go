// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/8/2022
// Letcode problem: 20. Valid Parentheses
// Letcode link: https://leetcode.com/problems/valid-parentheses/
// Level: easy
// Topics: String, Stack
//Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//An input string is valid if:
//  * Open brackets must be closed by the same type of brackets.
//  * Open brackets must be closed in the correct order.
package main

import (
	"fmt"
	"math"
)

func validParentheses2(s string) bool {
	var balance float64 = 0
	for i := range s {
		if s[i] == '(' {
			balance += math.Sqrt(2)
		} else if s[i] == ')' {
			balance -= math.Sqrt(2)
		} else if s[i] == '[' {
			balance += math.Sqrt(3)
		} else if s[i] == ']' {
			balance -= math.Sqrt(3)
		} else if s[i] == '{' {
			balance += math.Sqrt(4)
		} else if s[i] == '}' {
			balance -= math.Sqrt(4)
		}
		fmt.Printf("s[%d]:%s balance:%f\n", i, string(s[i]), balance)
		if balance < 0 {
			break
		}

	}
	fmt.Printf("Balance:%f\n", balance*1000000)
	return balance == 0
}

func validParentheses3(s string) bool {
	balance := make([]int, 3)
	for i := range s {
		if s[i] == '(' {
			balance[0]++
		} else if s[i] == ')' {
			balance[0]--
			if i > 0 && s[i-1] != '(' {
				//	invalid close
				if balance[1] != 0 && balance[2] != 0 {
					return false
				}
			}
		} else if s[i] == '[' {
			balance[1]++
		} else if s[i] == ']' {
			balance[1]--
			if i > 0 && s[i-1] != '[' {
				//	invalid close
				if balance[0] != 0 && balance[2] != 0 {
					return false
				}

			}
		} else if s[i] == '{' {
			balance[2]++
		} else if s[i] == '}' {
			balance[2]--
			if balance[2] < 0 {
				return false
			}
			if i > 0 && s[i-1] != '{' {
				//	invalid close
				if balance[0] != 0 && balance[1] != 0 {
					return false
				}
			}
		}
	}
	for i := range balance {
		if balance[i] != 0 {
			return false
		}
	}
	return true
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 47.40% of Go online submissions for Valid Parentheses.
func validParentheses(s string) bool {
	// Map of parentheses
	openParentheseMap := make(map[string]int)
	closeParentheseMap := make(map[string]int)
	openParentheseMap["("] = 1
	openParentheseMap["["] = 2
	openParentheseMap["{"] = 3
	closeParentheseMap[")"] = 1
	closeParentheseMap["]"] = 2
	closeParentheseMap["}"] = 3

	//Array  to keep track of open/close
	// balance[0][] = '('
	// balance[1][] = '['
	// balance[2][0] = '{'
	balance := make([][]int, 3)
	for i := range s {
		openIndex := openParentheseMap[string(s[i])] - 1
		closeIndex := closeParentheseMap[string(s[i])] - 1
		fmt.Printf("s[%d]:%s  openIndex:%d closeIndex:%d last balance:%v\n", i, string(s[i]), openIndex, closeIndex, balance)
		if openIndex >= 0 {
			// Update balance
			balance[openIndex] = append(balance[openIndex], i)
		} else {
			// Check close ) is balanced with (
			if len(balance[closeIndex]) == 0 {
				return false
			}
			// Check close ) is in order
			lastOpen := make([]int, 3)
			for i := range lastOpen {
				if len(balance[i]) == 0 {
					lastOpen[i] = -1
				} else {
					lastOpen[i] = balance[i][len(balance[i])-1]
				}
			}

			fmt.Printf(" lastOpen:%v\n", lastOpen)
			for i := range lastOpen {
				if i != closeIndex && lastOpen[i] > lastOpen[closeIndex] {
					return false
				}
			}
			// update balance
			balance[closeIndex] = balance[closeIndex][:len(balance[closeIndex])-1]

		}
	}

	// Check balance
	fmt.Printf("Remain balance:%v\n", balance)
	for i := range balance {
		if len(balance[i]) > 0 {
			return false
		}
	}

	return true
}

func mainValidParentheses() {
	s := "()[]{}"
	// s = "([)]"
	// s = "([])"
	// s = c
	// s = "(())"
	fmt.Printf("s:%s\n", s)
	fmt.Printf("isValid:%t\n", validParentheses(s))
}
