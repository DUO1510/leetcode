package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicates2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

//不满足空间复杂O（1）
func findDuplicates1(nums []int) []int {
	res := make([]int, 0)
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			res = append(res, nums[i])
		} else {
			m[nums[i]] = struct{}{}
		}
	}
	return res
}

func findDuplicates2(nums []int) []int {
	res := make([]int, 0)
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] < 0 {
			res = append(res, x)
		}
		nums[x-1] = -nums[x-1]
	}
	return res
}
