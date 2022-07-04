package questions

import "fmt"

// SearchInsert
// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
func SearchInsert(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	mid := start + (end-start)>>1
	for start <= end {
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = start + (end-start)>>1
		fmt.Println(start, mid, end)
	}

	if (mid == 0 || mid == len(nums)-1) && target <= nums[mid] {
		return mid
	}

	return mid + 1
}
