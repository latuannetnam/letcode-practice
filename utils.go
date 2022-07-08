// Utilities for letcode
package main

import (
	"fmt"
	"math"
	"strings"
)

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

//---------- Array utils
//Compared 2 sorted array
func arrayEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
func checkSubArrayExits(subArr []int, arr [][]int) bool {
	for i := range arr {
		if arrayEqual(arr[i], subArr) {
			return true
		}
	}
	return false
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

//-------- String utils -------
func arrayToString(arr []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprintf("%03v", arr), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

//---------- Binary search ------
func binarySearch(nums []int, target int, start int, end int) int {
	result := -1
	left := start
	right := end
	fmt.Printf("binarySearch start:%d end:%d %v:%d\n", start, end, nums[start:end+1], target)
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left:%d mid:%d right:%d:%d\n", left, mid, right, nums[mid])
		if nums[mid] == target {
			fmt.Printf("Found target %d at pos:%d\n", target, mid)
			result = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Printf("binarySearch result:%v\n", result)
	return result
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

//--------------- Queue of Array------------
type queueOfArray [][]int

func (s queueOfArray) Push(v []int) queueOfArray {
	return append(s, v)
}

func (s queueOfArray) Pop() (queueOfArray, []int) {
	// FIXME: What do we do if the queue is empty, though?
	return s[1:], s[0]
}

func (s queueOfArray) isEmpty() bool {
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
