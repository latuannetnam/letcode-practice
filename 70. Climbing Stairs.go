// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/8/2022
// Letcode problem: 70. Climbing Stairs
// Letcode link: https://leetcode.com/problems/climbing-stairs/
// Level: easy
//You are climbing a staircase. It takes n steps to reach the top.
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
package main

import "fmt"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
//Memory Usage: 1.9 MB, less than 86.60% of Go online submissions for Climbing Stairs.
func climbStairs(n int) int {
	if n <= 2 {
		return n
	} else {
		results := make([]int, n)
		results[0] = 1
		results[1] = 2
		for i := 2; i < n; i++ {
			results[i] = results[i-1] + results[i-2]
		}
		return results[n-1]
	}

}

func mainClimbStairs() {
	n := 3
	result := climbStairs(n)
	fmt.Printf("N:%d number of ways:%d\n", n, result)

}
