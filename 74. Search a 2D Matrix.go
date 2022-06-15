// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/15/2022
// Letcode problem: 74. Search a 2D Matrix
// Letcode link: https://leetcode.com/problems/search-a-2d-matrix/
// Level: medium
//Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
//This matrix has the following properties:
// +Integers in each row are sorted from left to right.
// +The first integer of each row is greater than the last integer of the previous row.
package main

import "fmt"

//Time: O(n)
//Runtime: 6 ms, faster than 28.78% of Go online submissions for Search a 2D Matrix.
//Memory Usage: 2.7 MB, less than 68.33% of Go online submissions for Search a 2D Matrix.
func searchMatrix2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Printf("Rows:%d cols:%d\n", rows, cols)
	//Find row to search for target
	row := -1
	for i := 0; i < rows; i++ {
		if matrix[i][0] == target {
			return true
		} else if matrix[i][cols-1] == target {
			return true
		} else if matrix[i][0] < target && matrix[i][cols-1] > target {
			row = i
			break
		}
	}
	fmt.Printf("Found row contains target:%d\n", row)
	if row < 0 {
		return false
	}

	return binarySearch(matrix[row], target, 0, cols-1) >= 0
}

//Runtime: 3 ms, faster than 70.96% of Go online submissions for Search a 2D Matrix.
//Memory Usage: 2.7 MB, less than 68.33% of Go online submissions for Search a 2D Matrix.
func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Printf("Rows:%d cols:%d\n", rows, cols)
	//Target is out of range
	if matrix[0][0] > target || matrix[rows-1][cols-1] < target {
		return false
	}

	row := -1
	left := 0
	right := rows - 1

	//Find row to search for target
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid][0] == target || matrix[mid][cols-1] == target {
			return true
		}
		fmt.Printf("left:%d mid:%d right:%d:%d\n", left, mid, right, matrix[mid][0])
		if matrix[mid][0] < target && matrix[mid][cols-1] > target {
			row = mid
			fmt.Printf("Found row contains target:%d\n", row)
			break
		}
		if matrix[mid][0] > target {
			right = mid - 1
		} else {
			left = left + 1
		}

		fmt.Printf("New left:%d right:%d\n", left, right)
	}

	if row < 0 {
		return false
	}

	return binarySearch(matrix[row], target, 0, cols-1) >= 0

}
func mainSearchMatrix() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Printf("Matrix:%v target:%d\n", matrix, target)
	fmt.Printf("Target:%d found:%t\n", target, searchMatrix(matrix, target))

}
