package questions

/**
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

 

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
**/

func SearchRange(nums []int, target int) []int {
	left_l, left_r := 0, 0
	right_l, right_r := len(nums)-1, len(nums)-1

	for left_l <= right_l || left_r <= right_r {
		// 确定左边界
		if left_l <= right_l {
			mid_l := left_l + (right_l-left_l)/2
			if nums[mid_l] == target {
				right_l = mid_l - 1
			} else if nums[mid_l] > target {
				right_l = mid_l - 1
			} else if nums[mid_l] < target {
				left_l = mid_l + 1
			}
		}

		// 确定右边界
		if left_r <= right_r {
			mid_r := left_r + (right_r-left_r)/2
			if nums[mid_r] == target {
				left_r = mid_r + 1
			} else if nums[mid_r] > target {
				right_l = mid_r - 1
			} else if nums[mid_r] < target {
				left_r = mid_r + 1
			}
		}
	}
	left_ans, right_ans := -1, -1
	if left_l != len(nums) && nums[left_l] == target {
		left_ans = left_l
	}
	if left_r-1 != -1 && nums[left_r-1] == target {
		right_ans = left_r - 1
	}

	return []int{left_ans, right_ans}
}
