package sort

func InsertSort(nums []int) []int {
	count := len(nums)
	for i := 0; i < count; i++ {
		for j := i; j > 0; j-- {
			if Greater(nums, j - 1, j) {
				Swap(nums , j - 1 , j)
			}
		}
	}
	return nums
}


// 1, 9,2,8,3,7,4,6,5

// 1,9 ,2,8,3,7,4,6,5

// 1,2,9 ,8,3,7,4,6,5

// 1,2,8,9 ,3,7,4,6,5

// 1,2,3,8,9 ,7,4,6,5

// 1,2,3,7,8,9 ,4,6,5

// 1,2,3,4,7,8, 9,6,5
