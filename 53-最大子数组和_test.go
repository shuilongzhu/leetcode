package leetcode

import "testing"

func Test53(t *testing.T) {
	t.Log(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// maxSubArray @description: leetCodeUrl:https://leetcode.cn/problems/maximum-subarray/
// @parameter nums
// @return int
func maxSubArray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}

		if result < nums[i] {
			result = nums[i]
		}

	}

	return result
}
