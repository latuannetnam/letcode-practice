// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Level: medium
// Given the head of a linked list, remove the nth node from the end of the list and return its head.

package main

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	node_slice := []*ListNode{}
	var node *ListNode = head
	for node != nil {
		node_slice = append(node_slice, node)
		node = node.Next
	}

	list_len := len(node_slice)
	index := list_len - n
	if index <= 0 {
		return head.Next
	} else if index == list_len-1 {
		node_slice[index-1].Next = nil
	} else {

		node_slice[index-1].Next = node_slice[index+1]
	}

	return head
}

func mainRemoveNthFromEnd() {
	// lists := []int{1, 2, 3, 4, 5}
	lists := []int{1, 2}
	n := 2
	head := slice2List(lists)
	// lists_2 := list2Slice(head)
	// fmt.Printf("lists: %v\n", lists_2)
	head_after := removeNthFromEnd(head, n)
	lists_after := list2Slice(head_after)
	fmt.Printf("nth: %d head: %v\n", n, lists_after)
}
