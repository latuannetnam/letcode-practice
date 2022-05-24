// Utilities for letcode
package main

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func slice2List(lists []int) *ListNode {
	var head *ListNode = nil
	var node *ListNode = nil
	for index := range lists {
		if head == nil {
			head = &ListNode{
				Val:  lists[index],
				Next: nil,
			}
			node = head
		} else {
			new_node := &ListNode{
				Val:  lists[index],
				Next: nil,
			}
			node.Next = new_node
			node = new_node
		}
	}
	return head
}

func list2Slice(head *ListNode) []int {
	lists := []int{}
	var node *ListNode = nil
	node = head
	for node != nil {
		lists = append(lists, node.Val)
		node = node.Next
	}
	return lists
}

type stackOfArray [][]int

func (s stackOfArray) Push(v []int) stackOfArray {
	return append(s, v)
}

func (s stackOfArray) Pop() (stackOfArray, []int) {
	// FIXME: What do we do if the stackOfArray is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stackOfArray) isEmpty() bool {
	return len(s) == 0
}
