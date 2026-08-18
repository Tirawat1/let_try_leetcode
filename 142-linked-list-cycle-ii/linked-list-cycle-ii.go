/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	curr := head
	fast := head
	// จุดตัดแรกบนวงกลม
	for fast != nil && fast.Next != nil {
		curr = curr.Next
		fast = fast.Next.Next
		if curr == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	curr = head
	for curr != fast {
		curr = curr.Next
		fast = fast.Next
	}
	return curr
}
