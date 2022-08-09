// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 36. Valid Sudoku
// Letcode link: https://leetcode.com/problems/valid-sudoku/
// Level: medium
// Topics: Array, Hash, Matrix
// Problem detail:
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
//    - A Sudoku board (partially filled) could be valid but is not necessarily solvable.
//    - Only the filled cells need to be validated according to the mentioned rules.

package main

import (
	"fmt"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
// Memory Usage: 3.5 MB, less than 36.75% of Go online submissions for Valid Sudoku.
func isValidSudoku(board [][]byte) bool {
	// Check row is valid
	for i := range board {
		fmt.Printf("Check row[%d]\n", i)
		numberMap := make([]int, 10)
		for j := range board[i] {
			index := board[i][j] - '0'
			if index <= 9 {
				numberMap[index]++
				if numberMap[index] > 1 {
					return false
				}
			}

		}
	}

	// Check col is valid
	for j := range board {
		fmt.Printf("Check col[%d]\n", j)
		numberMap := make([]int, 10)
		for i := range board[j] {
			index := board[i][j] - '0'
			if index <= 9 {
				numberMap[index]++
				if numberMap[index] > 1 {
					return false
				}
			}
		}
	}

	// Check sub board
	isSubBoardValid := func(boardRow int, boardCol int) bool {
		boardRow = boardRow * 3
		boardCol = boardCol * 3
		fmt.Printf("Check subboard[%d][%d]\n", boardRow, boardCol)
		numberMap := make([]int, 10)
		for i := boardRow; i < boardRow+3; i++ {
			// fmt.Printf(" check row[%d]\n", i)
			for j := boardCol; j < boardCol+3; j++ {
				index := board[i][j] - '0'
				if index <= 9 {
					numberMap[index]++
					if numberMap[index] > 1 {
						return false
					}
				}
			}
		}

		return true
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isSubBoardValid(i, j) {
				return false
			}
		}
	}

	return true
}

func mainValidSudoku() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board = [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board = [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Printf("board:%v\n", board)
	fmt.Printf("is valid Sudoku board:%t\n", isValidSudoku(board))
}
