// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 7/1/2022
// Letcode problem: 130. Surrounded Regions
// Letcode link: https://leetcode.com/problems/surrounded-regions/
// Level: medium
// Topics: matrix, Deep-first-search, BFS, Union find
//Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally surrounded by 'X'.
//
//A region is captured by flipping all 'O's into 'X's in that surrounded region.

package main

import "fmt"

func isBorderCell(row int, col int, totalRows int, totalCols int) bool {
	if row == 0 || col == 0 || row == totalRows-1 || col == totalCols-1 {
		return true
	}
	return false
}

func solveSurroundedRegions_failed(board [][]byte) {
	totalRows := len(board)
	totalCols := len(board[0])
	//Direction vector
	dVector130 := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	//Visited matrix
	visitedCells := make([][]bool, totalRows)
	for row := range visitedCells {
		visitedCells[row] = make([]bool, totalCols)
	}

	//init queue
	queueMatrix := queueOfArray{}
	queueMatrix = queueMatrix.Push([]int{0, 0})
	var corr []int
	count := 0

	for !queueMatrix.isEmpty() {
		queueMatrix, corr = queueMatrix.Pop()
		row := corr[0]
		col := corr[1]
		canFlipped := false
		count++
		if count >= 20 {
			break
		}
		if !visitedCells[row][col] {
			fmt.Printf("Checking board[%d][%d]:%d\n", row, col, board[row][col])
			//visitedCells[row][col] = true
			if board[row][col] == 'O' && !isBorderCell(row, col, totalRows, totalCols) {
				//Check if cell is not in border
				canFlipped = true
			}

			//	Find adjacent cells
			for i := range dVector130 {
				adjRow := row + dVector130[i][0]
				adjCol := col + dVector130[i][1]
				if adjRow >= 0 && adjCol >= 0 && adjRow < totalRows && adjCol < totalCols && !visitedCells[adjRow][adjCol] {
					queueMatrix = queueMatrix.Push([]int{adjRow, adjCol})
					//	Check if adjacent is not border
					if isBorderCell(adjRow, adjCol, totalRows, totalCols) && board[adjRow][adjCol] == 'O' {
						canFlipped = false
					}
				}
			}

			if canFlipped {
				fmt.Printf("Cell [%d][%d] is flipped\n", row, col)
				board[row][col] = 'X'
			}
			if board[row][col] == 'X' {
				visitedCells[row][col] = true
			}
		}

	}
}

//https://leetcode.com/problems/surrounded-regions/discuss/2206678/Python-85-Simple-Beginner-friendly
// Search for all border cell and connected border cell -> mark=A
// Convert A -> O
// O -> X
//Runtime: 11 ms, faster than 100.00% of Go online submissions for Surrounded Regions.
//Memory Usage: 6.5 MB, less than 84.86% of Go online submissions for Surrounded Regions.
func solveSurroundedRegions(board [][]byte) {
	totalRows := len(board)
	totalCols := len(board[0])
	//Direction vector
	dVector130 := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	//Visited matrix
	visitedCells := make([][]bool, totalRows)
	for row := range visitedCells {
		visitedCells[row] = make([]bool, totalCols)
	}

	//	Search for border cells == 0
	//init queue
	queueMatrix := queueOfArray{}
	fmt.Printf("Search for border cell rows:%d\n", totalRows)
	for row := range board {
		if board[row][0] == 'O' {
			board[row][0] = 'A'
			queueMatrix = queueMatrix.Push([]int{row, 0})
		}
		if board[row][totalCols-1] == 'O' {
			board[row][totalCols-1] = 'A'
			queueMatrix = queueMatrix.Push([]int{row, totalCols - 1})
		}
	}

	fmt.Printf("Search for border cell cols:%d\n", totalCols)
	for col := range board[0] {
		if board[0][col] == 'O' {
			board[0][col] = 'A'
			queueMatrix = queueMatrix.Push([]int{0, col})
		}

		if board[totalRows-1][col] == 'O' {
			board[totalRows-1][col] = 'A'
			queueMatrix = queueMatrix.Push([]int{totalRows - 1, col})
		}
	}

	//	Use BFS to search for all adjacent cells of border cell
	var corr []int
	for !queueMatrix.isEmpty() {
		queueMatrix, corr = queueMatrix.Pop()
		row := corr[0]
		col := corr[1]
		if !visitedCells[row][col] {
			visitedCells[row][col] = true
			fmt.Printf("Checking board[%d][%d]:%d\n", row, col, board[row][col])
			//	Find adjacent cells
			for i := range dVector130 {
				adjRow := row + dVector130[i][0]
				adjCol := col + dVector130[i][1]
				if adjRow >= 0 && adjCol >= 0 && adjRow < totalRows && adjCol < totalCols && !visitedCells[adjRow][adjCol] {
					if board[adjRow][adjCol] == 'O' {
						board[adjRow][adjCol] = 'A'
						queueMatrix = queueMatrix.Push([]int{adjRow, adjCol})
					}

				}

			}
		}
	}

	//	Scan all board can convert A->O, O->X
	for row := range board {
		for col := range board[0] {
			if board[row][col] == 'A' {
				board[row][col] = 'O'
			} else {
				board[row][col] = 'X'
			}
		}
	}

}

func mainSurroundedRegions() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	//board = [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}
	board = [][]byte{{'O', 'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X', 'O'}, {'X', 'O', 'X', 'O', 'X'}, {'O', 'X', 'O', 'O', 'O'}, {'X', 'X', 'O', 'X', 'O'}}
	fmt.Printf("Board:%v\n", board)
	solveSurroundedRegions(board)
	fmt.Printf("Solved board:%v\n", board)
}
