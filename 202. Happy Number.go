// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 202. Happy Number
// Letcode link: https://leetcode.com/problems/happy-number/
// Level: easy
// Topics: Math, Two Pointers, Hashtable
// Problem detail:
// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
//  - Starting with any positive integer, replace the number by the sum of the squares of its digits.
//  - Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// - Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

package main

import "fmt"

// Runtime: 1 ms, faster than 49.06% of Go online submissions for Happy Number.
// Memory Usage: 2.2 MB, less than 30.45% of Go online submissions for Happy Number.
func isHappySlow(n int) bool {
	squareDigitMap := make(map[int]bool)
	calcSquareDigit := func(num int) int {
		sum := 0
		for num > 0 {
			remain := num % 10
			sum += remain * remain
			num = num / 10
		}
		return sum
	}

	squareDigitMap[n] = true
	for n != 1 {
		n = calcSquareDigit(n)
		// fmt.Printf("n:%d -> Square digit:%d\n", n, n1)
		fmt.Printf("Square digit:%d\n", n)
		// n = n1
		if squareDigitMap[n] {
			break
		} else {
			squareDigitMap[n] = true
		}
	}

	return n == 1
}

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
// Memory Usage: 2.2 MB, less than 30.45% of Go online submissions for Happy Number.
func isHappyHash(n int) bool {
	squareDigitMap := make(map[int]bool)
	var squareLookup [10]int
	for i := range squareLookup {
		squareLookup[i] = i * i
	}
	calcSquareDigit := func(num int) int {
		sum := 0
		for num > 0 {
			remain := num % 10
			sum += squareLookup[remain]
			num = num / 10
		}
		return sum
	}

	squareDigitMap[n] = true
	for n != 1 {
		n = calcSquareDigit(n)
		// fmt.Printf("n:%d -> Square digit:%d\n", n, n1)
		fmt.Printf("Square digit:%d\n", n)
		// n = n1
		if squareDigitMap[n] {
			break
		} else {
			squareDigitMap[n] = true
		}
	}

	return n == 1
}

//  Floyd's Cycle-Finding Algorithm
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
// Memory Usage: 2 MB, less than 94.66% of Go online submissions for Happy Number.
func isHappy(n int) bool {
	var squareLookup [10]int
	for i := range squareLookup {
		squareLookup[i] = i * i
	}
	calcSquareDigit := func(num int) int {
		sum := 0
		for num > 0 {
			remain := num % 10
			sum += squareLookup[remain]
			num = num / 10
		}
		return sum
	}

	slow := n
	fast := calcSquareDigit(n)
	for fast != 1 && slow != fast {
		slow = calcSquareDigit(slow)
		fast = calcSquareDigit(calcSquareDigit(fast))
		fmt.Printf("slow:%d fast:%d\n", slow, fast)
	}

	return fast == 1
}

func mainIsHappy() {
	n := 19
	n = 2
	n = 8
	// n = 10
	fmt.Printf("num:%d\n", n)
	fmt.Printf("isHappy:%t\n", isHappy(n))
}
