// Utilities for letcode
package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// ---------------Other utils
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

// ---------- Array utils
// Compared 2 sorted array
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

// -------- String utils -------
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func arrayToString(arr []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprintf("%03v", arr), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// Runtime: 47 ms, faster than 26.09% of Go online submissions for Reverse String.
// Memory Usage: 6.7 MB, less than 57.52% of Go online submissions for Reverse String.
func reverseStringByte(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	n := len(s)
	for i := 0; i < n/2; i++ {
		rns[i], rns[n-i-1] = rns[n-i-1], rns[i]
	}
	return string(rns)
}

func isPalindromeString(s1 string) bool {
	if len(s1) == 1 {
		return true
	}
	n := len(s1)
	for i := 0; i < n/2; i++ {
		if s1[i] != s1[n-i-1] {
			return false
		}
	}
	return true
}

func isAlphanumeric(s byte) bool {
	return (s >= '0' && s <= '9') || (s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
}

func charLowerCase(s byte) byte {
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return s
}

// ---------- Binary search ------
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

// --------------- Stack of Array------------
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

// --------------- Queue of Array------------
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

// --------------- Stack of String Array------------
type stackOfString []string

func (s stackOfString) Push(v string) stackOfString {
	return append(s, v)
}

func (s stackOfString) Pop() (stackOfString, string) {
	// FIXME: What do we do if the stackOfArray is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stackOfString) isEmpty() bool {
	return len(s) == 0
}

// --------------- Stack of Int Array------------
type stackOfInt []int

func (s stackOfInt) Push(v int) stackOfInt {
	return append(s, v)
}

func (s stackOfInt) Pop() (stackOfInt, int) {
	// FIXME: What do we do if the stackOfArray is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stackOfInt) isEmpty() bool {
	return len(s) == 0
}

func (s stackOfInt) Top() int {
	if len(s) == 0 {
		return int(math.Inf(0))
	} else {
		return s[0]
	}
}

func (s stackOfInt) Bottom() int {
	if len(s) == 0 {
		return int(math.Inf(0))
	} else {
		return s[len(s)-1]
	}
}

// --------------- Stack of ListNode------------
type stackOfListNode []*ListNode

func (s stackOfListNode) Push(v *ListNode) stackOfListNode {
	return append(s, v)
}

func (s stackOfListNode) Pop() (stackOfListNode, *ListNode) {
	// FIXME: What do we do if the stackOfArray is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stackOfListNode) isEmpty() bool {
	return len(s) == 0
}

func (s stackOfListNode) Bottom() *ListNode {
	if len(s) == 0 {
		return nil
	} else {
		return s[len(s)-1]
	}
}

// ------------- Binary Tree
// https://www.golangprograms.com/golang-program-to-implement-binary-tree.html
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
