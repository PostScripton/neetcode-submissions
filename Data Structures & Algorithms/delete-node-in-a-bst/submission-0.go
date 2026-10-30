/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

    if val > root.Val {
		root.Right = deleteNode(root.Right, val)
	} else if val < root.Val {
		root.Left = deleteNode(root.Left, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minimum := findMinimum(root.Right)
			root.Val = minimum.Val
			root.Right = deleteNode(root.Right, minimum.Val)
		}
	}

	return root
}

func findMinimum(root *TreeNode) *TreeNode {
	curr := root

	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

//        Best     | Worst
// Time:  O(log n) | O(n)
// Space: O(log n) | O(n)