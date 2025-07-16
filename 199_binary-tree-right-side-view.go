/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	seq := make([]int, 0)

	if root == nil {
		return seq
	}

	queue := []TreeNode{*root}

	for len(queue) > 0 {
        seq = append(seq, queue[len(queue)-1].Val)

		var tempQueue = make([]TreeNode, 0)

		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				tempQueue = append(tempQueue, *queue[i].Left)
			}

			if queue[i].Right != nil {
				tempQueue = append(tempQueue, *queue[i].Right)
			}

		}

		queue = tempQueue
	}

	return seq
}
