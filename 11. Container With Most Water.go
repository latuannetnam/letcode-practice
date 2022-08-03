// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 11. Container With Most Water
// Letcode link: https://leetcode.com/problems/container-with-most-water/
// Level: medium
// Topics: Array, Two pointers
// Problem detail:
// You are given an integer array height of length n.
// There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

package main

import "fmt"

// Runtime: 136 ms, faster than 45.32% of Go online submissions for Container With Most Water.
// Memory Usage: 9.3 MB, less than 40.11% of Go online submissions for Container With Most Water.
func maxAreaWaterContainer(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		fmt.Printf("Area [%d][%d]:%d->%d\n", left, right, area, maxArea)
		if area > maxArea {
			maxArea = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func mainMaxAreaWaterContainer() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height = []int{1, 1}
	fmt.Printf("height:%v\n", height)
	fmt.Printf("max water:%d\n", maxAreaWaterContainer(height))

}
