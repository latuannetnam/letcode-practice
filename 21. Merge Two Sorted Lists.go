// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/6/2022
// Letcode problem: 21. Merge Two Sorted Lists
// Letcode link: https://leetcode.com/problems/merge-two-sorted-lists/
// Level: easy
//You are given the heads of two sorted linked lists list1 and list2.
//Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
//Return the head of the merged linked list.
package main

import "fmt"

//Runtime: 5 ms, faster than 27.15% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.6 MB, less than 10.42% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	var root *ListNode = nil
	var node3 *ListNode = nil
	var node *ListNode = nil
	node1 := list1
	node2 := list2
	for node1 != nil || node2 != nil {
		if node1 != nil && node2 != nil {
			if node1.Val <= node2.Val {
				node = &ListNode{node1.Val, nil}
				node1 = node1.Next
			} else {
				node = &ListNode{node2.Val, nil}
				node2 = node2.Next
			}
		} else if node1 != nil {
			node = &ListNode{node1.Val, nil}
			node1 = node1.Next
		} else {
			node = &ListNode{node2.Val, nil}
			node2 = node2.Next
		}
		//Assign new node to list3
		if root == nil {
			root = node
			node3 = root
		} else {
			node3.Next = node
			node3 = node
		}
	}
	return root
}

//Runtime: 7 ms, faster than 15.43% of Go online submissions for Merge Two Sorted Lists.
//Memory Usage: 2.6 MB, less than 43.92% of Go online submissions for Merge Two Sorted Lists.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root *ListNode = nil
	var node3 *ListNode = nil
	var node *ListNode = nil
	node1 := list1
	node2 := list2
	for node1 != nil || node2 != nil {
		if node1 != nil && node2 != nil {
			if node1.Val <= node2.Val {
				node = node1
				node1 = node1.Next
			} else {
				node = node2
				node2 = node2.Next
			}
		} else if node1 != nil {
			node = node1
			node1 = node1.Next
		} else {
			node = node2
			node2 = node2.Next
		}
		//Assign new node to list3
		if root == nil {
			root = node
			node3 = root
		} else {
			node3.Next = node
			node3 = node
		}
	}
	return root
}
func mainMergeTwoLists() {
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}
	root1 := slice2List(list1)
	root2 := slice2List(list2)
	fmt.Printf("List1:%v root1:%v\n", list1, root1)
	fmt.Printf("List2:%v root2:%v\n", list2, root2)
	root3 := mergeTwoLists(root1, root2)
	list3 := list2Slice(root3)
	fmt.Printf("Merger 2 list:%v\n", list3)
}
