// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 150. Evaluate Reverse Polish Notation
// Letcode link: https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Level: medium
// Topics: Array, Math, Stack
// Problem detail:
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
// Note that division between two integers should truncate toward zero.
// It is guaranteed that the given RPN expression is always valid.
// That means the expression would always evaluate to a result, and there will not be any division by zero operation.

package main

import (
	"fmt"
	"strconv"
)

// Runtime: 5 ms, faster than 72.25% of Go online submissions for Evaluate Reverse Polish Notation.
// Memory Usage: 4.4 MB, less than 31.14% of Go online submissions for Evaluate Reverse Polish Notation.
func evalRPN(tokens []string) int {
	var rpnStack stackOfString
	var result int
	if len(tokens) == 1 {
		result, _ = strconv.Atoi(tokens[0])
		return result
	}
	for i := range tokens {
		fmt.Printf("Token:%s\n", tokens[i])
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			rpnStack = rpnStack.Push(tokens[i])
		} else {
			var token1s, token2s string
			rpnStack, token2s = rpnStack.Pop()
			rpnStack, token1s = rpnStack.Pop()
			token1, _ := strconv.Atoi(token1s)
			token2, _ := strconv.Atoi(token2s)

			if tokens[i] == "+" {
				result = token1 + token2
			} else if tokens[i] == "*" {
				result = token1 * token2
			} else if tokens[i] == "-" {
				result = token1 - token2
			} else if tokens[i] == "/" {
				result = token1 / token2
			}
			fmt.Printf("%d %s %d = %d\n", token1, tokens[i], token2, result)
			rpnStack = rpnStack.Push(strconv.Itoa(result))
		}

	}

	return result
}

func mainEvalRPN() {
	tokens := []string{"2", "1", "+", "3", "*"}
	tokens = []string{"4", "13", "5", "/", "+"}
	tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	tokens = []string{"18"}

	fmt.Printf("Tokens:%v\n", tokens)
	fmt.Printf("RPN:%d\n", evalRPN(tokens))

}
