// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Letcode problem: 2. Add Two Numbers
// Letcode link: https://leetcode.com/problems/add-two-numbers/
// Level: medium
// Topics: Link list, Math, Recursion
// Problem detail:
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

package main

import "fmt"

func addTwoNumbersWrong(l1 *ListNode, l2 *ListNode) *ListNode {
	// Reverse 2 list
	rl1 := reverseList(l1)
	rl2 := reverseList(l2)

	// Do rl3 = rl1 + rl2
	prevNode1 := rl1
	node1 := rl1
	node2 := rl2
	remain := 0
	for node1 != nil && node2 != nil {
		sum := node1.Val + node2.Val + remain
		newVal := sum % 10
		fmt.Printf("node1:%d + node2:%d = %d ( remain:%d) => newVal:%d, newRemain:%d\n", node1.Val, node2.Val, sum, remain, newVal, sum/10)
		remain = sum / 10
		node1.Val = newVal
		prevNode1 = node1
		node1 = node1.Next
		node2 = node2.Next
	}

	if node1 == nil && node2 != nil {
		// list1 size < list 2 size
		prevNode1.Next = node2
	}

	fmt.Printf("rl3:%v\n", list2Slice(rl1))

	l3 := reverseList(rl1)

	return l3
}

// Runtime: 20 ms, faster than 32.56% of Go online submissions for Add Two Numbers.
// Memory Usage: 4.4 MB, less than 98.68% of Go online submissions for Add Two Numbers.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prevNode1 := l1
	node1 := l1
	node2 := l2
	remain := 0
	for node1 != nil && node2 != nil {
		sum := node1.Val + node2.Val + remain
		newVal := sum % 10
		fmt.Printf("node1:%d + node2:%d = %d ( remain:%d) => newVal:%d, newRemain:%d\n", node1.Val, node2.Val, sum, remain, newVal, sum/10)
		remain = sum / 10
		node1.Val = newVal
		prevNode1 = node1
		node1 = node1.Next
		node2 = node2.Next
	}

	fmt.Printf("Last remain1:%d\n", remain)

	if node1 == nil && node2 != nil {
		// list1 size < list 2 size
		fmt.Println("Extend list1 for list2")
		prevNode1.Next = node2
		for node2 != nil {
			sum := node2.Val + remain
			newVal := sum % 10
			fmt.Printf("node2:%d + remain:%d = %d => newVal:%d, newRemain:%d\n", node2.Val, remain, sum, newVal, sum/10)
			remain = sum / 10
			node2.Val = newVal
			prevNode1 = node2
			node2 = node2.Next
		}
	} else if node1 != nil {
		// list1 size > list 2 size
		fmt.Println("Continue list1")
		for node1 != nil {
			sum := node1.Val + remain
			newVal := sum % 10
			fmt.Printf("node1:%d + remain:%d = %d => newVal:%d, newRemain:%d\n", node1.Val, remain, sum, newVal, sum/10)
			remain = sum / 10
			node1.Val = newVal
			prevNode1 = node1
			node1 = node1.Next
		}
	}

	fmt.Printf("Last remain2:%d\n", remain)

	if remain > 0 {
		fmt.Println("Extend list1 for remain")
		node2 = &ListNode{
			Val: remain,
		}
		prevNode1.Next = node2
	}

	return l1
}

func mainAddTwoNumbers() {
	arr1 := []int{2, 4, 3}
	arr2 := []int{5, 6, 4}

	arr2 = []int{9, 9, 9, 9, 9, 9, 9}
	arr1 = []int{9, 9, 9, 9}

	arr2 = []int{0}
	arr1 = []int{0}

	l1 := slice2List(arr1)
	l2 := slice2List(arr2)
	fmt.Printf("l1:%v l2:%v\n", arr1, arr2)
	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("l1+l2=%v\n", list2Slice(l3))

}
