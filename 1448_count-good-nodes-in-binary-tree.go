/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

    count := 0;
	if max <= node.Val {
		count = 1
		max = node.Val
	}

	count = count + dfs(node.Left, max)
	count = count + dfs(node.Right, max)

    return count
}