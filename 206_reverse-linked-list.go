/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev, curr, next *ListNode
    curr = head

    for curr != nil {
        // assign Next value before update it
        next = curr.Next

        // update Next value with previous value
        curr.Next = prev

        // move cursor to next iteration
        prev = curr
        curr = next
    }

    return prev
}
