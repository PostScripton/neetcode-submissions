/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var result int
    
	var dfs func(curr *TreeNode) int
	dfs = func(curr *TreeNode) int {
		if curr == nil {
			return 0
		}

		left, right := dfs(curr.Left), dfs(curr.Right)
		diameter := left + right

		if diameter > result {
			result = diameter
		}

		return 1 + max(left, right)
	}

	dfs(root)

	return result
}

//        Best     | Worst
// Time:  O(n)     | O(n)
// Space: O(log n) | O(n)
