// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Problem:22. Generate Parentheses
// Letcode problem: https://leetcode.com/problems/generate-parentheses/
// Level: medium
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]

// Input: n = 1
// Output: ["()"]

// Input: n = 4
// [(())(()) ()()()() (()()()) ()(()()) (()())() ((()())) ()()(()) ()(())() (()(())) (())()() ((())()) ()((())) ((()))() (((())))]

package main

import (
	"fmt"
	"sort"
	"strings"
)

func generateParenthesisRecusive2(n int) []string {
	results := []string{}
	result_map := map[string]bool{}
	var str string = ""
	if n == 1 {
		return []string{"()", ")("}
	} else {
		results_1 := generateParenthesisRecusive2(n - 1)
		for _, item := range results_1 {
			str = "(" + item + ")"
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = "(" + item + "("
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = ")" + item + "("
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = ")" + item + ")"
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = "()" + item
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = "((" + item
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = item + "()"
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}

			str = item + "))"
			if !result_map[str] {
				results = append(results, str)
				result_map[str] = true
			}
		}
	}
	return results
}

func generateParenthesisRecusive(n int) map[string]bool {
	result_map := map[string]bool{}
	var str string = ""
	if n == 1 {
		return map[string]bool{"()": true, ")(": true}
	} else {
		results_1 := generateParenthesisRecusive(n - 1)
		for item := range results_1 {
			str = "(" + item + ")"
			result_map[str] = true

			str = "(" + item + "("
			result_map[str] = true

			str = ")" + item + "("
			result_map[str] = true

			str = ")" + item + ")"
			result_map[str] = true

			str = "()" + item
			result_map[str] = true

			str = "((" + item
			result_map[str] = true

			str = item + "()"
			result_map[str] = true

			str = item + "))"
			result_map[str] = true

		}
	}
	return result_map
}

func isParenthesis(pattern string) bool {
	sum := 0
	for index := range pattern {
		if pattern[index] == '(' {
			sum += 1
		} else {
			sum += -1
		}
		if sum < 0 {
			return false
		}
	}
	return sum == 0
}

func isParenthesisArray(pattern []string) bool {
	sum := 0
	for index := range pattern {
		if pattern[index] == "(" {
			sum += 1
		} else {
			sum += -1
		}
		if sum < 0 {
			return false
		}
	}
	return sum == 0
}

func arrayOfStringToString(pairs []string, s string) string {
	return strings.Join(pairs, "")
}

func generateParenthesisold(n int) []string {
	results := generateParenthesisRecusive(n)
	final_result := []string{}
	for item := range results {
		if isParenthesis(item) {
			final_result = append(final_result, item)
		}
	}
	return final_result
}

//Backtrack 1
//Runtime: 8 ms, faster than 10.56% of Go online submissions for Generate Parentheses.
//Memory Usage: 2.7 MB, less than 74.45% of Go online submissions for Generate Parentheses
func generateParenthesisBacktrack1(n int) []string {
	var result []string
	//Generate pairs of parenthesis
	var parenthesisPairs []string
	for i := 0; i < n; i++ {
		parenthesisPairs = append(parenthesisPairs, "(", ")")
	}
	sort.Strings(parenthesisPairs)

	var backTracking func(int)
	backTracking = func(start int) {
		if start == len(parenthesisPairs) {
			if isParenthesisArray(parenthesisPairs) {
				parenthesis := arrayOfStringToString(parenthesisPairs, "")
				fmt.Printf("Found parenthesis:%s\n", parenthesis)
				result = append(result, parenthesis)
			}
			return
		}

		for i := start; i < len(parenthesisPairs); i++ {
			if i != start && parenthesisPairs[i-1] == parenthesisPairs[i] {
				//Skip duplicate
				continue
			}
			//	Swap position
			parenthesisPairs[start], parenthesisPairs[i] = parenthesisPairs[i], parenthesisPairs[start]
			fmt.Printf("Start[%d]:%s i[%d]:%s pairs:%v\n", start, parenthesisPairs[start], i, parenthesisPairs[i], parenthesisPairs)
			//	Backtrack
			backTracking(start + 1)
			parenthesisPairs[start], parenthesisPairs[i] = parenthesisPairs[i], parenthesisPairs[start]

		}
	}
	fmt.Printf("n:%d parenthesisPairs:%v\n", n, parenthesisPairs)
	backTracking(0)
	return result
}

//Backtrack 2
//https://leetcode.com/problems/generate-parentheses/solution/
func generateParenthesis(n int) []string {
	var result []string
	//Generate pairs of parenthesis

	var backTracking func(int, int, int, []string)
	backTracking = func(open int, close int, max int, combination []string) {
		if len(combination) == 2*max {
			parenthesis := arrayOfStringToString(combination, "")
			fmt.Printf("Found parenthesis:%s\n", parenthesis)
			result = append(result, parenthesis)
		}
		if open < max {
			combination = append(combination, "(")
			backTracking(open+1, close, max, combination)
			//Backtrack
			combination = combination[:len(combination)-1]
		}

		if close < open {
			combination = append(combination, ")")
			backTracking(open, close+1, max, combination)
			//Backtrack
			combination = combination[:len(combination)-1]
		}
	}

	backTracking(0, 0, n, []string{})
	return result
}

func mainGenerateParenthesis() {

	var n int = 0
	//n = 2
	//results := generateParenthesis(n)
	//fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	n = 3
	results = generateParenthesis(n)
	fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	//n = 4
	//results = generateParenthesis(n)
	//fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	// n = 5
	// results = generateParenthesis(n)
	// fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	//n = 8
	//results = generateParenthesis(n)
	//fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

}
