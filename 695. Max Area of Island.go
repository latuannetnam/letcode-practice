// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 5/30/2022
// Letcode problem: 695. Max Area of Island
// Letcode link: https://leetcode.com/problems/max-area-of-island/
// Level: medium
//You are given an m x n binary matrix grid.
//An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)
//You may assume all four edges of the grid are surrounded by water.
//The area of an island is the number of cells with a value 1 in the island.
////Return the maximum area of an island in grid. If there is no island, return 0

package main

import "fmt"

//Use DFS to scan all related rocks of island (related cells = 1)
func scanIslandArea(grid [][]int, row int, col int) int {
	area := 0
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
		if grid[row][col] == 1 {
			area += grid[row][col]
			fmt.Printf("scanIsland Cell[%d:%d]:%d area:%d\n", row, col, grid[row][col], area)
			//  Push adjacent cells
			if row-1 >= 0 && grid[row-1][col] == 1 {
				stackMatrix = stackMatrix.Push([]int{row - 1, col})
			}
			if row+1 < totalRows && grid[row+1][col] == 1 {
				stackMatrix = stackMatrix.Push([]int{row + 1, col})
			}

			if col-1 >= 0 && grid[row][col-1] == 1 {
				stackMatrix = stackMatrix.Push([]int{row, col - 1})
			}

			if col+1 < totalCols && grid[row][col+1] == 1 {
				stackMatrix = stackMatrix.Push([]int{row, col + 1})
			}
			grid[row][col] = 0
		}

	}

	fmt.Printf("scanIsland: total items:%d area:%d\n", count, area)
	return area

}

//Runtime: 14 ms, faster than 58.29% of Go online submissions for Max Area of Island.
//Memory Usage: 6.1 MB, less than 30.06% of Go online submissions for Max Area of Island.
func maxAreaOfIsland(grid [][]int) int {
	totalRows := len(grid)
	totalCols := len(grid[0])
	maxArea := 0
	islands := 0
	fmt.Printf("Size of matrix:%dx%d\n", totalRows, totalCols)
	for row := 0; row < totalRows; row++ {
		for col := 0; col < totalCols; col++ {
			fmt.Printf("maxAreaOfIsland Cell[%d:%d]:%d\n", row, col, grid[row][col])
			if grid[row][col] == 1 {
				islands++
				area := scanIslandArea(grid, row, col)
				if maxArea < area {
					maxArea = area
				}
			}
		}
	}

	fmt.Printf("Total island:%d max area:%d\n", islands, maxArea)
	return maxArea
}

func mainMaxAreaOfIsland() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	result := maxAreaOfIsland(grid)
	fmt.Printf("maxAreaOfIsland:%d\n", result)
}
