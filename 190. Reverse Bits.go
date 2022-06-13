// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/13/2022
// Letcode problem: 190. Reverse Bits
// Letcode link: https://leetcode.com/problems/reverse-bits/
// Level: easy
//Reverse bits of a given 32 bits unsigned integer.
package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
)

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Bits.
//Memory Usage: 2.6 MB, less than 7.43% of Go online submissions for Reverse Bits.
func reverseBits(num uint32) uint32 {
	var reverseNum uint32 = 0
	for i := 0; i < 32 && num >= 0; i++ {
		bitNum := num % 2
		if bitNum > 0 {
			reverseNum += uint32(math.Pow(2, float64(31-i)))
		}
		num = num / 2
	}
	return reverseNum
}

//Best, using string
func reverseBits2(num uint32) uint32 {
	str := fmt.Sprintf("%032b", num)
	fmt.Printf("%v\n", str)
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse = reverse + string(str[i])
	}
	fmt.Printf("%v\n", reverse)
	revNum, _ := strconv.ParseUint(reverse, 2, 32)
	fmt.Printf("%v\n", revNum)
	return uint32(revNum)
}

func mainReverseBits() {
	var num uint32 = 4294967293
	reverse1 := reverseBits2(num)
	reverse2 := bits.Reverse32(num)
	fmt.Printf("num:%032b  reverse:%d/%032b reverse2:%d/%032b\n", num, reverse1, reverse1, reverse2, reverse2)
}
