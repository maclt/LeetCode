/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, []int{}, targetSum)
}

func dfs(node *TreeNode, prevSums []int, ts int) int {
	if node == nil {
		return 0
	}

    sums := make([]int, len(prevSums)+1)
	copy(sums[:len(sums)-1], prevSums)
	sums[len(sums)-1] = 0

	count := 0;
	for i:=0; i<len(sums); i++ {
		sum := sums[i] + node.Val

		if sum == ts {
			count = count + 1
		}

        sums[i] = sum;
	}

	return count + dfs(node.Left, sums, ts) + dfs(node.Right, sums, ts)
}
