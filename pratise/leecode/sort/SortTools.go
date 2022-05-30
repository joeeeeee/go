package sort

func Swap(nums []int, a int, b int) []int {
	nums[a], nums[b] = nums[b], nums[a]

	return nums
}

func Greater(nums []int, a int, b int) bool {
	return nums[a] > nums[b]
}

func LessThan(nums []int, a int, b int) bool {
	return nums[a] > nums[b]
}
