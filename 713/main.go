package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 6}
	k := 100
	fmt.Println(numSubarrayProductLessThanK(nums, k))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		product := nums[i]
		for j := i + 1; j < len(nums)+1; j++ {
			if product < k {
				res++
				//此处需判断 j 的值，防止边界溢出。
				if j < len(nums) {
					product = product * nums[j]
				}
				fmt.Printf("i is %d,j is %d,product is %d,res is %d\n", i, j-1, product, res)
			} else {
				break
			}
		}
	}
	return res
}
