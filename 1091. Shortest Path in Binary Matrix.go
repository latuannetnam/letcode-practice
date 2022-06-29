// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/28/2022
// Letcode problem: 1091. Shortest Path in Binary Matrix
// Letcode link: https://leetcode.com/problems/shortest-path-in-binary-matrix/
// Level: medium
// Topics: Matrix, Breath-First-Search
//Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix. If there is no clear path, return -1.
//A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell (i.e., (n - 1, n - 1)) such that:
//	All the visited cells of the path are 0.
//	All the adjacent cells of the path are 8-directionally connected (i.e., they are different and they share an edge or a corner).
//The length of a clear path is the number of visited cells of this path.
package main

import (
	"fmt"
	"math"
)

// Visit direction:  D+L,    D+R,    U+L,     U+R ,     D,      L ,     R,        U
var dVector [][]int = [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {1, 0}, {0, 1}, {0, -1}, {-1, 0}}

func isIdxValidMatrix01(grid [][]int, row, col, totalRows, totalCols int) bool {
	result := true

	if row < 0 || col < 0 || row >= totalRows || col >= totalCols {
		result = false
	} else if grid[row][col] == 1 {
		result = false
	}
	//fmt.Printf("Check valid row:%d col:%d valid:%t\n", row, col, result)
	return result
}

func isValidAdjDistance(distant [][]int, row, col, totalRows, totalCols int) bool {
	result := true

	if row < 0 || col < 0 || row >= totalRows || col >= totalCols {
		result = false
	} else if distant[row][col] == 0 {
		result = false
	}
	//fmt.Printf("Check valid row:%d col:%d valid:%t\n", row, col, result)
	return result
}

func calShortestDistance(distant [][]int, row int, col int, totalRows, totalCols int) int {
	result := math.MaxInt
	if row == 0 && col == 0 {
		return distant[0][0]
	}
	for i := 0; i < len(dVector); i++ {
		idxRow := row + dVector[i][0]
		idxCol := col + dVector[i][1]
		if isValidAdjDistance(distant, idxRow, idxCol, totalRows, totalCols) {
			result = min(result, distant[idxRow][idxCol]+1)
			//result = distant[idxRow][idxCol] + 1
			//break
		}
	}
	return result
}

//https://www.geeksforgeeks.org/breadth-first-traversal-bfs-on-a-2d-array/
//Runtime: 127 ms, faster than 17.28% of Go online submissions for Shortest Path in Binary Matrix.
//Memory Usage: 8.4 MB, less than 20.43% of Go online submissions for Shortest Path in Binary Matrix.
func shortestPathBinaryMatrix(grid [][]int) int {
	totalRows := len(grid)
	totalCols := len(grid[0])
	//Check top cell and bottom cell
	if grid[0][0] != 0 || grid[totalRows-1][totalCols-1] != 0 {
		return -1
	}

	//distant array
	distant := make([][]int, totalRows)
	for row := range distant {
		distant[row] = make([]int, totalCols)
	}

	distant[0][0] = 1

	queueMatrix := queueOfArray{}
	//init stack
	queueMatrix = queueMatrix.Push([]int{0, 0})
	var corr []int
	for !queueMatrix.isEmpty() {
		queueMatrix, corr = queueMatrix.Pop()
		row := corr[0]
		col := corr[1]
		if grid[row][col] == 0 {
			distant[row][col] = calShortestDistance(distant, row, col, totalCols, totalCols)
			for i := 0; i < len(dVector); i++ {
				idxRow := row + dVector[i][0]
				idxCol := col + dVector[i][1]
				if isIdxValidMatrix01(grid, idxRow, idxCol, totalRows, totalCols) {
					queueMatrix = queueMatrix.Push([]int{idxRow, idxCol})
				}
			}
		}

		fmt.Printf("row:%d col:%d value:%d distant:%d\n", row, col, grid[row][col], distant[row][col])
		grid[row][col] = 1
		if row == totalRows-1 && col == totalCols-1 {
			return distant[row][col]
		}

	}

	return -1
}

func mainShortestPathBinaryMatrix() {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	//grid = [][]int{{0, 1}, {1, 0}}
	//grid = [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	//Correct = 4
	grid = [][]int{{0, 1, 0, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0}}
	fmt.Printf("grid:%v\n", grid)
	fmt.Printf("Shortest path:%d\n", shortestPathBinaryMatrix(grid))
}
