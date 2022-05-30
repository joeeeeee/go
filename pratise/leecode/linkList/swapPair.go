package linkList


func SwapPair(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	curr := dummy

	for curr.Next != nil && curr.Next.Next != nil {
		first  := curr.Next
		second := curr.Next.Next
		first.Next = second.Next
		second.Next = first
		curr.Next = second
		curr = first
	}

	return dummy.Next
}


// [0 1 2 3 4 5 6 7 8 9] cur nil
// [1 0 3 2 5 4 7 6 9 8]
func SwapPair2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := head.Next.Next
	newHead := head.Next // newHead => 1
	newHead.Next = head
	head.Next = SwapPair2(last)

	return newHead
}
