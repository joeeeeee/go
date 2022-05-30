package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var idx int

	for i := 0; i < len(nums); i++ {
		idx = target - nums[i]

		if j, ok := m[nums[i]]; ok {
			return []int{j, i}
		}

		m[idx] = i
	}

	return []int {}
}


