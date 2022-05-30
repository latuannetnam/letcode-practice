// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 5/30/2022
// Letcode problem: 200. Number of Islands
// Letcode link: https://leetcode.com/problems/number-of-islands/
// Level: medium
//Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
//An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
//You may assume all four edges of the grid are all surrounded by water.
package main

import "fmt"

const water = '0'
const visitedWater = '3'
const rock = '1'
const visitedRock = '2'

func visitCell(grid [][]byte, row, col int) {
	if grid[row][col] == water {
		grid[row][col] = visitedWater
	} else if grid[row][col] == rock {
		grid[row][col] = visitedRock
	}
}

func unVisitCell(grid [][]byte, row, col int) {
	if grid[row][col] == visitedWater {
		grid[row][col] = water
	} else if grid[row][col] == visitedRock {
		grid[row][col] = rock
	}
}

func isVisited(grid [][]byte, row, col int) bool {
	return grid[row][col] > '1'
}

//Use DFS to scan all related rocks in island (adjancent cells with the same color)
func scanIsland(grid [][]byte, row int, col int) {
	totalRows := len(grid)
	totalCols := len(grid[0])
	fmt.Printf("scanIsland Matrix size:%dx%d\n", totalRows, totalCols)
	stackMatrix := stackOfArray{}
	stackMatrix = stackMatrix.Push([]int{row, col})
	corr := []int{}
	count := 0
	for !stackMatrix.isEmpty() {
		stackMatrix, corr = stackMatrix.Pop()
		row := corr[0]
		col := corr[1]
		count++
		if grid[row][col] == rock {
			fmt.Printf("scanIsland Cell[%d:%d]:%d\n", row, col, grid[row][col])
			//  Push adjacent cells
			if row-1 >= 0 && grid[row-1][col] == rock {
				stackMatrix = stackMatrix.Push([]int{row - 1, col})
			}
			if row+1 < totalRows && grid[row+1][col] == rock {
				stackMatrix = stackMatrix.Push([]int{row + 1, col})
			}

			if col-1 >= 0 && grid[row][col-1] == rock {
				stackMatrix = stackMatrix.Push([]int{row, col - 1})
			}

			if col+1 < totalCols && grid[row][col+1] == rock {
				stackMatrix = stackMatrix.Push([]int{row, col + 1})
			}
		}
		visitCell(grid, row, col)
	}

	fmt.Printf("scanIsland: total items:%d \n", count)
}

func scanIsland2(grid [][]byte, row int, col int) {
	totalRows := len(grid)
	totalCols := len(grid[0])
	fmt.Printf("scanIsland Matrix size:%dx%d\n", totalRows, totalCols)
	stackMatrix := stackOfArray{}
	stackMatrix = stackMatrix.Push([]int{row, col})
	corr := []int{}
	count := 0
	for !stackMatrix.isEmpty() {
		stackMatrix, corr = stackMatrix.Pop()
		row := corr[0]
		col := corr[1]
		count++
		if grid[row][col] == visitedRock {
			fmt.Printf("scanIsland Cell[%d:%d]:%d\n", row, col, grid[row][col])
			//  Push adjacent cells
			if row-1 >= 0 && grid[row-1][col] == visitedRock {
				stackMatrix = stackMatrix.Push([]int{row - 1, col})
			}
			if row+1 < totalRows && grid[row+1][col] == visitedRock {
				stackMatrix = stackMatrix.Push([]int{row + 1, col})
			}

			if col-1 >= 0 && grid[row][col-1] == visitedRock {
				stackMatrix = stackMatrix.Push([]int{row, col - 1})
			}

			if col+1 < totalCols && grid[row][col+1] == visitedRock {
				stackMatrix = stackMatrix.Push([]int{row, col + 1})
			}
		}
		unVisitCell(grid, row, col)
	}

	fmt.Printf("scanIsland: total items:%d \n", count)
}

//Use DFS to scan all items in matrix
//Runtime: 69 ms, faster than 5.52% of Go online submissions for Number of Islands.
//Memory Usage: 14.6 MB, less than 5.04% of Go online submissions for Number of Islands.
func numIslands2(grid [][]byte) int {
	totalRows := len(grid)
	totalCols := len(grid[0])
	fmt.Printf("numIslands Matrix size:%dx%d\n", totalRows, totalCols)

	stackMatrix := stackOfArray{}
	stackMatrix = stackMatrix.Push([]int{0, 0})
	rockStack := stackOfArray{}
	corr := []int{}
	count := 0
	islands := 0
	rockCount := 0
	for !stackMatrix.isEmpty() {
		stackMatrix, corr = stackMatrix.Pop()
		row := corr[0]
		col := corr[1]
		count++
		if !isVisited(grid, row, col) {
			fmt.Printf("numIslands Cell[%d:%d]:%d\n", row, col, grid[row][col])

			if grid[row][col] == rock {
				rockCount++
				rockStack = rockStack.Push([]int{row, col})
			}
			visitCell(grid, row, col)
			//  Push adjacent cells
			if row-1 >= 0 && !isVisited(grid, row-1, col) {
				stackMatrix = stackMatrix.Push([]int{row - 1, col})
			}
			if row+1 < totalRows && !isVisited(grid, row+1, col) {
				stackMatrix = stackMatrix.Push([]int{row + 1, col})
			}
			if col-1 >= 0 && !isVisited(grid, row, col-1) {
				stackMatrix = stackMatrix.Push([]int{row, col - 1})
			}
			if col+1 < totalCols && !isVisited(grid, row, col+1) {
				stackMatrix = stackMatrix.Push([]int{row, col + 1})
			}
		}
	}

	fmt.Printf("\n-------Total items:%d total Rocks:%d\n ", count, rockCount)
	//Scan all rocks
	for !rockStack.isEmpty() {
		rockStack, corr = rockStack.Pop()
		row := corr[0]
		col := corr[1]

		if grid[row][col] == visitedRock {
			islands++
			fmt.Printf("island:%d Rock[%d:%d]:%d\n", islands, row, col, grid[row][col])
			scanIsland2(grid, row, col)
		}
	}

	fmt.Printf("total islands:%d\n", islands)
	return islands
}

//Runtime: 5 ms, faster than 68.90% of Go online submissions for Number of Islands.
//Memory Usage: 6 MB, less than 32.36% of Go online submissions for Number of Islands.

func numIslands(grid [][]byte) int {
	totalRows := len(grid)
	totalCols := len(grid[0])
	islands := 0
	fmt.Printf("numIslands Matrix size:%dx%d\n", totalRows, totalCols)
	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalCols; col++ {
			fmt.Printf("numIslands Cell[%d:%d]:%d\n", row, col, grid[row][col])
			if grid[row][col] == rock {
				islands++
				scanIsland(grid, row, col)
			}
		}
	}
	return islands
}

func mainNumIslands() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("------------------")
	result := numIslands(grid)
	fmt.Printf("numIslands:%d\n", result)

	//grid = [][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}
	//fmt.Println("------------------")
	//result = numIslands(grid)
	//fmt.Printf("numIslands:%d\n", result)
}
