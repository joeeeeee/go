package sort

func SelectSort(nums []int) []int {

	count := len(nums)

	for i := 0; i < count - 1; i++ {
		min := i
		for j := i ; j <= count - 1 ; j ++ {
			if Greater(nums, min, j) {
				min = j
			}
		}
		if i != min {
			Swap(nums, i, min)
		}

	}

	return nums
}


//  i = 0 [1 9 2 8 3 7 4 6 5] j = 0

//  i = 1 [1 2 9 8 3 7 4 6 5] j = 1
//  i = 2 [1 2 3 8 9 7 4 6 5] j = 2
//  i = 3 [1 2 3 4 9 7 8 6 5]
//  i = 4 [1 2 3 4 5 7 8 6 9]
//  i = n - 1  [1 2 3 4 5 6 7 8 9]