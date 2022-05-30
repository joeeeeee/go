package leetcode

import (
	"fmt"
	"sort"
)
func ThreeSum(nums []int) [][]int {
	ret := [][]int{}

	sort.Ints(nums)

	repeat := map[string]bool{}

	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]

		res := TwoSum(nums[i+1:], target)

		if len(res) == 0 {
			continue
		}

		for _, value := range res {
			s := []int{nums[i]}
			k := string(nums[i])
			for _, x := range value {
				s = append(s, nums[x+i+1])
				k = fmt.Sprint(k, ",", string(nums[x+i+1]))
			}
			if repeat[k] {
				continue
			} else {
				repeat[k] = true
				ret = append(ret, s)
			}
		}
	}
   
	return ret
}

func TwoSum(nums []int, target int) map[int][]int {

	m := make(map[int]int)

	ret := map[int][]int{}

	i := 0

	for x, v := range nums {
		if _, ok := m[target-v]; ok {
			ret[i] = []int{m[target-v], x}
			i++
		} else {
			m[v] = x
		}
	}
	return ret
}
