/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	isCarriedOver := false

	var firstNode *ListNode = &ListNode{}
	var lastNode *ListNode = firstNode

	for l1 != nil || l2 != nil {
		var sum int

		if l1 != nil {
			sum = sum + l1.Val
            l1 = l1.Next
		} 
        
        if l2 != nil {
			sum = sum + l2.Val
            l2 = l2.Next
		}

		if isCarriedOver {
			sum++
		}

		if sum >= 10 {
			isCarriedOver = true
			sum = sum - 10
		} else {
			isCarriedOver = false
		}

		lastNode.Next = &ListNode{Val: sum}
		lastNode = lastNode.Next
	}

    if isCarriedOver {
        lastNode.Next = &ListNode{Val: 1}
    }

	return firstNode.Next
}
