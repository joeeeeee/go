package sort

// 递归分组排序
func QuickSort(nums []int) []int{
	if len(nums) == 1 || len(nums) == 0 {
		return nums
	}

	middle := len(nums) >> 1

	mid := nums[middle]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == middle {
			continue
		}
		if nums[i] <= mid {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	values := make([]int, 0)

	left = QuickSort(left)
	right = QuickSort(right)

	values = append(values, left...)
	values = append(values, mid)
	values = append(values, right...)
	return values
}

