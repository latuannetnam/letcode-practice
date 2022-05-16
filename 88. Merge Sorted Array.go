// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 88. Merge Sorted Array
// Letcode problem: https://leetcode.com/problems/merge-sorted-array/
// Level: medium

package main

import "fmt"

func mergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m)
	copy(nums3, nums1[:m])
	m_pos := 0
	n_pos := 0
	pos := 0
	for pos < (m + n) {
		if m_pos < m && n_pos < n {
			if nums2[n_pos] < nums3[m_pos] {
				nums1[pos] = nums2[n_pos]
				n_pos++
			} else {
				nums1[pos] = nums3[m_pos]
				m_pos++
			}
		} else if m_pos == m {
			nums1[pos] = nums2[n_pos]
			n_pos++
		} else {
			nums1[pos] = nums3[m_pos]
			m_pos++
		}
		pos++
	}
}

func mainMergeSortedArrays() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Printf("nums1:%v nums2:%v\n", nums1[:m], nums2[:n])
	mergeSortedArrays(nums1, m, nums2, n)
	fmt.Printf("merge sort array:%v\n", nums1)
}
