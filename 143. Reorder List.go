// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 143. Reorder List
// Letcode link: https://leetcode.com/problems/reorder-list/
// Level: medium
// Topics: Linked List, Two pointers, Stack, Recursion
// Problem detail:
// You are given the head of a singly linked-list. The list can be represented as:
// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:
// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.

package main

import "fmt"

// Stack
// Runtime: 18 ms, faster than 34.03% of Go online submissions for Reorder List.
// Memory Usage: 6.3 MB, less than 22.99% of Go online submissions for Reorder List.
func reorderListStack(head *ListNode) {
	var tempStack stackOfListNode
	node := head

	// Push all node to stack to make reverse node list
	for node != nil {
		tempStack = tempStack.Push(node)
		node = node.Next
	}

	node = head
	var nodeOfStack *ListNode = nil

	for node != tempStack.Bottom() && node.Next != tempStack.Bottom() {
		tempStack, nodeOfStack = tempStack.Pop()
		tempNode := node.Next
		fmt.Printf("node:%d -> nodeOfStack:%d -> node.Next:%d\n", node.Val, nodeOfStack.Val, tempNode.Val)
		node.Next = nodeOfStack
		nodeOfStack.Next = tempNode
		node = tempNode
	}

	fmt.Printf("Last node:%v\n", node)
	if node.Next == tempStack.Bottom() {
		tempStack.Bottom().Next = nil
	} else {
		node.Next = nil
	}

	fmt.Printf("Done reorder list\n")
}

// Neetcode
// Two pointers
// Runtime: 13 ms, faster than 59.70% of Go online submissions for Reorder List.
// Memory Usage: 5.4 MB, less than 46.57% of Go online submissions for Reorder List.
func reorderListTwoPointers(head *ListNode) {
	// Find middle of list
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fmt.Printf("Midle node:%v\n", slow)

	// Reverse second haft of list
	second := slow.Next
	slow.Next = nil
	prev := slow.Next

	for second != nil {
		fmt.Printf("Second:%v\n", second)
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	// Merge 2 haft of list
	first := head
	second = prev
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next
		first.Next = second
		second.Next = tmp1
		first = tmp1
		second = tmp2
	}

}

func mainReorderList() {
	headArr := []int{1, 2, 3, 4}
	head := slice2List(headArr)
	fmt.Printf("head:%v\n", headArr)
	// reorderListStack(head)
	reorderListTwoPointers(head)
	headArr = list2Slice(head)
	fmt.Printf("Reverse list:%v\n", headArr)
}
