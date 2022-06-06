// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/6/2022
// Letcode problem: 206. Reverse Linked List
// Letcode link: https://leetcode.com/problems/reverse-linked-list/
// Level: easy
//Given the head of a singly linked list, reverse the list, and return the reversed list.

package main

import "fmt"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
//Memory Usage: 2.6 MB, less than 73.56% of Go online submissions for Reverse Linked List.
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head
	var nextNode *ListNode = nil
	if node != nil {
		nextNode = node.Next
	}
	for node != nil && nextNode != nil {
		fmt.Printf("node:%v nextNode:%v\n", node, nextNode)
		tempNode := nextNode.Next
		nextNode.Next = node
		node = nextNode
		nextNode = tempNode
	}
	head.Next = nil
	return node
}
func mainReverseList() {
	list := []int{1, 2, 3, 4, 5}
	head := slice2List(list)
	fmt.Printf("List:%v\n", list)
	head2 := reverseList(head)
	list2 := list2Slice(head2)
	fmt.Printf("Reverse list:%v\n", list2)

	fmt.Printf("\n-----------\n")
	list = []int{1, 2}
	head = slice2List(list)
	fmt.Printf("List:%v\n", list)
	head2 = reverseList(head)
	list2 = list2Slice(head2)
	fmt.Printf("Reverse list:%v\n", list2)

}
