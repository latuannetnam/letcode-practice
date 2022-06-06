// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/6/2022
// Letcode problem: 994. Rotting Oranges
// Letcode link: https://leetcode.com/problems/rotting-oranges/
// Level: medium
//You are given an m x n grid where each cell can have one of three values:
//
//0 representing an empty cell,
//1 representing a fresh orange, or
//2 representing a rotten orange.
//Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
//
//Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

package main

import "fmt"

const FRESH_ORANGE = 1
const ROTTEN_ORANGE = 2

type MatrixCell struct {
	row int
	col int
}

func makeOrangeFresh2Rotten(grid [][]int, rottenOranges []MatrixCell) []MatrixCell {
	fmt.Printf("Make fresh oranges to rotten. Rotten:%v \n", rottenOranges)
	tempOranges := []MatrixCell{}
	for _, rottenOrange := range rottenOranges {
		row := rottenOrange.row
		col := rottenOrange.col
		fmt.Printf("Turn adjacent of cell[%d][%d] to rotten\n", row, col)
		if row-1 >= 0 && grid[row-1][col] == FRESH_ORANGE {
			fmt.Printf("Turn cell[%d][%d] to rotten\n", row-1, col)
			grid[row-1][col] = ROTTEN_ORANGE
			tempOranges = append(tempOranges, MatrixCell{row - 1, col})
		}
		if row+1 < len(grid) && grid[row+1][col] == FRESH_ORANGE {
			fmt.Printf("Turn cell[%d][%d] to rotten\n", row+1, col)
			grid[row+1][col] = ROTTEN_ORANGE
			tempOranges = append(tempOranges, MatrixCell{row + 1, col})
		}

		if col-1 >= 0 && grid[row][col-1] == FRESH_ORANGE {
			fmt.Printf("Turn cell[%d][%d] to rotten\n", row, col-1)
			grid[row][col-1] = ROTTEN_ORANGE
			tempOranges = append(tempOranges, MatrixCell{row, col - 1})
		}
		if col+1 < len(grid[row]) && grid[row][col+1] == FRESH_ORANGE {
			fmt.Printf("Turn cell[%d][%d] to rotten\n", row, col+1)
			grid[row][col+1] = ROTTEN_ORANGE
			tempOranges = append(tempOranges, MatrixCell{row, col + 1})
		}
		//	Remove current rotten cell
		grid[row][col] = 0
	}
	//Clean up rotten oranges
	rottenOranges = []MatrixCell{}
	//Refresh rotten oranges
	for _, rottenOrange := range tempOranges {
		row := rottenOrange.row
		col := rottenOrange.col
		if isAdjacentFreshOrange(grid, row, col) {
			rottenOranges = append(rottenOranges, MatrixCell{row, col})
		} else {
			grid[row][col] = 0
		}
		//	Remove orange from fresh list
		//freshOranges = deleteOrange(freshOranges, index)
	}

	return rottenOranges
}

func isAdjacentFreshOrange(grid [][]int, row int, col int) bool {

	//fmt.Printf("Check rotten [%d][%d] with adjacent fresh\n", row, col)
	if row-1 >= 0 && grid[row-1][col] == FRESH_ORANGE {
		return true
	}
	if row+1 < len(grid) && grid[row+1][col] == FRESH_ORANGE {
		return true
	}

	if col-1 >= 0 && grid[row][col-1] == FRESH_ORANGE {
		return true
	}
	if col+1 < len(grid[row]) && grid[row][col+1] == FRESH_ORANGE {
		return true
	}

	return false
}

//Runtime: 2 ms, faster than 85.71% of Go online submissions for Rotting Oranges.
//Memory Usage: 3 MB, less than 48.63% of Go online submissions for Rotting Oranges.
func orangesRotting(grid [][]int) int {
	minTime := 0
	//Find all rotten orange and fresh orange
	totalRows := len(grid)
	totalCols := len(grid[0])
	rottenOranges := []MatrixCell{}
	freshOranges := []MatrixCell{}
	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalCols; col++ {
			if grid[row][col] == FRESH_ORANGE {
				freshOranges = append(freshOranges, MatrixCell{row, col})
			} else if grid[row][col] == ROTTEN_ORANGE {
				//Only add rotten orange if adjacent cells are fresh oranges
				if isAdjacentFreshOrange(grid, row, col) {
					rottenOranges = append(rottenOranges, MatrixCell{row, col})
				} else {
					grid[row][col] = 0
				}

			}
		}
	}

	fmt.Printf("Number of rotten oranges with adjacent fresh:%d\n", len(rottenOranges))
	fmt.Printf("Number of fresh oranges:%d\n", len(freshOranges))

	//Return minTime=0 if no fresh oranges
	if len(freshOranges) == 0 {
		return minTime
	}

	//Return -1 if no fresh oranges can be turned into rotten
	if len(rottenOranges) == 0 {
		return -1
	}

	for len(rottenOranges) > 0 {
		rottenOranges = makeOrangeFresh2Rotten(grid, rottenOranges)
		minTime++
	}

	//Check if any fresh orange left
	for _, orange := range freshOranges {
		if grid[orange.row][orange.col] == FRESH_ORANGE {
			fmt.Printf("Remain fresh cell[%d][%d]\n", orange.row, orange.col)
			return -1
		}
	}

	return minTime
}

func mainOrangesRotting() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Printf("Grid:%v\n", grid)
	minTime := orangesRotting(grid)
	fmt.Printf("Min minute to rotten all:%d\n", minTime)

	fmt.Printf("\n---------------\n")
	grid = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Printf("Grid:%v\n", grid)
	minTime = orangesRotting(grid)
	fmt.Printf("Min minute to rotten all:%d\n", minTime)

	fmt.Printf("\n---------------\n")
	grid = [][]int{{0, 2}}
	fmt.Printf("Grid:%v\n", grid)
	minTime = orangesRotting(grid)
	fmt.Printf("Min minute to rotten all:%d\n", minTime)

	fmt.Printf("\n---------------\n")
	grid = [][]int{{1, 2}}
	fmt.Printf("Grid:%v\n", grid)
	minTime = orangesRotting(grid)
	fmt.Printf("Min minute to rotten all:%d\n", minTime)

}
