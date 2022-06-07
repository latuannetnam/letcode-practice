// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/6/2022
// Letcode problem: 77. Combinations
// Letcode link: https://leetcode.com/problems/combinations/
// Level: medium
//Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
//You may return the answer in any order.
package main

import "fmt"

func makeCombineKofN2(arr []int, start int, end int, k int, combineResult [][]int) [][]int {
	var newResult [][]int
	//fmt.Printf("makeCombineKofN  start:%d end:%d k:%d combineResult:%v\n", start, end, k, combineResult)
	if k == 1 {
		fmt.Printf("k:%d final combine\n", k)
		for _, subArr := range combineResult {
			subArr = append(subArr, arr[start])
			fmt.Printf("subArr:%v\n", subArr)
			newResult = append(newResult, subArr)
		}
	} else {

		tempResult := make([][]int, 0)
		for _, subArr := range combineResult {
			subArr = append(subArr, arr[start])
			subArr1 := make([]int, len(subArr))
			copy(subArr1, subArr)
			tempResult = append(tempResult, subArr1)
		}
		for i := start + 1; i <= end-k+1; i++ {
			fmt.Printf("Make next combine: from %d-%d of %d tempResult:%v\n", i, end, k-1, tempResult)
			newResult = append(newResult, makeCombineKofN2(arr, i, end, k-1, tempResult)...)
			fmt.Printf("After Make next combine: from %d-%d of %d new-result:%v\n", i, end, k-1, newResult)

		}
	}
	//fmt.Printf("After makeCombineKofN  start:%d end:%d k:%d combineResult:%v\n", start, end, k, newResult)
	return newResult
}

func makeCombineKofN(arr []int, start int, end int, k int, combineResult []int) [][]int {
	var newResult [][]int
	//fmt.Printf("makeCombineKofN  start:%d end:%d k:%d combineResult:%v\n", start, end, k, combineResult)
	combineResult = append(combineResult, arr[start])
	if k == 1 {
		newResult = append(newResult, combineResult)
	} else {
		for i := start + 1; i <= end-k+1; i++ {
			tempResult := make([]int, len(combineResult))
			copy(tempResult, combineResult)
			fmt.Printf("Make next combine: from %d-%d of %d tempResult:%v\n", i, end, k-1, tempResult)
			newResult = append(newResult, makeCombineKofN(arr, i, end, k-1, tempResult)...)
			fmt.Printf("After Make next combine: from %d-%d of %d new-result:%v\n", i, end, k-1, newResult)
		}
	}
	return newResult
}

//Runtime: 10 ms, faster than 86.00% of Go online submissions for Combinations.
//Memory Usage: 8.8 MB, less than 34.89% of Go online submissions for Combinations.

//Runtime: 13 ms, faster than 73.43% of Go online submissions for Combinations.
//Memory Usage: 9.1 MB, less than 29.57% of Go online submissions for Combinations.
func combineKofNInteger(n int, k int) [][]int {
	var combineResult [][]int
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i+1)

	}

	for i := 0; i <= n-k; i++ {
		fmt.Printf("\n-------\n")
		fmt.Printf("makeCombineKofN for i:%d k:%d\n", i, k)
		combineResult = append(combineResult, makeCombineKofN(arr, i, len(arr), k, []int{})...)
	}

	return combineResult
}

func mainCombineKofNInteger() {
	n := 10
	k := 7
	fmt.Printf("N:%d K:%d\n", n, k)
	combineResult := combineKofNInteger(n, k)
	fmt.Printf("Combine result:%v\n", combineResult)

}
