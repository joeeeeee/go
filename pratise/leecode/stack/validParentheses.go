package stack

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 1.字段串是否为空

func IsValid(s string) bool {
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	revPairs := make(map[byte]byte)

	for pairKey, pairValue := range pairs {
		revPairs[pairValue] = pairKey
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		ls, inPair := pairs[s[i]]
		_, inRevPair := revPairs[s[i]]
		// 如果不存在 map里，则不入栈
		if !(inPair || inRevPair) {
			continue
		}
		// 存在map 里，判断奇偶
		if inRevPair {
			// 如果是奇数 则为左括号，入栈
			stack = append(stack, s[i])
		} else {
			// 当前右侧， 如果队列为空
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != ls {
				return false
			}
			// 匹配到数据
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
