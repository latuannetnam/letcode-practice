// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/20/2022
// Letcode problem: 82. Remove Duplicates from Sorted List II
// Letcode link: https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
// Level: medium
//Given the head of a sorted linked list,
//delete all nodes that have duplicate numbers,
//leaving only distinct numbers from the original list.
//Return the linked list sorted as well.
package main

import "fmt"

//Runtime: 2 ms, faster than 86.28% of Go online submissions for Remove Duplicates from Sorted List II.
//Memory Usage: 3.1 MB, less than 15.83% of Go online submissions for Remove Duplicates from Sorted List II.
func deleteDuplicatesII(head *ListNode) *ListNode {
	node := head
	var prevNode *ListNode = nil
	for node != nil && node.Next != nil {
		isDuplicated := false
		fmt.Printf("prev:%v node:%v next:%d\n", prevNode, node, node.Next)
		for node.Next != nil && node.Val == node.Next.Val {
			node = node.Next
			isDuplicated = true
		}
		if isDuplicated {
			if prevNode == nil {
				head = node.Next
				node = head
				fmt.Printf("Move head to:%v\n", head)
			} else {
				prevNode.Next = node.Next
				node = prevNode.Next
			}
		} else {
			prevNode = node
			node = prevNode.Next
		}

	}
	return head
}

func mainDeleteDuplicatesII() {
	//list1 := []int{1, 2, 3, 3, 4, 4, 5}
	//list1 := []int{1, 1, 1, 2, 2, 3}
	list1 := []int{1, 2, 2, 2}
	head := slice2List(list1)
	fmt.Printf("list:%v\n", list1)
	head = deleteDuplicatesII(head)
	list2 := list2Slice(head)
	fmt.Printf("After delete:%v\n", list2)
}
