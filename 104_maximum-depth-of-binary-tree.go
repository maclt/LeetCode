/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    depth := 0;

    if root == nil {
        return depth;
    }

    queue := []TreeNode{*root};
    
    for len(queue) > 0 {
        tempQueue := make([]TreeNode, 0);
        for i:=0; i<len(queue); i++ {
            if queue[i].Left != nil {
                tempQueue = append(tempQueue, *queue[i].Left);
            }
            if queue[i].Right != nil {
                tempQueue = append(tempQueue, *queue[i].Right);
            }
        }

        queue = tempQueue;
        depth++;
    }

    return depth;
}
