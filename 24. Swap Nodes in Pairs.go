// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/swap-nodes-in-pairs/
// Level: medium

// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)

// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

package main

import "fmt"

//  Definition for singly-linked list.
// Defined in 19. Remove Nth Node From End of List
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	var first_node, second_node, new_head, last_node *ListNode

	if head == nil || head.Next == nil {
		return head
	}

	first_node = head
	new_head = head.Next
	last_node = nil
	for first_node != nil && first_node.Next != nil {
		second_node = first_node.Next
		temp_node := second_node.Next
		second_node.Next = first_node
		first_node.Next = temp_node
		if last_node != nil {
			last_node.Next = second_node
		}
		last_node = first_node
		// lists_after := list2Slice(new_head)
		// fmt.Printf("first node: %d second node:%d temp node:%d \n", first_node.Val, second_node.Val, temp_node)
		// fmt.Printf("head after:%v\n", lists_after)

		first_node = first_node.Next
	}

	return new_head
}

func mainSwapPairs() {
	lists := []int{1, 2, 3, 4, 5}
	// lists := []int{}
	head := slice2List(lists)
	head_after := swapPairs(head)
	lists_after := list2Slice(head_after)
	fmt.Printf("head before: %v head after:%v\n", lists, lists_after)
}
