package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = algorithms.TreeNode

func preorderTraversal(root *TreeNode) []int {
	var result []int
	helper(&result, root)
	return result
}

func helper(result *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	helper(result, node.Left)
	helper(result, node.Right)
}

func iteration(root *TreeNode) []int {
	var result []int
	stack := new(algorithms.Stack)
	stack.Push(root)
	for !stack.IsEmpty() {
		node := stack.Pop().(*TreeNode)

		result = append(result, node.Val)

		if node.Right != nil {
			stack.Push(node.Right)
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
	return result
}
