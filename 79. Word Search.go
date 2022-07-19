// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 79. Word Search
// Letcode link: https://leetcode.com/problems/word-search/
// Level: medium
// Topics: Array, Backtrack, Matrix, DFS
// Problem detail:
//Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells,
// where adjacent cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once.

package main

import "fmt"

func existWordInMatrixFailed(board [][]byte, word string) bool {
	totalRows := len(board)
	totalCols := len(board[0])
	//Direction vector
	// dVector := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	dVector := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	//Visited matrix
	visitedCells := make([][]bool, totalRows)
	for row := range visitedCells {
		visitedCells[row] = make([]bool, totalCols)
	}

	// Init wordsearch
	wordSearch := false
	// Init pointer
	// wordPointer := -1

	// Backtracking function
	var backTracking func(int, int, int)
	isNextChar := func(row, col int, wordPointer int) (int, int) {
		if wordPointer < len(word)-1 {
			for i := range dVector {
				adjRow := row + dVector[i][0]
				adjCol := col + dVector[i][1]
				if adjRow >= 0 && adjRow < totalRows && adjCol >= 0 && adjCol < totalCols && !visitedCells[adjRow][adjCol] {
					if board[adjRow][adjCol] == word[wordPointer+1] {
						return adjRow, adjCol
					}
				}
			}
		}

		return -1, -1

	}

	backTracking = func(row int, col int, wordPointer int) {
		fmt.Printf("Search for word in [%d][%d]:%s->%s pointer:%d\n", row, col, string(board[row][col]), string(word[wordPointer+1]), wordPointer)
		if wordPointer == len(word)-1 {
			wordSearch = true
			return
		}
		visitedCells[row][col] = true
		charFound := false

		if board[row][col] == word[wordPointer+1] {
			wordPointer++
			fmt.Printf("  found char %s in position:%d at [%d][%d]\n", string(board[row][col]), wordPointer, row, col)
			if wordPointer == len(word)-1 {
				wordSearch = true
				return
			}
			adjRow, adjCol := isNextChar(row, col, wordPointer)
			if adjRow >= 0 && adjCol >= 0 {
				charFound = true
				backTracking(adjRow, adjCol, wordPointer)
			} else {
				wordPointer--
			}
		}

		if !charFound {
			for i := range dVector {
				adjRow := row + dVector[i][0]
				adjCol := col + dVector[i][1]
				if adjRow >= 0 && adjRow < totalRows && adjCol >= 0 && adjCol < totalCols && !visitedCells[adjRow][adjCol] {

					backTracking(adjRow, adjCol, wordPointer)

				}
			}

		}

	}

	// Call backTracking for 1st cell
	backTracking(0, 0, -1)

	return wordSearch
}

// https://leetcode.com/problems/word-search/discuss/2296101/Clean-Code-%2B-Heavily-commented-code-for-better-understanding
// Runtime: 227 ms, faster than 33.38% of Go online submissions for Word Search.
// Memory Usage: 2.1 MB, less than 66.18% of Go online submissions for Word Search.
func existWordInMatrix(board [][]byte, word string) bool {
	totalRows := len(board)
	totalCols := len(board[0])
	//Direction vector
	// dVector := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	dVector := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	//Visited matrix
	visitedCells := make([][]bool, totalRows)
	for row := range visitedCells {
		visitedCells[row] = make([]bool, totalCols)
	}

	// Declare backTracking function
	var backTracking func(int, int, int) bool
	backTracking = func(row int, col int, wordPointer int) bool {
		if wordPointer == len(word) {
			return true
		}
		// Check boundary
		if row < 0 || row >= totalRows || col < 0 || col >= totalCols {
			return false
		}

		fmt.Printf("Search for word in [%d][%d]:%s pointer:%d\n", row, col, string(board[row][col]), wordPointer)

		if visitedCells[row][col] || board[row][col] != word[wordPointer] {
			return false
		}

		visitedCells[row][col] = true
		result := false
		for i := range dVector {
			adjRow := row + dVector[i][0]
			adjCol := col + dVector[i][1]
			fmt.Printf(" subSearch:[%d][%d]\n", adjRow, adjCol)
			result = result || backTracking(adjRow, adjCol, wordPointer+1)
		}
		fmt.Printf("  subResult[%d][%d]:%t\n", row, col, result)

		// Backtrack
		visitedCells[row][col] = false
		return result
	}

	//------------------------------------- main WordSearch
	for row := range board {
		for col := range board[0] {
			fmt.Printf("Search for first word in [%d][%d]:%s -> %s => %t\n", row, col, string(board[row][col]), string(word[0]), board[row][col] == word[0])
			if board[row][col] == word[0] {
				if backTracking(row, col, 0) {
					return true
				}
			}
		}
	}
	return false
}

func mainExistWordInMatrix() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	// word = "ABFS"
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word = "SEE"

	// board = [][]byte{{'a'}}
	// word = "b"

	board = [][]byte{{'a', 'b'}}
	word = "ba"

	board = [][]byte{{'a'}}
	word = "a"

	fmt.Printf("Board:%v\n", board)
	fmt.Printf("Word to search:%s\n", word)
	fmt.Printf("Word search result:%t\n", existWordInMatrix(board, word))
}
