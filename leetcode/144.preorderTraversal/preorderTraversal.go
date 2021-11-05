package leetcode

import "leetcode-go/structures"

type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversal(root.Left)
		for _, t := range tmp {
			res = append(res, t)
		}
		tmp = preorderTraversal(root.Right)
		for _, t := range tmp {
			res = append(res, t)
		}
	}
	return res
}
