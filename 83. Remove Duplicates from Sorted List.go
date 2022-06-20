// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/20/2022
// Letcode problem: 83. Remove Duplicates from Sorted List
// Letcode link: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// Level: easy
//Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
//Return the linked list sorted as well.
package main

import "fmt"

//Runtime: 10 ms, faster than 10.33% of Go online submissions for Remove Duplicates from Sorted List.
//Memory Usage: 3.2 MB, less than 11.21% of Go online submissions for Remove Duplicates from Sorted List.
func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		fmt.Printf("node:%v node Next:%v\n", node, node.Next)
		nextNode := node.Next
		for nextNode != nil && node.Val == nextNode.Val {
			node.Next = nextNode.Next
			nextNode = node.Next
		}
		node = node.Next

	}
	return head
}

func mainDeleteDuplicates() {
	list1 := []int{1, 1, 2, 3, 3}
	//list1 := []int{1, 1, 1}
	head := slice2List(list1)
	fmt.Printf("list:%v\n", list1)
	head = deleteDuplicates(head)
	list2 := list2Slice(head)
	fmt.Printf("List after remove:%v\n", list2)
}
