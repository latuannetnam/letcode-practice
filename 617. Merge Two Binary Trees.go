// Letcode practice
// Author: Le Anh Tuan (latuannetnam@gmail.com)
// Created on 5/31/2022
// Letcode problem: 617. Merge Two Binary Trees
// Letcode link: https://leetcode.com/problems/merge-two-binary-trees/
// Level: easy
//You are given two binary trees root1 and root2.
//Imagine that when you put one of them to cover the other,
//some nodes of the two trees are overlapped while the others are not.
//You need to merge the two trees into a new binary tree.
//The merge rule is that if two nodes overlap,
//then sum node values up as the new value of the merged node.
//Otherwise, the NOT null node will be used as the node of the new tree.
////Return the merged tree.
////Note: The merging process must start from the root nodes of both trees.
package main

import "fmt"

func doMergeTrees2(node1 *TreeNode, node2 *TreeNode, node *TreeNode, isLeft bool) {

	if node1 != nil || node2 != nil {
		newNode := &TreeNode{}
		if isLeft {
			node.Left = newNode
		} else {
			node.Right = newNode
		}

		var node1Left *TreeNode = nil
		var node2Left *TreeNode = nil
		var node1Right *TreeNode = nil
		var node2Right *TreeNode = nil

		if node1 != nil && node2 != nil {
			newNode.Val = node1.Val + node2.Val
			node1Left = node1.Left
			node2Left = node2.Left
			node1Right = node1.Right
			node2Right = node2.Right
		} else if node1 != nil {
			newNode.Val = node1.Val
			node1Left = node1.Left
			node1Right = node1.Right
		} else if node2 != nil {
			newNode.Val = node2.Val
			node2Left = node2.Left
			node2Right = node2.Right
		}

		fmt.Printf("Merge node1:%v and node2:%v from node:%v to newnNode:%v on Branch:%t\n", node1, node2, node, newNode, isLeft)
		doMergeTrees2(node1Left, node2Left, newNode, true)
		doMergeTrees2(node1Right, node2Right, newNode, false)

	}

}

//Runtime: 39 ms, faster than 5.77% of Go online submissions for Merge Two Binary Trees.
//Memory Usage: 7.9 MB, less than 6.25% of Go online submissions for Merge Two Binary Trees.
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	root := &TreeNode{}
	fmt.Printf("Merge Root node1:%v and node2:%v to node:%v \n", root1, root2, root)
	if root1 != nil && root2 != nil {
		root.Val = root1.Val + root2.Val
		doMergeTrees2(root1.Left, root2.Left, root, true)
		doMergeTrees2(root1.Right, root2.Right, root, false)
	} else if root1 != nil {
		root.Val = root1.Val
		doMergeTrees2(root1.Left, nil, root, true)
		doMergeTrees2(root1.Right, nil, root, false)
	} else if root2 != nil {
		root.Val = root2.Val
		doMergeTrees2(nil, root2.Left, root, true)
		doMergeTrees2(nil, root2.Right, root, false)
	} else {
		return nil
	}

	return root
}

func doMergeNodes(node1, node2 *TreeNode) *TreeNode {
	node := &TreeNode{}
	if node1 != nil && node2 != nil {
		node.Val = node1.Val + node2.Val
		node.Left = doMergeNodes(node1.Left, node2.Left)
		node.Right = doMergeNodes(node1.Right, node2.Right)
	} else if node1 != nil {
		node.Val = node1.Val
		node.Left = doMergeNodes(node1.Left, nil)
		node.Right = doMergeNodes(node1.Right, nil)
	} else if node2 != nil {
		node.Val = node2.Val
		node.Left = doMergeNodes(nil, node2.Left)
		node.Right = doMergeNodes(nil, node2.Right)
	} else {
		return nil
	}
	return node
}

//Runtime: 10 ms, faster than 95.67% of Go online submissions for Merge Two Binary Trees.
//Memory Usage: 7.2 MB, less than 50.48% of Go online submissions for Merge Two Binary Trees.
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var root *TreeNode = nil
	if root1 != nil || root2 != nil {
		root = doMergeNodes(root1, root2)
	}

	return root

}

func mainMergeTrees() {
	//	Test Treenode
	lists := []int{1, 2, 3, 4}
	fmt.Printf("List:%v\n", lists)
	root1 := slice2BTree(lists)
	fmt.Printf("BTree:%v\n", root1)
	lists2 := bTree2slice(root1)
	//fmt.Printf("BTree 2 slice 1:%v\n", lists2)

	lists = []int{4, 5, 6, 7, 8}
	fmt.Printf("List:%v\n", lists)
	root2 := slice2BTree(lists)
	fmt.Printf("BTree:%v\n", root2)
	lists2 = bTree2slice(root2)
	//fmt.Printf("BTree 2 slice 2:%v\n", lists2)

	root := mergeTrees(root1, root2)
	lists2 = bTree2slice(root)
	fmt.Printf("\n\nBTree merge:%v\n", lists2)

}
