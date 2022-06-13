// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/13/2022
// Letcode problem: 191. Number of 1 Bits
// Letcode link: https://leetcode.com/problems/number-of-1-bits/
// Level: easy
//Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
package main

import "fmt"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of 1 Bits.
//Memory Usage: 1.9 MB, less than 42.22% of Go online submissions for Number of 1 Bits.
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num % 2)
		num = num / 2
	}
	return count
}

func mainHammingWeight() {
	var num uint32 = 11
	fmt.Printf("num:%b  -  num of bits:%d", num, hammingWeight(num))
}
