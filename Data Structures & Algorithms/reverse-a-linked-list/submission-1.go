/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var dummyHead *ListNode

	for head != nil {
		if dummyHead == nil {
			dummyHead = &ListNode{
				Val: head.Val,
			}
		} else {
			current := &ListNode{
				Val: head.Val,
				Next: dummyHead,
			}
			dummyHead = current
		}
		head = head.Next
	}

	return dummyHead
}

// Time:  O(n)
// Space: O(n)