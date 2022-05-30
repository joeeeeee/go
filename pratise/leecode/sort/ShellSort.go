package sort

func ShellSort(nums []int) []int {
	count := len(nums)
	step := count >> 1
	for {
		for i := step; i < count; i++ {
			for j := i; j - step >= 0 && Greater(nums, j-step, j); j -= step {
				Swap(nums, j-step, j)
			}
		}
		step = step >> 1
		if step == 0 {
			break
		}
	}

	return nums
}