// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/21/2022
// Letcode problem: 986. Interval List Intersections
// Letcode link: https://leetcode.com/problems/interval-list-intersections/
// Level: medium
//You are given two lists of closed intervals,
//firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj].
//Each list of intervals is pairwise disjoint and in sorted order.
//
//Return the intersection of these two interval lists.
//
//A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.
//
//The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval.
//For example, the intersection of [1, 3] and [2, 4] is [2, 3].
package main

import "fmt"

//Runtime: 37 ms, faster than 21.84% of Go online submissions for Interval List Intersections.
//Memory Usage: 7 MB, less than 64.37% of Go online submissions for Interval List Intersections.
func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	firstRows := len(firstList)
	firstPointer := 0
	secondRows := len(secondList)
	secondPointer := 0
	interSectionTypeMap := map[int]string{
		-1: "No intersect",
		0:  "secondList inside firstList",
		1:  "firstList inside secondList",
		2:  "firstList and secondList intersect type 1",
		3:  "firstList and secondList intersect type 2",
		4:  "Firstlist < second list",
		5:  "Secondlist < Firstlist",
	}
	for firstPointer < firstRows && secondPointer < secondRows {
		fmt.Printf("Assert firstPointer:%d firstList[%d]:%v\n", firstPointer, firstPointer, firstList[firstPointer])
		fmt.Printf("Assert secondPointer:%d secondList[%d]:%v\n", secondPointer, secondPointer, secondList[secondPointer])
		firstCols := len(firstList[firstPointer])
		secondCols := len(secondList[secondPointer])
		var interSection []int
		interSectionType := -1
		if firstList[firstPointer][0] <= secondList[secondPointer][0] && firstList[firstPointer][firstCols-1] >= secondList[secondPointer][secondCols-1] {
			//secondList inside firstList
			interSection = append(interSection, secondList[secondPointer][0], secondList[secondPointer][secondCols-1])
			secondPointer++
			interSectionType = 0
		} else if firstList[firstPointer][0] >= secondList[secondPointer][0] && firstList[firstPointer][firstCols-1] <= secondList[secondPointer][secondCols-1] {
			//firstList inside secondList
			interSection = append(interSection, firstList[firstPointer][0], firstList[firstPointer][firstCols-1])
			firstPointer++
			interSectionType = 1
		} else if firstList[firstPointer][0] <= secondList[secondPointer][0] && secondList[secondPointer][0] <= firstList[firstPointer][firstCols-1] && firstList[firstPointer][firstCols-1] <= secondList[secondPointer][secondCols-1] {
			//firstList and secondList intersect type 1
			interSection = append(interSection, secondList[secondPointer][0], firstList[firstPointer][firstCols-1])
			firstPointer++
			interSectionType = 2
		} else if firstList[firstPointer][0] >= secondList[secondPointer][0] && firstList[firstPointer][0] <= secondList[secondPointer][secondCols-1] && firstList[firstPointer][firstCols-1] >= secondList[secondPointer][secondCols-1] {
			//firstList and secondList intersect type 2
			interSection = append(interSection, firstList[firstPointer][0], secondList[secondPointer][secondCols-1])
			secondPointer++
			interSectionType = 3
		} else if firstList[firstPointer][0] <= secondList[secondPointer][0] && firstList[firstPointer][firstCols-1] <= secondList[secondPointer][secondCols-1] {
			//	Firstlist < second list
			firstPointer++
			interSectionType = 4
		} else {
			//	secondlist < firstlist
			secondPointer++
			interSectionType = 5
		}
		fmt.Printf("Intersection type:%s:%v\n", interSectionTypeMap[interSectionType], interSection)
		if len(interSection) > 0 {
			result = append(result, interSection)
		}

	}
	return result
}

//https://leetcode.com/problems/interval-list-intersections/solution/
//Runtime: 29 ms, faster than 40.23% of Go online submissions for Interval List Intersections.
//Memory Usage: 6.9 MB, less than 74.14% of Go online submissions for Interval List Intersections.
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	firstRows := len(firstList)
	firstPointer := 0
	secondRows := len(secondList)
	secondPointer := 0
	for firstPointer < firstRows && secondPointer < secondRows {
		fmt.Printf("Assert firstPointer:%d firstList[%d]:%v\n", firstPointer, firstPointer, firstList[firstPointer])
		fmt.Printf("Assert secondPointer:%d secondList[%d]:%v\n", secondPointer, secondPointer, secondList[secondPointer])
		lo := max(firstList[firstPointer][0], secondList[secondPointer][0])
		hi := min(firstList[firstPointer][1], secondList[secondPointer][1])
		fmt.Printf("lo:%d hi:%d\n", lo, hi)
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}
		if firstList[firstPointer][1] < secondList[secondPointer][1] {
			firstPointer++
		} else {
			secondPointer++
		}
	}
	return result
}

func mainIntervalIntersection() {
	firstList := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	secondList := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	//secondList := [][]int{}
	fmt.Printf("First list:%v\n", firstList)
	fmt.Printf("Second list:%v\n", secondList)
	fmt.Printf("Interval section:%v\n", intervalIntersection(firstList, secondList))

}
