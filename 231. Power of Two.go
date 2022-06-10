// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/10/2022
// Letcode problem: 231. Power of Two
// Letcode link: https://leetcode.com/problems/power-of-two/
// Level: easy
//Given an integer n, return true if it is a power of two. Otherwise, return false.//
//An integer n is a power of two, if there exists an integer x such that n == 2x.
package main

import "fmt"

//Recursion
//Runtime: 3 ms, faster than 54.55% of Go online submissions for Power of Two.
//Memory Usage: 2.2 MB, less than 7.16% of Go online submissions for Power of Two
func isPowerOfTwo2(n int) bool {
	fmt.Printf("n:%d n mod 2:%d\n", n, n%2)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%2 == 0 {
		return isPowerOfTwo(n / 2)
	}
	return false
}

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Power of Two.
//Memory Usage: 2 MB, less than 100.00% of Go online submissions for Power of Two.
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n%2 == 0 {
		fmt.Printf("n:%d mod 2:%d\n", n, n%2)
		n = n / 2
	}
	if n == 1 {
		return true
	}
	return false
}

func mainIsPowerOfTwo() {
	n := 16
	result := isPowerOfTwo(n)
	fmt.Printf("N:%d isPower2:%t\n", n, result)
}
