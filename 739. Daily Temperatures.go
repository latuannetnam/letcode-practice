// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 739. Daily Temperatures
// Letcode link: https://leetcode.com/problems/daily-temperatures/
// Level: medium
// Topics: Array, Stack
// Problem detail:
// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
// If there is no future day for which this is possible, keep answer[i] == 0 instead.

package main

import "fmt"

// Runtime: 247 ms, faster than 51.75% of Go online submissions for Daily Temperatures.
// Memory Usage: 9.8 MB, less than 56.11% of Go online submissions for Daily Temperatures.
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	answer := make([]int, n)
	answer[n-1] = 0
	if n == 1 {
		return answer
	}

	var tempStack stackOfInt
	tempStack = tempStack.Push(n - 1)
	for i := n - 2; i >= 0; i-- {
		fmt.Printf("Find answer for [%d]:%d\n", i, temperatures[i])
		for !tempStack.isEmpty() && temperatures[tempStack.Bottom()] <= temperatures[i] {
			tempStack, _ = tempStack.Pop()
		}
		if tempStack.isEmpty() {
			answer[i] = 0
		} else {
			answer[i] = tempStack.Bottom() - i
		}
		tempStack = tempStack.Push(i)
		fmt.Printf(" stack:%v\n", tempStack)
		fmt.Printf(" answer for [%d]:%d->%d\n", i, temperatures[i], answer[i])

	}

	return answer
}

func mainDailyTemperatures() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	temperatures = []int{30, 40, 50, 60}

	fmt.Printf("Temperatures:%v\n", temperatures)
	fmt.Printf("Answer:%v\n", dailyTemperatures(temperatures))
}
