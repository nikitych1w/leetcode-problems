package containers

import (
	"fmt"
)

type treeNode struct {
	Val   interface{}
	Left  *treeNode
	Right *treeNode
}

type BinaryTree struct {
	root *treeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

func insertHelper(node *treeNode, data interface{}) {
	switch data.(type) {
	case int:
		if data.(int) < node.Val.(int) {
			if node.Left == nil {
				node.Left = &treeNode{Val: data, Left: nil, Right: nil}
			} else {
				insertHelper(node.Left, data)
			}
		}

		if data.(int) > node.Val.(int) {
			if node.Right == nil {
				node.Right = &treeNode{Val: data, Left: nil, Right: nil}
			} else {
				insertHelper(node.Right, data)
			}
		}
	}
}

func (bt *BinaryTree) Insert(data interface{}) {
	if bt.root == nil {
		bt.root = &treeNode{Val: data, Left: nil, Right: nil}
		return
	}

	insertHelper(bt.root, data)
}

func inorderTreeWalk(node *treeNode) {
	if node.Left != nil {
		inorderTreeWalk(node.Left)
	}
	if node.Val != nil {
		fmt.Println(node.Val)
	}
	if node.Right != nil {
		inorderTreeWalk(node.Right)
	}
}

func (bt *BinaryTree) BinaryTreeTraversal() {
	inorderTreeWalk(bt.root)
}
