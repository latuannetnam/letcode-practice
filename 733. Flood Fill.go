// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 733. Flood Fill
// Letcode problem: https://leetcode.com/problems/flood-fill/

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
// You are also given three integers sr, sc, and newColor. You should perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
//  plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
//  Replace the color of all of the aforementioned pixels with newColor.

// Return the modified image after performing the flood fill.
// Level: easy

package main

import "fmt"

func fillImage(image [][]int, sr int, sc int, oldColor, newColor int, total_rows, total_cols int) {
	fmt.Printf("Fill image:%d %d\n", sr, sc)
	image[sr][sc] = newColor
	// Check cell on the left of  current cell on the same row
	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		fillImage(image, sr, sc-1, oldColor, newColor, total_rows, total_cols)
	}
	// Check cell on the right of current cell on the same row
	if sc+1 < total_cols && image[sr][sc+1] == oldColor {
		fillImage(image, sr, sc+1, oldColor, newColor, total_rows, total_cols)
	}

	// Check cell on the above of current cell on the same col
	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		fillImage(image, sr-1, sc, oldColor, newColor, total_rows, total_cols)
	}

	// Check cell on the below of current cell on the same col
	if sr+1 < total_rows && image[sr+1][sc] == oldColor {
		fillImage(image, sr+1, sc, oldColor, newColor, total_rows, total_cols)
	}
}

//  Using recursive and Deep first search
// Runtime: 19 ms, faster than 5.06% of Go online submissions for Flood Fill.
// Memory Usage: 4.2 MB, less than 72.59% of Go online submissions for Flood Fill.
func floodFill2(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	total_rows := len(image)
	total_cols := len(image[sr])
	fmt.Printf("Total rows:%d total cols:%d\n", total_rows, total_cols)
	fillImage(image, sr, sc, oldColor, newColor, total_rows, total_cols)
	return image
}

// Use stackOfArray and Deep first search
// Runtime: 10 ms, faster than 42.41% of Go online submissions for Flood Fill.
// Memory Usage: 4.1 MB, less than 72.59% of Go online submissions for Flood Fill.
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	total_rows := len(image)
	total_cols := len(image[sr])
	fmt.Printf("Total rows:%d total cols:%d\n", total_rows, total_cols)
	color_stackOfArray := stackOfArray{}
	color_stackOfArray = color_stackOfArray.Push([]int{sr, sc})
	corr := []int{}
	for !color_stackOfArray.isEmpty() {
		color_stackOfArray, corr = color_stackOfArray.Pop()
		row := corr[0]
		col := corr[1]
		if image[row][col] == oldColor {
			fmt.Printf("Checking cell[%d][%d]: %d -> %d\n", row, col, image[row][col], newColor)
			image[row][col] = newColor
			// Check cell on the left of  current cell on the same row
			if col-1 >= 0 && image[row][col-1] == oldColor {
				color_stackOfArray = color_stackOfArray.Push([]int{row, col - 1})
			}
			// Check cell on the right of current cell on the same row
			if col+1 < total_cols && image[row][col+1] == oldColor {
				color_stackOfArray = color_stackOfArray.Push([]int{row, col + 1})
			}

			// Check cell on the above of current cell on the same col
			if row-1 >= 0 && image[row-1][col] == oldColor {
				color_stackOfArray = color_stackOfArray.Push([]int{row - 1, col})
			}

			// Check cell on the below of current cell on the same col
			if row+1 < total_rows && image[row+1][col] == oldColor {
				color_stackOfArray = color_stackOfArray.Push([]int{row + 1, col})
			}

		}

	}
	return image
}

func mainFloodFill() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	fmt.Printf("Image: %v\n", image)
	newImage := floodFill(image, sr, sc, newColor)
	fmt.Printf("Flood fill:%v\n", newImage)

	image = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	sr = 0
	sc = 0
	newColor = 2
	fmt.Printf("Image: %v\n", image)
	newImage = floodFill(image, sr, sc, newColor)
	fmt.Printf("Flood fill:%v\n", newImage)

	image = [][]int{{0, 0, 0}, {0, 1, 1}}
	sr = 1
	sc = 1
	newColor = 1
	fmt.Printf("Image: %v\n", image)
	newImage = floodFill(image, sr, sc, newColor)
	fmt.Printf("Flood fill:%v\n", newImage)
}
