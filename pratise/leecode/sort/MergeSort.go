package sort

func MergeSort(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)

	return nums
}

func mergeSort(nums []int, left int, right int) {
	if left < right {
		middle := (left + right) >> 1

		mergeSort(nums, left, middle)

		mergeSort(nums, middle+1, right)

		merge(nums, left, middle, right)
	}
}

func merge(nums []int, left, middle, right int) {

	temp := make([]int, 0)
	i, j := left, middle+1

	for i <= middle && j <= right {
		if nums[i] < nums[j] {
			temp = append(temp, nums[i])
			i++
		} else {
			temp = append(temp, nums[j])
			j++
		}
	}

	if i <= middle {
		temp = append(temp, nums[i:middle+1]...)
	}

	if j <= right {
		temp = append(temp, nums[j:]...)
	}

	n := left

	for _, v := range temp {
		nums[n] = v
		n++
	}

}
