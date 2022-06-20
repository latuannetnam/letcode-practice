// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/15/2022
// Letcode problem: 240. Search a 2D Matrix II
// Letcode link: https://leetcode.com/problems/search-a-2d-matrix-ii/
// Level: medium
//Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix.
//This matrix has the following properties:
// - Integers in each row are sorted in ascending from left to right.
// - Integers in each column are sorted in ascending from top to bottom.
package main

import "fmt"

func searchMatrixII_old(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Printf("Rows:%d cols:%d\n", rows, cols)
	//Target is out of range
	if matrix[0][0] > target || matrix[rows-1][cols-1] < target {
		return false
	}
	maxCol := 0
	maxRow := 0

	for maxCol < cols && maxRow < rows {
		fmt.Printf("Row:%d maxCol:%d item:%d -> %d \n", maxRow, maxCol, matrix[maxRow][maxCol], target)
		if matrix[maxRow][maxCol] == target {
			return true
		} else if matrix[maxRow][maxCol] > target {
			break
		}
		if maxRow+1 < rows {
			maxRow++
		}
		if maxCol+1 < cols {
			maxCol++
		}

	}

	fmt.Printf("Target from maxRow:0-%d maxCol:0-%d\n", maxRow, maxCol)
	result := -1

	for row := 0; result < 0 && row <= maxRow; row++ {
		fmt.Printf("Check Row:%d matrix[%d][%d]:%d\n", row, row, maxCol, matrix[row][maxCol])
		if matrix[row][maxCol] >= target {
			result = binarySearch(matrix[row], target, 0, maxCol)
		}
	}

	return result >= 0
}

//Runtime: 41 ms, faster than 23.21% of Go online submissions for Search a 2D Matrix II.
//Memory Usage: 7.5 MB, less than 18.77% of Go online submissions for Search a 2D Matrix II.
func searchMatrixII_old2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Printf("Rows:%d cols:%d\n", rows, cols)
	//Target is out of range
	if matrix[0][0] > target || matrix[rows-1][cols-1] < target {
		return false
	}
	fmt.Println("---Find min of row to search---")
	minRow := 0
	lo := 0
	high := rows - 1
	for lo <= high {
		mid := lo + (high-lo)/2
		fmt.Printf("Row lo:%d mid:%d high:%d matrix[%d][cols]:%d\n", lo, mid, high, mid, matrix[mid][cols-1])
		if matrix[mid][cols-1] == target {
			return true
		}
		if matrix[mid][cols-1] > target {
			high = mid - 1
			minRow = mid
		} else {
			lo = mid + 1
		}
	}

	fmt.Println("---Find max of row to search---")
	maxRow := -1
	lo = 0
	high = rows - 1
	for lo <= high {
		mid := lo + (high-lo)/2
		fmt.Printf("Row lo:%d mid:%d high:%d matrix[%d][0]:%d\n", lo, mid, high, mid, matrix[mid][0])
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			high = mid - 1
			maxRow = high
		} else {
			lo = mid + 1
		}
	}

	if maxRow < 0 {
		maxRow = rows - 1
	}

	fmt.Println("---Find min of col to search---")
	minCol := 0
	lo = 0
	high = cols - 1
	for lo <= high {
		mid := lo + (high-lo)/2
		fmt.Printf("Row lo:%d mid:%d high:%d matrix[rows][%d]:%d\n", lo, mid, high, mid, matrix[rows-1][mid])
		if matrix[rows-1][mid] == target {
			return true
		}
		if matrix[rows-1][mid] > target {
			high = mid - 1
			minCol = mid
		} else {
			lo = mid + 1
		}
	}
	fmt.Println("---Find max of col to search---")
	maxCol := -1
	lo = 0
	high = cols - 1
	for lo <= high {
		mid := lo + (high-lo)/2
		fmt.Printf("Row lo:%d mid:%d high:%d matrix[0][%d]:%d\n", lo, mid, high, mid, matrix[0][mid])
		if matrix[0][mid] == target {
			return true
		}
		if matrix[0][mid] > target {
			high = mid - 1
			maxCol = high
		} else {
			lo = mid + 1
		}
	}

	if maxCol < 0 {
		maxCol = cols - 1
	}

	fmt.Printf("Row from:%d - %d. Col from %d - %d \n", minRow, maxRow, minCol, maxCol)
	for row := minRow; row <= maxRow; row++ {
		fmt.Printf("Search target:%d in row:%d from col:%d-%d\n", target, row, minCol, maxCol)
		if binarySearch(matrix[row], target, minCol, maxCol) >= 0 {
			return true
		}
	}
	return false
}

//https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/2168010/Search-a-2D-Matrix-II-oror-90-faster-solution-C%2B%2B
//Runtime: 34 ms, faster than 40.96% of Go online submissions for Search a 2D Matrix II.
//Memory Usage: 7.1 MB, less than 66.21% of Go online submissions for Search a 2D Matrix II.
func searchMatrixII(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Printf("Rows:%d cols:%d\n", rows, cols)
	//Target is out of range
	if matrix[0][0] > target || matrix[rows-1][cols-1] < target {
		return false
	}
	row := 0
	col := cols - 1
	for row < rows && col >= 0 {
		fmt.Printf("row:%d col:%d item:%d\n", row, col, matrix[row][col])
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
func mainSearchMatrixII() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//matrix := [][]int{{-1, 3}}
	//matrix := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}

	target := 6
	fmt.Printf("Matrix:%v target:%d\n", matrix, target)
	fmt.Printf("Target:%d found:%t\n", target, searchMatrixII(matrix, target))
}
