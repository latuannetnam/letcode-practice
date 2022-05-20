// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Level: medium
// Given the head of a linked list, remove the nth node from the end of the list and return its head.
// Remove in one pass
package main

import "fmt"

// Convert Linklist -> slice and remove Nth node from end of list
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
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

// Runtime: 3 ms, faster than 43.38% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.1 MB, less than 73.05% of Go online submissions for Remove Nth Node From End of List.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	end := node
	if node.Next == nil {
		return nil
	}
	// Move end forward nth node
	i := 1
	for end.Next != nil && i <= n {
		end = end.Next
		i++
	}
	fmt.Printf("init end node:%d i:%d n:%d\n", end.Val, i, n)
	// If delete head
	if end.Next == nil {
		return head.Next
	}

	// Move end node to end of list
	// node = end - n + 1
	for end.Next != nil {
		node = node.Next
		end = end.Next
	}
	fmt.Printf("node:%d end node:%d\n", node.Val, end.Val)
	// Remove node.Next from list
	node.Next = node.Next.Next

	return head
}

func mainRemoveNthFromEnd() {
	lists := []int{1, 2, 3, 4, 5}
	// lists := []int{1, 2}
	n := 5
	head := slice2List(lists)
	// lists_2 := list2Slice(head)
	fmt.Printf("lists: %v n:%d\n", lists, n)
	head_after := removeNthFromEnd(head, n)
	lists_after := list2Slice(head_after)
	fmt.Printf("nth: %d head: %v\n", n, lists_after)
}
