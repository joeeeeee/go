package sort

func BubbleSort(nums []int) []int {
	count := len(nums)
	for i := 0; i < count-1; i++ {
		flag := false
		for j := 0; j < count-1-i; j++ {
			greater := Greater(nums, j, j + 1)
			if  greater {
				nums = Swap(nums, j, j + 1)
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
	return nums
}

