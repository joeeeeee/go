package linkList

type ListNode struct {
	data interface{}
	Next *ListNode
}

func CreateListNode(head *ListNode, n int) {
	cur := head

	for i := 0; i < n; i++ {
		node := ListNode{data: i}
		cur.Next = &node
		cur = cur.Next
	}
}

func Reverses(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func PrintListNode(head *ListNode) []interface{} {
	all := make([]interface{}, 0)
	cur := head
	for {
		if cur == nil {
			break
		}
		all = append(all, cur.data)
		cur = cur.Next
	}
	return all
}
