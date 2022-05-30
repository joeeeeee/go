package search

func BinarySearch(num []int, value int) int {
	return search(num, 0, len(num) - 1, value)
}


func search(num []int, left, right, value int) int {
	if left > right {
		return -1
	}

	mid := (right - left) >> 1 + left

	if num[mid] > value {
		return search(num, left,   mid - 1, value)
	}  else if num[mid] < value {
		return search(num, mid + 1, right, value)
	} else  {
		return mid
	}
}

/**
 * 根据二分搜索查找第一个匹配的元素
 */
func BinarySearch2(num []int, value int) int {

	n := len(num)

	low := 0

	high := n - 1

	for true {
		mid := (high - low) >> 1 + low

		if num[mid] > value {
			high = mid - 1
		} else if num[mid] < value {
			low = mid + 1
		} else {
			if mid == 0 || num[mid - 1] != num [mid] {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}
