package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees() *treeNode {
	t1 := NewBinaryTree()
	t1.Insert(1)
	t1.Insert(3)
	t1.Insert(2)
	t1.Insert(5)

	t2 := NewBinaryTree()
	t2.Insert(2)
	t2.Insert(1)
	t2.Insert(3)
	t2.Insert(4)
	t2.Insert(7)

	t1.BinaryTreeTraversal()
	fmt.Println("")
	t2.BinaryTreeTraversal()

	return nil
}

func main() {
	mergeTrees()
}
