package linkList



func Circle(head *ListNode) bool {
	cur := head
	m := make(map[*ListNode]bool)

	for {
		if cur == nil {
			break
		}

		addr, ok := m[cur]

		if !ok {
			m[cur] = true
		}

		if addr {
			return true
		}
		cur = cur.Next
	}
	return false
}

