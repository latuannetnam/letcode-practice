// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 278. First Bad Version
// Letcode problem: https://leetcode.com/problems/first-bad-version/
// Level: Easy
// You are a product manager and currently leading a team to develop a new product.
// Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version, all the versions after a bad version are also bad.

// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

// You are given an API bool isBadVersion(version) which returns whether version is bad.
// Implement a function to find the first bad version. You should minimize the number of calls to the API.

package main

import "fmt"

func isBadVersion(version int) bool {
	is_bad := false
	if version >= 2 {
		is_bad = true
	}
	return is_bad
}

func firstBadVersion(n int) int {
	left := 1
	right := n
	var pos int = left
	bad := pos
	for left <= right {
		pos = left + (right-left)/2
		fmt.Printf("left: %d right:%d pos:%d isBad:%t\n", left, right, pos, isBadVersion(pos))
		if isBadVersion(pos) {
			bad = pos
			right = pos - 1
		} else {
			left = pos + 1
		}
	}

	return bad
}

func mainFirstBadVersion() {
	n := 3
	bad := firstBadVersion(n)
	fmt.Printf("n:%d bad:%d\n", n, bad)

}
