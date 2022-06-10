// Utilities for letcode
package main

import "math"

//---------------Other utils
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(num1 int, num2 int) int {
	return int(math.Min(float64(num1), float64(num2)))
}

func max(num1 int, num2 int) int {
	return int(math.Max(float64(num1), float64(num2)))
}

func reverseArray(nums []int, start, end int) {
	left := start
	right := end - 1
	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// ----------------  Definition for singly-linked list.
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

//--------------- Stack of Array------------
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

// ------------- Binary Tree
//https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) insert(item int) *TreeNode {
	if t == nil {
		t = &TreeNode{
			Val:   item,
			Left:  nil,
			Right: nil,
		}
	} else if t.Val <= item {
		if t.Left == nil {
			t.Left = &TreeNode{
				Val:   item,
				Left:  nil,
				Right: nil,
			}
		} else {
			t.Left = t.Left.insert(item)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{
				Val:   item,
				Left:  nil,
				Right: nil,
			}
		} else {
			t.Right = t.Right.insert(item)
		}
	}
	return t
}

func (t *TreeNode) leftTraverse(lists []int) []int {
	if t != nil {
		item := t.Val
		lists = append(lists, item)
		if t.Left != nil {
			lists = t.Left.leftTraverse(lists)
		}
		if t.Right != nil {
			lists = t.Right.leftTraverse(lists)
		}
	}
	return lists
}
func slice2BTree(lists []int) *TreeNode {
	var root *TreeNode = nil
	for _, item := range lists {
		root = root.insert(item)
	}
	return root
}

func bTree2slice(node *TreeNode) []int {
	lists := []int{}
	lists = node.leftTraverse(lists)
	return lists
}
