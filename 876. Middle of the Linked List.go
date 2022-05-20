// Author: Le Anh Tuan (latuannetnam@gmail.com)
// 876. Middle of the Linked List
// Letcode problem: Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
// Level: Easy
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.

package main

import "fmt"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
// Memory Usage: 1.9 MB, less than 73.42% of Go online submissions for Middle of the Linked List
func middleNode(head *ListNode) *ListNode {
	node := head
	end := head
	for node.Next != nil && end != nil && end.Next != nil {
		fmt.Printf("head:%d node:%d end:%d\n", head.Val, node.Val, end.Val)
		node = node.Next
		end = end.Next.Next
	}
	return node
}

func mainMiddleNode() {
	head_arr := []int{1, 2, 3, 4, 5}
	head := slice2List(head_arr)
	fmt.Printf("head:%v\n", head_arr)
	node := middleNode(head)
	fmt.Printf("middle node:%d\n", node.Val)

	head_arr = []int{1, 2, 3, 4, 5, 6}
	head = slice2List(head_arr)
	fmt.Printf("head:%v\n", head_arr)
	node = middleNode(head)
	fmt.Printf("middle node:%d\n", node.Val)
}
