// Author: Le Anh Tuan (latuannetnam@gmail.com)
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

import "fmt"

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

func generateParenthesis(n int) []string {
	results := generateParenthesisRecusive(n)
	final_result := []string{}
	for item := range results {
		if isParenthesis(item) {
			final_result = append(final_result, item)
		}
	}
	return final_result
}

func mainGenerateParenthesis() {

	var n int = 0
	n = 2
	results := generateParenthesis(n)
	fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	n = 3
	results = generateParenthesis(n)
	fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	// n = 4
	// results = generateParenthesis(n)
	// fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	// n = 5
	// results = generateParenthesis(n)
	// fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

	n = 8
	results = generateParenthesis(n)
	fmt.Printf("n:%d total:%d result:%v\n", n, len(results), results)

}
