// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 6/3/2022
// Letcode problem: 116. Populating Next Right Pointers in Each Node
// Letcode link: https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// Level: medium
//You are given a perfect binary tree where all leaves are on the same level, and every parent has two children.
//Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
//Initially, all next pointers are set to NULL.
package main

import (
	"fmt"
	"math"
	"strconv"
)

//Definition for a Node.

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (t *Node) midTraverse() {
	if t != nil {
		fmt.Printf("Node:%d Left:%v Right:%v\n", t.Val, t.Left, t.Right)
		if t.Left != nil {
			t.Left.midTraverse()
		}
		if t.Right != nil {
			t.Right.midTraverse()
		}
	}
}

func (t *Node) addNode(level int, offset int, item int) *Node {
	levelLimit := int(math.Pow(2, float64(level)))
	fmt.Printf("Addnode: parent: %v level:%d offset:%d/%d item:%d\n", t, level, offset, levelLimit, item)
	if t == nil {
		t = &Node{Val: item}
	} else {
		if level == 1 {
			if offset%2 == 0 {
				t.Left = &Node{Val: item}
			} else {
				t.Right = &Node{Val: item}
			}
		} else {
			if offset < levelLimit/2 {
				t.Left = t.Left.addNode(level-1, offset, item)
			} else {
				t.Right = t.Right.addNode(level-1, offset-levelLimit/2, item)
			}
		}

	}
	fmt.Printf("After Addnode: parent: %v level:%d offset:%d item:%d\n", t, level, offset, item)
	return t
}

func _116Slice2BTreeNode(lists []int) *Node {
	var root *Node = nil
	level := 0
	totalItemPerLevel := 0

	for index := range lists {
		levelLimit := math.Pow(2, float64(level))
		if index-totalItemPerLevel >= int(levelLimit) {
			totalItemPerLevel += int(levelLimit)
			level++
		}
		offset := index - totalItemPerLevel
		//fmt.Printf("index:%d level:%d levelLimit:%d totalItemPerLevel:%d offset:%d/%d item:%d\n", index, level, int(levelLimit), totalItemPerLevel, offset, int(math.Pow(2, float64(level))), lists[index])
		root = root.addNode(level, offset, lists[index])
	}
	return root
}

func traverseNextPointer(root *Node, lists []string) []string {
	if root == nil {
		return lists
	}
	lists = append(lists, strconv.Itoa(root.Val))
	node := root.Next
	for node != nil {
		lists = append(lists, "->"+strconv.Itoa(node.Val))
		node = node.Next
	}
	lists = append(lists, "->#")
	if root.Left != nil {
		lists = traverseNextPointer(root.Left, lists)
	}
	return lists
}

//Runtime: 10 ms, faster than 38.35% of Go online submissions for Populating Next Right Pointers in Each Node.
//Memory Usage: 7 MB, less than 18.84% of Go online submissions for Populating Next Right Pointers in Each Node.
func populatingNextRightPointers(root *Node) *Node {
	node := root
	if root == nil {
		return node
	}
	if node.Left != nil {
		node.Left.Next = node.Right
	}
	if node.Next != nil {
		if node.Right != nil {
			node.Right.Next = node.Next.Left
		}
	}

	fmt.Printf("Node:%v Next:%v Left:%v Right:%v\n", node, node.Next, node.Left, node.Right)
	populatingNextRightPointers(node.Left)
	populatingNextRightPointers(node.Right)

	return node
}

func mainPopulatingNextRightPointers() {
	lists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root := _116Slice2BTreeNode(lists)
	fmt.Printf("\n\n------------------\n\n")
	//root.midTraverse()
	root = populatingNextRightPointers(root)
	lists2 := []string{}
	lists2 = traverseNextPointer(root, lists2)
	fmt.Printf("\n\nPopulatingNextRightPointers:%v\n", lists2)
}
