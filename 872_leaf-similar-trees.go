/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
   seq1 := getSequence(root1);
   seq2 := getSequence(root2);

   return isSameSequence(seq1, seq2);
}

func getSequence(root *TreeNode) []int {
    queue := []TreeNode{*root};

    sequence := make([]int, 0);

    for len(queue) > 0 {
        var node = queue[0];
        queue = queue[1:];

        if node.Right != nil {
            queue = append([]TreeNode{*node.Right}, queue...);
        }

        if node.Left != nil {
            queue = append([]TreeNode{*node.Left}, queue...);
        }

        if node.Left == nil && node.Right == nil {
            sequence = append(sequence, node.Val);
        }
    }
    
    return sequence;
}

func isSameSequence(s1 []int, s2 []int) bool {
    if len(s1) != len(s2) {
        return false;
    }

    for i:=0; i<len(s1); i++ {
        if s1[i] != s2[i] {
            return false
        }
    }
    
    return true;
}
