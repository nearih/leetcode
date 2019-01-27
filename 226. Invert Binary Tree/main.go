package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	r := switchbi(root)
	return r
}

func switchbi(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := &TreeNode{}
	temp = root.Right
	root.Right = root.Left
	root.Right = switchbi(root.Right)
	root.Left = temp
	root.Left = switchbi(root.Left)
	return root
}
