package sort

func ShellSort2(nums []int) []int {
	count := len(nums)

	step := count >> 1

	for {
		for i := step ; i <= count - 1; i ++ {
			for j := i ; j - step >= 0 && Greater(nums, j - step, j) ; j -= step {
				Swap(nums, j , j - step)
			}
		}
		step = step >> 1
		if step == 0 {
			break
		}
	}

	return nums
}


//  [1 9 2 8 3 7 4 6 5]
// step 4  i = 4 -> 8 j = 4 swap j => j - step
//  [1 7 2 6 3 9 4 8 5]

// step 2  i = 2 -> 8 j = 2  swap j => j - step
//  [1 7 2 6 3 9 4 8 5]

